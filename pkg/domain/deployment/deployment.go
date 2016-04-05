package deployment

// DeploymentRepository represents a deployment resource interface
type DeploymentRepository interface {
	Create(*Deployment) (string, error)
	Get(string) (*Deployment, error)
}

// DeploymentExecutor represents a deployment job interface
type DeploymentExecutor interface {
	Schedule(*Deployment) error
}

// Deployment represents a deployment resource
type Deployment struct {
	ID       string
	App      string
	Env      []string
	Image    string
	Port     int32
	Replicas int32
	Status   int // TODO: add enum
	Zone     string
}

type deloymentChan chan *Deployment

// Queue serves to communicate new deployment events between db and domain
var Queue deloymentChan

// Enqueue adds a deployment to the channel
func Enqueue(d *Deployment) {
	Queue <- d
}
