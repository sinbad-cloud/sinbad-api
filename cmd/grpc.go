package cmd

import (
	"errors"
	"net"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/dchest/uniuri"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/sinbad-cloud/sinbad-api/apipb"
	"github.com/sinbad-cloud/sinbad-api/pkg/domain/app"
	"github.com/sinbad-cloud/sinbad-api/pkg/domain/build"
	"github.com/sinbad-cloud/sinbad-api/pkg/domain/deployment"
	"github.com/sinbad-cloud/sinbad-api/pkg/domain/user"
	"github.com/sinbad-cloud/sinbad-api/pkg/manager"
)

// Server is a representation of an API server
type Server struct {
	Address           string
	AppManager        *manager.AppManager
	BuildManager      *manager.BuildManager
	DeploymentManager *manager.DeploymentManager
	UserManager       *manager.UserManager
}

// NewServer creates a new server
func NewServer(addr string, appManager *manager.AppManager, buildManager *manager.BuildManager, deploymentManager *manager.DeploymentManager, userManager *manager.UserManager) *Server {
	return &Server{
		Address:           addr,
		AppManager:        appManager,
		BuildManager:      buildManager,
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
	api.RegisterBuildServiceServer(grpcServer, s)
	api.RegisterDeploymentServiceServer(grpcServer, s)
	api.RegisterAuthServiceServer(grpcServer, s)
	return grpcServer.Serve(listener)
}

// SignIn signs a user in
func (s *Server) SignIn(cxt context.Context, usr *api.User) (*api.AuthResponse, error) {
	log.Infof("Signing user %s with context %#v", usr.Email, cxt)
	u, err := s.UserManager.UserRepo.Get(usr.Email)
	if err != nil {
		log.Errorf("Get user error: %v", err)
		return nil, escapeError(err)
	}
	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(usr.Password)); err != nil {
		return nil, escapeError(errors.New("incorrect password"))
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
	log.Infof("Receiving sign up request %s with context %#v", usr.Email, cxt)
	token := uniuri.New()
	_, err := s.UserManager.UserRepo.Create(&user.User{Email: usr.Email, Name: usr.Name, Password: usr.Password, Token: token})
	if err != nil {
		log.Errorf("Create user error: %v", err)
		return nil, escapeError(err)
	}
	return &api.AuthResponse{Email: usr.Email, Name: usr.Name, Token: token}, nil
}

// GetDeployment gets a deployment from its ID
func (s *Server) GetDeployment(cxt context.Context, job *api.DeploymentJob) (*api.DeploymentRequest, error) {
	log.Infof("Getting deployment %s with context %#v", job.Id, cxt)
	d, err := s.DeploymentManager.DeploymentRepo.Get(job.Id)
	if err != nil {
		log.Errorf("Get deployment error: %v", err)
		return nil, escapeError(err)
	}
	return &api.DeploymentRequest{App: d.App, Env: d.Env, Image: d.Image, Port: d.Port, Replicas: d.Replicas, Zone: d.Zone}, nil
}

// CreateDeployment creates a deployment in the DB
func (s *Server) CreateDeployment(cxt context.Context, d *api.DeploymentRequest) (*api.DeploymentJob, error) {
	log.Infof("Receiving deployment request %#v with context %#v", d, cxt)
	ID, err := s.DeploymentManager.DeploymentRepo.Create(&deployment.Deployment{App: d.App, Env: d.Env, Image: d.Image, Port: d.Port, Replicas: d.Replicas, Zone: d.Zone})
	if err != nil {
		log.Errorf("Create deployment error: %v", err)
		return nil, escapeError(err)
	}
	return &api.DeploymentJob{Id: ID}, nil
}

// GetApp gets an app from its name
func (s *Server) GetApp(cxt context.Context, a *api.App) (*api.App, error) {
	log.Infof("Getting app %s with context %#v", a.Name, cxt)
	r, err := s.AppManager.AppRepo.Get(a.Name)
	if err != nil {
		log.Errorf("Get app error: %v", err)
		return nil, escapeError(err)
	}
	return &api.App{Name: r.Name, Config: r.Config, Owner: r.User, Repo: r.RepoURL}, nil
}

// CreateApp creates an app in the DB
func (s *Server) CreateApp(cxt context.Context, a *api.App) (*api.AppCreateResponse, error) {
	log.Infof("Receiving app request %#v with context %#v", a, cxt)
	_, err := s.AppManager.AppRepo.Create(&app.App{Name: a.Name, RepoURL: a.Repo, User: a.Owner})

	if err != nil {
		log.Errorf("Create app error: %v", err)
		return nil, escapeError(err)
	}
	return &api.AppCreateResponse{Name: a.Name}, nil
}

// GetAppConfig gets the config for an app
func (s *Server) GetAppConfig(cxt context.Context, cfg *api.ConfigRequest) (*api.ConfigResponse, error) {
	log.Infof("Getting app config %s with context %#v", cfg.Name, cxt)
	c, err := s.AppManager.AppRepo.GetConfig(cfg.Name, cfg.Key)
	if err != nil {
		log.Errorf("Get app config error: %v", err)
		return nil, escapeError(err)
	}
	return &api.ConfigResponse{Key: cfg.Key, Value: c[cfg.Key]}, nil
}

// SetAppConfig sets an app config
func (s *Server) SetAppConfig(cxt context.Context, cfg *api.ConfigRequest) (*api.ConfigResponse, error) {
	log.Infof("Setting app config %s with context %#v", cfg.Name, cxt)
	c, err := s.AppManager.AppRepo.SetConfig(cfg.Name, cfg.Key, cfg.Value)
	if err != nil {
		log.Errorf("Set app config error: %v", err)
		return nil, escapeError(err)
	}
	return &api.ConfigResponse{Key: cfg.Key, Value: c[cfg.Key]}, nil
}

// GetBuild gets a build from its ID
func (s *Server) GetBuild(cxt context.Context, b *api.GetBuildRequest) (*api.GetBuildResponse, error) {
	log.Infof("Getting build %s with context %#v", b.Id, cxt)
	r, err := s.BuildManager.BuildRepo.Get(b.Id)
	if err != nil {
		log.Errorf("Get app error: %v", err)
		return nil, escapeError(err)
	}
	return &api.GetBuildResponse{
		Id: r.ID,
		// TODO: returns entire build or status?
	}, nil
}

// CreateBuild creates a new build
func (s *Server) CreateBuild(cxt context.Context, b *api.CreateBuildRequest) (*api.CreateBuildResponse, error) {
	log.Infof("Receiving build request %#v with context %#v", b, cxt)
	ID, err := s.BuildManager.BuildRepo.Create(&build.Build{
		App:          b.App,
		Author:       b.Author,
		Commit:       b.Commit,
		Organization: b.Organisation,
		Origin:       b.Origin,
		Repository:   b.Repository,
		Timestamp:    time.Unix(b.Timestamp.Seconds, int64(b.Timestamp.Nanos)),
	})

	if err != nil {
		log.Errorf("Create app error: %v", err)
		return nil, escapeError(err)
	}
	return &api.CreateBuildResponse{Id: ID}, nil
}

func escapeError(err error) error {
	// https://github.com/grpc/grpc-go/issues/576
	msg := strings.Replace(err.Error(), "\n", " ", -1)
	return errors.New(msg)
}
