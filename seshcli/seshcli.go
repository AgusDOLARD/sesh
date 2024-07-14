package seshcli

import (
	"github.com/joshmedeski/sesh/cmds"
	"github.com/joshmedeski/sesh/config"

	"github.com/urfave/cli/v2"
)

func App() cli.App {
	return cli.App{
		Name:    "sesh",
		Version: config.Version,
		Usage:   "Smart session manager for the terminal",
		Commands: []*cli.Command{
			cmds.List(),
			cmds.Connect(),
			cmds.Clone(),
		},
	}
}
