package app

import (
	"os"

	"matt-rickard.com/virgo/cmd/virgo/app/cmd"
)

func Run() error {
	c := cmd.NewRootCommand(os.Stdout, os.Stderr)
	return c.Execute()
}
