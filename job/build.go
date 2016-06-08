package job

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/apis/extensions"

	"github.com/sinbad-cloud/sinbad-api/cluster"
	"github.com/sinbad-cloud/sinbad-api/pkg/domain/build"
)

const (
	containerSourcePath string = "/src"
	dockerSocket        string = "/var/run/docker.sock"
	kubeNameSpace       string = "default"
)

type buildExec struct {
	client   *cluster.Client
	image    string
	registry string
	zone     string
}

// BuildJob represents a build job
type BuildJob struct {
	*build.Build
}

// NewBuildExecutor returns a new build executor
func NewBuildExecutor(client *cluster.Client, dockerRegistry, image, zone string) *buildExec {
	return &buildExec{
		image:    image,
		client:   client,
		registry: dockerRegistry,
		zone:     zone,
	}
}

func (be *buildExec) Schedule(b *build.Build) error {
	gitRepo := fmt.Sprintf("https://%s/%s/%s.git", b.Origin, b.Organization, b.Repository)
	jobID := fmt.Sprintf("build-%s", b.ID)
	job := extensions.Job{
		TypeMeta:   unversioned.TypeMeta{Kind: "Job", APIVersion: "extensions/v1beta1"},
		ObjectMeta: api.ObjectMeta{Name: jobID},
		Spec: extensions.JobSpec{
			Template: api.PodTemplateSpec{
				ObjectMeta: api.ObjectMeta{Labels: map[string]string{
					"id": jobID, "origin": b.Origin, "commit": b.Commit,
					"organization": b.Organization, "repository": b.Repository,
				}},
				Spec: api.PodSpec{
					Containers: []api.Container{
						{
							Name:            "builder",
							Image:           be.image,
							ImagePullPolicy: api.PullAlways,
							Args: []string{
								"--dir=" + containerSourcePath,
								"--dns-zone=" + be.zone,
								"--log-json",
								"--namespace=" + b.Organization,
								"--origin=" + b.Origin,
								"--registry=" + be.registry,
								"--repo=" + b.Repository,
								"--verbose",
							},
							VolumeMounts: []api.VolumeMount{{
								Name:      "src",
								MountPath: containerSourcePath,
							}, {
								Name:      "docker-socket",
								MountPath: "/var/run/docker.sock",
							}, {
								Name:      "docker-config",
								MountPath: "/root/.docker",
								ReadOnly:  true,
							}},
						},
					},
					Volumes: []api.Volume{{Name: "src", VolumeSource: api.VolumeSource{
						GitRepo: &api.GitRepoVolumeSource{Repository: gitRepo, Revision: b.Commit},
					}}, {Name: "docker-socket", VolumeSource: api.VolumeSource{
						HostPath: &api.HostPathVolumeSource{Path: dockerSocket},
					}}, {Name: "docker-config", VolumeSource: api.VolumeSource{
						// FIXME: hardcoding - how can we pass as config.json with a different secret name?
						// copy to a different location in builder?
						Secret: &api.SecretVolumeSource{SecretName: "config.json"},
					}}},
					RestartPolicy: "OnFailure",
				},
			},
		},
	}
	result, err := be.client.BatchClient.Jobs(kubeNameSpace).Create(&job)
	if err != nil {
		return err
	}
	log.Debugf("Job result: %+v", result)
	log.WithFields(log.Fields{
		"jobID":  jobID,
		"origin": b.Origin,
		"org":    b.Organization,
		"repo":   b.Repository,
		"commit": b.Commit,
	}).Info("Submitted job")
	return nil
}
