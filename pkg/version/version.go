package version

import (
	"fmt"
	"runtime"

	"github.com/blang/semver"
	"github.com/pkg/errors"
)

var project, version, gitCommit, gitTreeState, buildDate string
var platform = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)

type Info struct {
	Project      string
	Version      string
	GitCommit    string
	GitTreeState string
	BuildDate    string
	GoVersion    string
	Compiler     string
	Platform     string
}

// Get returns the version and buildtime information about the binary
func Get() *Info {
	return &Info{
		Project:      project,
		Version:      version,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     platform,
	}
}

func ParseVersion(version string) (semver.Version, error) {
	v, err := semver.Parse(version)
	if err != nil {
		return semver.Version{}, errors.Wrap(err, "parsing semver")
	}
	return v, nil
}
