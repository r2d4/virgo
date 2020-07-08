package cmd

import (
	"io"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"matt-rickard.com/virgo/pkg/constants"
	"matt-rickard.com/virgo/pkg/version"
)

var (
	v string
)

func NewRootCommand(out, err io.Writer) *cobra.Command {
	ver := version.Get()
	rootCmd := &cobra.Command{
		Use:   "email",
		Short: "",
	}
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if err := SetUpLogs(err, v); err != nil {
			return err
		}
		rootCmd.SilenceUsage = true
		logrus.Infof("Version: %+v", ver)
		return nil
	}

	rootCmd.SilenceErrors = true
	rootCmd.AddCommand(NewCmdVersion(out))
	rootCmd.AddCommand(NewCmdRun(out))
	rootCmd.PersistentFlags().StringVarP(&v, "verbosity", "v", constants.DefaultLogLevel.String(), "Log level (debug, info, warn, error, fatal, panic)")
	return rootCmd
}

func SetUpLogs(out io.Writer, level string) error {
	logrus.SetOutput(out)
	lvl, err := logrus.ParseLevel(v)
	if err != nil {
		return errors.Wrap(err, "parsing log level")
	}
	logrus.SetLevel(lvl)
	return nil
}
