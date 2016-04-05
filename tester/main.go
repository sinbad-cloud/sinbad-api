package main

import (
	"flag"

	"bitbucket.org/jtblin/kigo-api/apipb"

	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:5080"
)

var readonly bool
var ID string

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := api.NewDeploymentServiceClient(conn)

	if readonly {
		d, err := c.GetDeployment(context.TODO(), &api.DeploymentJob{Id: ID})
		if err != nil {
			log.Fatalf("Get deployment error: %v", err)
		}
		log.Infof("Deployment: %+v", d)
	} else {
		d, err := c.CreateDeployment(context.TODO(), &api.DeploymentRequest{
			App:      "foobar",
			Env:      []string{"FOO"},
			Image:    "busybox:latest",
			Port:     8080,
			Replicas: 2,
			Zone:     "myzone.com",
		})
		if err != nil {
			log.Fatalf("Create deployment error: %v", err)
		}
		log.Infof("Deployment created: %+v", d)
	}

}

func init() {
	flag.BoolVar(&readonly, "readonly", false, "Get instead of create")
	flag.StringVar(&ID, "id", "", "ID of document to retrieve")
}
