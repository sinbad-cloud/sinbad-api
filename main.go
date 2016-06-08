package main

import (
	"runtime"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/pflag"

	"github.com/sinbad-cloud/sinbad-api/cluster"
	"github.com/sinbad-cloud/sinbad-api/cmd"
	"github.com/sinbad-cloud/sinbad-api/db"
	"github.com/sinbad-cloud/sinbad-api/job"
	"github.com/sinbad-cloud/sinbad-api/pkg/manager"
	"github.com/sinbad-cloud/sinbad-api/version"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := cmd.NewCmd()
	c.AddFlags(pflag.CommandLine)
	pflag.Parse()

	if c.Verbose {
		log.SetLevel(log.DebugLevel)
	}
	if c.LogJSON {
		log.SetFormatter(&log.JSONFormatter{})
	}

	if c.Version {
		version.PrintVersionAndExit()
	}

	dbClient := db.NewClient(c.DBAddress)
	clusterClient, err := cluster.NewClient(c.APIServer, c.APIUser, c.APIToken)
	if err != nil {
		log.Fatal(err)
	}

	appManager := manager.NewAppManager(db.NewAppRepository(dbClient))
	buildManager := manager.NewBuildManager(db.NewBuildRepository(dbClient), job.NewBuildExecutor(clusterClient, c.DockerRegistry, c.BuilderImage, c.Zone))
	deploymentManager := manager.NewDeploymentManager(db.NewDeploymentRepository(dbClient), job.NewDeploymentExecutor(clusterClient))
	userManager := manager.NewUserManager(db.NewUserRepository(dbClient))
	s := cmd.NewServer(c.ServerAddress, appManager, buildManager, deploymentManager, userManager)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}

	w := cmd.NewWorker(buildManager, deploymentManager)
	if err := w.Run(); err != nil {
		log.Fatal(err)
	}
}
