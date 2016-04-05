package cmd

import (
	"bitbucket.org/jtblin/kigo-api/pkg/domain/deployment"
	"bitbucket.org/jtblin/kigo-api/pkg/manager"
)

// Worker is a representation of a worker
type Worker struct {
	DeploymentManager *manager.DeploymentManager
}

// NewWorker creates a new worker
func NewWorker(deploymentManager *manager.DeploymentManager) *Worker {
	return &Worker{
		DeploymentManager: deploymentManager,
	}
}

func (w *Worker) Run() error {
	go w.ProcessDeployments()
	return nil
}

// TODO: make more generic
// ProcessDeployments process jobs from the deployment queue
func (w *Worker) ProcessDeployments() {
	for deployment := range deployment.Queue {
		go w.DeploymentManager.DeploymentExec.Schedule(deployment)
	}
}
