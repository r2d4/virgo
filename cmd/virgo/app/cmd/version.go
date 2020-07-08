package cmd

import (
	"io"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"matt-rickard.com/virgo/cmd/virgo/app/cmd/flags"
	"matt-rickard.com/virgo/pkg/version"
)

var versionFlag = flags.NewTemplateFlag("{{.Version}}\n", version.Info{})

func NewCmdVersion(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version information",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunVersion(out, cmd)
		},
	}

	cmd.Flags().VarP(versionFlag, "output", "o", versionFlag.Usage())
	return cmd
}

func RunVersion(out io.Writer, cmd *cobra.Command) error {
	if err := versionFlag.Template().Execute(out, version.Get()); err != nil {
		return errors.Wrap(err, "executing template")
	}
	return nil
}
