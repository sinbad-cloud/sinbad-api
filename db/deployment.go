package db

import (
	"errors"

	"bitbucket.org/jtblin/kigo-api/pkg/domain/deployment"

	r "github.com/dancannon/gorethink"
)

const deploymentTable = "deployment"

// DeploymentModel represents a deployment resource
type DeploymentModel struct {
	ID       string   `gorethink:"id,omitempty"`
	App      string   `gorethink:"app"`
	Env      []string `gorethink:"env"`
	Image    string   `gorethink:"image"`
	Port     int32    `gorethink:"port"`
	Replicas int32    `gorethink:"replicas"`
	Status   int      `gorethink:"status"` // TODO: add enum
	Zone     string   `gorethink:"zone"`
}

type deploymentRepo struct {
	*rethinkClient
}

// NewDeploymentRepository is an implementation for an DeploymentRepository
func NewDeploymentRepository(c *rethinkClient) deployment.DeploymentRepository {
	return &deploymentRepo{
		rethinkClient: c,
	}
}

// WatchDeployments watch deployments and send to queue
func (c *rethinkClient) watchDeployments() error {
	cursor, err := r.Table(deploymentTable).Changes().Run(c.session)
	if err != nil {
		return err
	}
	//deploymentChan := make(chan r.ChangeResponse)
	go func() {
		var model DeploymentModel // d deployment.Deployment
		// TODO: use channel and cursor.Listen instead?
		// TODO: check if new or updated or deleted
		// i.e. newValue != nil, oldValue != nil
		for cursor.Next(&model) {
			deployment.Enqueue(&deployment.Deployment{
			// TODO: fill this in
			})
		}
	}()
	return nil
}

// Get returns a deployment
func (rd *deploymentRepo) Get(ID string) (*deployment.Deployment, error) {
	cursor, err := r.Table(deploymentTable).Get(ID).Run(rd.session)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()
	if cursor.IsNil() {
		return nil, errors.New("no record found")
	}
	var model DeploymentModel
	if err = cursor.One(&model); err != nil {
		return nil, err
	}
	return &deployment.Deployment{
	// TODO: fill this in
	}, nil
}

// Create creates a new deployment
func (dr *deploymentRepo) Create(d *deployment.Deployment) (string, error) {
	if d.App == "" || d.Image == "" || d.Port == 0 || d.Zone == "" {
		return "", errors.New("missing required field")
	}
	if d.Replicas == 0 {
		d.Replicas = 1
	}
	response, err := r.Table(deploymentTable).Insert(DeploymentModel{
		App: d.App,
		// TODO: fill this in
	}).RunWrite(dr.session)
	if err != nil {
		return "", err
	}
	return response.GeneratedKeys[0], nil
}

func init() {
	tables = append(tables, deploymentTable)
}
