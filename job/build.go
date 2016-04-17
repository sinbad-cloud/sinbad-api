package job

import (
	"bitbucket.org/jtblin/kigo-api/pkg/domain/build"
)

type buildExec struct{}

// BuildJob represents a build job
type BuildJob struct {
	*build.Build
	// TODO: add stuff
}

// NewBuildExecutor returns a new build executor
func NewBuildExecutor() *buildExec {
	return &buildExec{}
}

func (be *buildExec) Schedule(d *build.Build) error {
	// TODO: schedule the job on kubernetes
	return nil
}
