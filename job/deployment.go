package job

import (
	"bitbucket.org/jtblin/kigo-api/cluster"
	"bitbucket.org/jtblin/kigo-api/pkg/domain/deployment"
)

type deploymentExec struct {
	client *cluster.Client
}

// DeploymentJob represents a deployment job
type DeploymentJob struct {
	*deployment.Deployment
	// TODO: add stuff
}

// NewDeploymentExecutor returns a new deployment executor
func NewDeploymentExecutor(client *cluster.Client) *deploymentExec {
	return &deploymentExec{client: client}
}

func (de *deploymentExec) Schedule(d *deployment.Deployment) error {
	// TODO: schedule the job on kubernetes
	return nil
}
