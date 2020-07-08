set -e -o pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

if ! [ -x "$(command -v golangci-lint)" ]; then
	echo "Installing GolangCI-Lint"
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.27.0
fi

golangci-lint run \
	--no-config \
    --skip-dirs vendor \
	-E goconst \
	-E goimports \
	-E golint \
	-E interfacer \
	-E maligned \
	-E misspell \
	-E unconvert \
	-E unparam \
	-E errcheck \