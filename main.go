package main

import (
	"runtime"

	"bitbucket.org/jtblin/kigo-api/cmd"
	"bitbucket.org/jtblin/kigo-api/db"
	"bitbucket.org/jtblin/kigo-api/job"
	"bitbucket.org/jtblin/kigo-api/pkg/manager"
	"bitbucket.org/jtblin/kigo-api/version"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/pflag"
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

	dbClient, err := db.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	appManager := manager.NewAppManager(db.NewAppRepository(dbClient))
	deploymentManager := manager.NewDeploymentManager(db.NewDeploymentRepository(dbClient), job.NewDeploymentExecutor())
	s := cmd.NewServer(c.ServerAddress, appManager, deploymentManager)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}

	w := cmd.NewWorker(deploymentManager)
	if err := w.Run(); err != nil {
		log.Fatal(err)
	}
}
