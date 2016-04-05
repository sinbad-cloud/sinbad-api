package cmd

import (
	"errors"
	"net"
	"strings"

	"bitbucket.org/jtblin/kigo-api/apipb"
	"bitbucket.org/jtblin/kigo-api/pkg/domain/app"
	"bitbucket.org/jtblin/kigo-api/pkg/domain/deployment"
	"bitbucket.org/jtblin/kigo-api/pkg/manager"

	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Server is a representation of an API server
type Server struct {
	Address           string
	AppManager        *manager.AppManager
	DeploymentManager *manager.DeploymentManager
	Verbose           bool
	Version           bool
}

// NewServer creates a new server
func NewServer(addr string, appManager *manager.AppManager, deploymentManager *manager.DeploymentManager) *Server {
	return &Server{
		Address:           addr,
		AppManager:        appManager,
		DeploymentManager: deploymentManager,
	}
}

// Run runs the server
func (s *Server) Run() error {
	listener, err := net.Listen("tcp", s.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	api.RegisterDeploymentServiceServer(grpcServer, s)
	api.RegisterAppServiceServer(grpcServer, s)
	return grpcServer.Serve(listener)
}

// GetDeployment gets a deployment from its ID
func (s *Server) GetDeployment(cxt context.Context, job *api.DeploymentJob) (*api.DeploymentRequest, error) {
	log.Infof("Getting deployment %s with context %#v", job.Id, cxt)
	d, err := s.DeploymentManager.DeploymentRepo.Get(job.Id)
	if err != nil {
		log.Errorf("Get deployment error: %v", escapeError(err))
		return nil, err
	}
	return &api.DeploymentRequest{App: d.App, Env: d.Env, Image: d.Image, Port: d.Port, Replicas: d.Replicas, Zone: d.Zone}, nil
}

// CreateDeployment creates a deployment in the DB
func (s *Server) CreateDeployment(cxt context.Context, d *api.DeploymentRequest) (*api.DeploymentJob, error) {
	log.Infof("Receiving deployment request %#v with context %#v", d, cxt)
	ID, err := s.DeploymentManager.DeploymentRepo.Create(&deployment.Deployment{App: d.App, Env: d.Env, Image: d.Image, Port: d.Port, Replicas: d.Replicas, Zone: d.Zone})
	if err != nil {
		log.Errorf("Create deployment error: %v", escapeError(err))
		return nil, err
	}
	return &api.DeploymentJob{Id: ID}, nil
}

// GetDeployment gets an app from its name
func (s *Server) GetApp(cxt context.Context, a *api.App) (*api.App, error) {
	log.Infof("Getting app %s with context %#v", a.Name, cxt)
	r, err := s.AppManager.AppRepo.Get(a.Name)
	if err != nil {
		log.Errorf("Get app error: %v", escapeError(err))
		return nil, err
	}
	return &api.App{Name: r.Name, Owner: r.User}, nil
}

// CreateApp creates an app in the DB
func (s *Server) CreateApp(cxt context.Context, a *api.App) (*api.AppCreateResponse, error) {
	log.Infof("Receiving app request %#v with context %#v", a, cxt)
	// TODO: ID?
	_, err := s.AppManager.AppRepo.Create(&app.App{Name: a.Name, User: a.Owner})

	if err != nil {
		log.Errorf("Create app error: %v", escapeError(err))
		return nil, err
	}
	return &api.AppCreateResponse{Status: "OK"}, nil
}

func escapeError(err error) error {
	// https://github.com/grpc/grpc-go/issues/576
	msg := strings.Replace(err.Error(), "\\n", " ", -1)
	return errors.New(msg)
}
