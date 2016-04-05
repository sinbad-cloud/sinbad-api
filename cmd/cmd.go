package cmd

import "github.com/spf13/pflag"

// Cmd is a representation of a cmd params
type Cmd struct {
	LogJSON       bool
	ServerAddress string
	Verbose       bool
	Version       bool
}

// NewCmd creates a new cmd struct
func NewCmd() *Cmd {
	return &Cmd{
		ServerAddress: ":5080",
	}
}

// AddFlags adds flags to the specified FlagSet
func (c *Cmd) AddFlags(fs *pflag.FlagSet) {
	fs.BoolVar(&c.LogJSON, "log-json", false, "Log as JSON")
	fs.StringVar(&c.ServerAddress, "server-addr", c.ServerAddress, "Address to listen on e.g. :80")
	fs.BoolVar(&c.Verbose, "verbose", false, "Verbose")
	fs.BoolVar(&c.Version, "version", false, "Print the version and exits")
}
