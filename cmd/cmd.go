package cmd

import "github.com/spf13/pflag"

// Cmd is a representation of a cmd params
type Cmd struct {
	APIServer      string
	APIToken       string
	APIUser        string
	BuilderImage   string
	DBAddress      string
	DockerRegistry string
	ServerAddress  string
	Zone           string
	LogJSON        bool
	Verbose        bool
	Version        bool
}

// NewCmd creates a new cmd struct
func NewCmd() *Cmd {
	return &Cmd{
		APIServer:      "http://192.168.64.2:8080", // Mac OSX kube-solo
		BuilderImage:   "jtblin/kigo-builder:latest",
		DBAddress:      "localhost:28015",
		DockerRegistry: "jtblin",
		ServerAddress:  ":5080",
		Zone:           "connectapp.cloud",
	}
}

// AddFlags adds flags to the specified FlagSet
func (c *Cmd) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&c.APIServer, "api-server", c.APIServer, "Endpoint for the k8s api server")
	fs.StringVar(&c.APIToken, "api-token", c.APIToken, "Token to authenticate with the  k8sapi server")
	fs.StringVar(&c.APIUser, "api-user", c.APIUser, "User to authenticate with the k8s api server")
	fs.StringVar(&c.BuilderImage, "builder-image", c.BuilderImage, "Image for builder job")
	fs.StringVar(&c.DBAddress, "db-address", c.DBAddress, "Address of the db")
	fs.StringVar(&c.Zone, "dns-zone", c.Zone, "DNS zone to which to deploy services")
	fs.StringVar(&c.DockerRegistry, "docker-registry", c.DockerRegistry, "Docker registry to push images")
	fs.BoolVar(&c.LogJSON, "log-json", false, "Log as JSON")
	fs.StringVar(&c.ServerAddress, "server-addr", c.ServerAddress, "Address to listen on e.g. :80")
	fs.BoolVar(&c.Verbose, "verbose", false, "Verbose")
	fs.BoolVar(&c.Version, "version", false, "Print the version and exits")
}
