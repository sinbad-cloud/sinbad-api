// Package **manager** performs resource-related "business-logic" functions on a resource and related objects.
// This is in contrast to the Repositories which perform little more than CRUD operations.
package manager

import (
	"bitbucket.org/jtblin/kigo-api/pkg/domain/app"
	"bitbucket.org/jtblin/kigo-api/pkg/domain/build"
	"bitbucket.org/jtblin/kigo-api/pkg/domain/deployment"
	"bitbucket.org/jtblin/kigo-api/pkg/domain/user"
)

// AppManager performs app-related "business-logic" functions on an app and related objects.
type AppManager struct {
	AppRepo app.AppRepository
}

// NewAppManager initialises a new app manager
func NewAppManager(appRepo app.AppRepository) *AppManager {
	return &AppManager{
		AppRepo: appRepo,
	}
}

// BuildManager performs build-related "business-logic" functions on a build and related objects.
type BuildManager struct {
	BuildRepo build.BuildRepository
	BuildExec build.BuildExecutor
}

// NewBuildManager initialises a new build manager
func NewBuildManager(buildRepo build.BuildRepository, buildExec build.BuildExecutor) *BuildManager {
	return &BuildManager{
		BuildRepo: buildRepo,
		BuildExec: buildExec,
	}
}

// DeploymentManager performs deployment-related "business-logic" functions on a deployment and related objects.
type DeploymentManager struct {
	DeploymentRepo deployment.DeploymentRepository
	DeploymentExec deployment.DeploymentExecutor
}

// NewAppManager initialises a new deployment manager
func NewDeploymentManager(deploymentRepo deployment.DeploymentRepository, deploymentExec deployment.DeploymentExecutor) *DeploymentManager {
	return &DeploymentManager{
		DeploymentRepo: deploymentRepo,
		DeploymentExec: deploymentExec,
	}
}

// UserManager performs user-related "business-logic" functions on a user and related objects.
type UserManager struct {
	UserRepo user.UserRepository
}

// NewUserManager initialises a new user manager
func NewUserManager(userRepo user.UserRepository) *UserManager {
	return &UserManager{
		UserRepo: userRepo,
	}
}
