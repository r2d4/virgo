package main

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"matt-rickard.com/virgo/cmd/virgo/app"
)

func main() {
	if err := app.Run(); err != nil {
		if errors.Cause(err) == context.Canceled {
			logrus.Debugln(errors.Wrap(err, "ignore error since context is cancelled"))
		} else {
			logrus.Fatal(err)
		}
	}
}
