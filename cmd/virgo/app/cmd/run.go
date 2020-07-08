package cmd

import (
	"io"

	"github.com/spf13/cobra"
	"matt-rickard.com/virgo/pkg/run"
)

func NewCmdRun(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return run.Run(args[0])
		},
	}

	return cmd
}
