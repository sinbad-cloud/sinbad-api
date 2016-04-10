package cmd

import (
	"errors"
	"net"
	"strings"

	"bitbucket.org/jtblin/kigo-api/apipb"
	"bitbucket.org/jtblin/kigo-api/pkg/domain/app"
	"bitbucket.org/jtblin/kigo-api/pkg/domain/deployment"
	"bitbucket.org/jtblin/kigo-api/pkg/domain/user"
	"bitbucket.org/jtblin/kigo-api/pkg/manager"

	log "github.com/Sirupsen/logrus"
	"github.com/dchest/uniuri"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Server is a representation of an API server
type Server struct {
	Address           string
	AppManager        *manager.AppManager
	DeploymentManager *manager.DeploymentManager
	UserManager       *manager.UserManager
	Verbose           bool
	Version           bool
}

// NewServer creates a new server
func NewServer(addr string, appManager *manager.AppManager, deploymentManager *manager.DeploymentManager, userManager *manager.UserManager) *Server {
	return &Server{
		Address:           addr,
		AppManager:        appManager,
		DeploymentManager: deploymentManager,
		UserManager:       userManager,
	}
}

// Run runs the server
func (s *Server) Run() error {
	listener, err := net.Listen("tcp", s.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	api.RegisterAppServiceServer(grpcServer, s)
	api.RegisterDeploymentServiceServer(grpcServer, s)
	api.RegisterAuthServiceServer(grpcServer, s)
	return grpcServer.Serve(listener)
}

// SignIn signs a user in
func (s *Server) SignIn(cxt context.Context, usr *api.User) (*api.AuthResponse, error) {
	log.Infof("Signing user %s with context %#v", usr.Email, cxt)
	u, err := s.UserManager.UserRepo.Get(usr.Email)
	if err != nil {
		log.Errorf("Get user error: %v", escapeError(err))
		return nil, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(usr.Password)); err != nil {
		return nil, errors.New("Incorrect password")
	}
	return &api.AuthResponse{Email: u.Email, Name: u.Name, Token: u.Token}, nil
}

// Reset resets a user password
func (s *Server) Reset(cxt context.Context, usr *api.User) (*api.AuthResponse, error) {
	// TODO: implement
	return &api.AuthResponse{}, nil
}

// SignUp creates a user in the DB
func (s *Server) SignUp(cxt context.Context, usr *api.User) (*api.AuthResponse, error) {
	// FIXME: do not log password
	log.Infof("Receiving sign up request %#v with context %#v", usr, cxt)
	token := uniuri.New()
	_, err := s.UserManager.UserRepo.Create(&user.User{Email: usr.Email, Name: usr.Name, Password: usr.Password, Token: token})
	if err != nil {
		log.Errorf("Create user error: %v", escapeError(err))
		return nil, err
	}
	return &api.AuthResponse{Email: usr.Email, Name: usr.Name, Token: token}, nil
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
