package manager

import (
	"bitbucket.org/jtblin/kigo-api/pkg/domain/app"
	"bitbucket.org/jtblin/kigo-api/pkg/domain/deployment"
)

// AppManager performs app-related "business-logic" functions on ab app and related objects.
// This is in contrast to the Repos which perform little more than CRUD operations.
type AppManager struct {
	AppRepo app.AppRepository
}

// NewAppManager initialises a new app manager
func NewAppManager(appRepo app.AppRepository) *AppManager {
	return &AppManager{
		AppRepo: appRepo,
	}
}

// DeploymentManager performs deployment-related "business-logic" functions on a deployment and related objects.
// This is in contrast to the Repos which perform little more than CRUD operations.
type DeploymentManager struct {
	DeploymentRepo deployment.DeploymentRepository
	DeploymentExec deployment.DeploymentExecutor
}

// NewAppManager initialises a new app manager
func NewDeploymentManager(deploymentRepo deployment.DeploymentRepository, deploymentExec deployment.DeploymentExecutor) *DeploymentManager {
	return &DeploymentManager{
		DeploymentRepo: deploymentRepo,
		DeploymentExec: deploymentExec,
	}
}
