package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/sinbad-cloud/sinbad-api/apipb"
)

const (
	address = "localhost:5080"
)

type getOptions struct{}
type createOptions struct {
	app      string
	email    string
	env      []string
	image    string
	name     string
	pwd      string
	port     int32
	replicas int32
	zone     string
}

func main() {
	cmd := newCmd(os.Stdin, os.Stdout, os.Stderr)
	cmd.Execute()
}

func connect() (*grpc.ClientConn, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("Did not connect: %v", err)
	}
	return conn, nil
}

func runCreate(cmd *cobra.Command, out io.Writer, args []string, options *createOptions) error {
	conn, err := connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	switch args[0] {
	case "deployment, deployments", "dep":
		c := api.NewDeploymentServiceClient(conn)
		d, err := c.CreateDeployment(context.TODO(), &api.DeploymentRequest{
			App:      options.app,
			Env:      options.env,
			Image:    options.image,
			Port:     options.port,
			Replicas: options.replicas,
			Zone:     options.zone,
		})
		if err != nil {
			return fmt.Errorf("create deployment error: %v", err)
		}
		log.Infof("Deployment created: %+v", d)
	case "users", "user", "usr", "u":
		c := api.NewAuthServiceClient(conn)
		usr, err := c.SignUp(context.TODO(), &api.User{
			Email:    options.email,
			Name:     options.name,
			Password: options.pwd,
		})
		if err != nil {
			return fmt.Errorf("create user error: %v", err)
		}
		log.Infof("User created: %+v", usr)
	}

	return nil
}

func runGet(cmd *cobra.Command, out io.Writer, args []string, options *getOptions) error {
	if len(args) == 0 {
		cmd.Help()
		os.Exit(1)
	}

	conn, err := connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	c := api.NewDeploymentServiceClient(conn)

	switch args[0] {
	case "deployment", "deployments", "de":
		switch len(args) {
		case 1:
			// TODO: list deployments
			return errors.New("listing deployments not implemented")
		case 2:
			d, err := c.GetDeployment(context.TODO(), &api.DeploymentJob{Id: args[1]})
			if err != nil {
				return fmt.Errorf("Get deployment error: %v", err)
			}
			log.Infof("Deployment: %+v", d)
		default:
			cmd.Help()
			os.Exit(1)
		}
	}
	return nil
}

func checkErr(err error, handleErr func(string)) {
	if err == nil {
		return
	}
	msg := fmt.Sprintf("error: %s", err.Error())
	handleErr(msg)
}

func fatal(msg string) {
	// add newline if needed
	if !strings.HasSuffix(msg, "\n") {
		msg += "\n"
	}

	fmt.Fprint(os.Stderr, msg)
	os.Exit(1)
}

func newCreateCommand(out io.Writer) *cobra.Command {
	options := &createOptions{}
	cmd := &cobra.Command{
		Use:     "create TYPE",
		Example: "sinbadctl create user",
		Short:   "sinbadctl create user from the API",
		Long:    "sinbadctl create user from the API",
		Run: func(cmd *cobra.Command, args []string) {
			checkErr(runCreate(cmd, out, args, options), fatal)
		},
	}
	cmd.Flags().StringVarP(&options.app, "app", "a", "", "App name")
	cmd.Flags().StringVarP(&options.email, "email", "m", "", "Email")
	cmd.Flags().StringSliceVarP(&options.env, "env", "e", []string{}, "Env name")
	cmd.Flags().StringVarP(&options.image, "image", "i", "", "Image name")
	cmd.Flags().StringVarP(&options.name, "name", "n", "", "User name")
	cmd.Flags().StringVarP(&options.pwd, "password", "w", "", "Password")
	cmd.Flags().Int32VarP(&options.port, "port", "p", 8080, "Port")
	cmd.Flags().Int32VarP(&options.replicas, "replicas", "r", 1, "Port")
	cmd.Flags().StringVarP(&options.zone, "zone", "z", "", "Zone name")
	return cmd
}

func newGetCommand(out io.Writer) *cobra.Command {
	options := &getOptions{}
	cmd := &cobra.Command{
		Use:     "get TYPE [ID]",
		Example: "sinbadctl get users",
		Short:   "sinbadctl get resources from the API",
		Long:    "sinbadctl get resources from the API",
		Run: func(cmd *cobra.Command, args []string) {
			checkErr(runGet(cmd, out, args, options), fatal)
		},
	}
	return cmd
}

func newCmd(in io.Reader, out, err io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sinbadctl",
		Short: "sinbadctl communicates with sinbad-api",
		Long:  "sinbadctl communicates with sinbad-api",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(newGetCommand(out))
	cmd.AddCommand(newCreateCommand(out))
	return cmd
}
