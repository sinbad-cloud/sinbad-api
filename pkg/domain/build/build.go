package build

import (
	"time"
)

// Build represents a build resource
type Build struct {
	App          string
	Author       string
	Commit       string
	ID           string
	Organization string
	Origin       string
	Repository   string
	Timestamp    time.Time
}

// BuildRepository represents a build interface
type BuildRepository interface {
	Create(build *Build) (string, error)
	Get(ID string) (*Build, error)
}

// BuildExecutor represents a build job interface
type BuildExecutor interface {
	Schedule(*Build) error
}

type buildQueue chan *Build

// Queue serves to communicate new build events between db and domain
var Queue buildQueue

// Enqueue adds a deployment to the channel
func Enqueue(d *Build) {
	Queue <- d
}
