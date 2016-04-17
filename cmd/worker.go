package cmd

import (
	"bitbucket.org/jtblin/kigo-api/pkg/domain/build"
	"bitbucket.org/jtblin/kigo-api/pkg/domain/deployment"
	"bitbucket.org/jtblin/kigo-api/pkg/manager"
)

// Worker is a representation of a worker
type Worker struct {
	BuilderImage      string
	DockerRegistry    string
	Zone              string

	BuildManager      *manager.BuildManager
	DeploymentManager *manager.DeploymentManager
}

// NewWorker creates a new worker
func NewWorker(khost, kusername, ktoken, dockerRegistry, image, zone string, buildManager *manager.BuildManager, deploymentManager *manager.DeploymentManager) *Worker {
	return &Worker{
		BuildManager:      buildManager,
		DeploymentManager: deploymentManager,
		BuilderImage:      image,
		DockerRegistry:    dockerRegistry,
		Zone:              zone,
	}
}

func (w *Worker) Run() error {
	go w.ProcessDeployments()
	return nil
}

// TODO: make more generic
// ProcessBuilds process jobs from the build queue
func (w *Worker) ProcessBuilds() {
	for build := range build.Queue {
		go w.BuildManager.BuildExec.Schedule(build)
	}
}

// ProcessDeployments process jobs from the deployment queue
func (w *Worker) ProcessDeployments() {
	for deployment := range deployment.Queue {
		go w.DeploymentManager.DeploymentExec.Schedule(deployment)
	}
}
