ORG := matt-rickard.com
PROJECT := virgo

BUILD_PACKAGE = $(ORG)/$(PROJECT)/cmd/$(PROJECT)

BUILD_DIR ?= out
REPOSITORY ?= $(ORG)/$(PROJECT)
DATE := $(shell date +'%Y-%m-%dT%H:%M:%SZ')

DOCKER_REPO ?= gcr.io/matt-rickard/

GOOS ?= $(shell go env GOOS)
GOARCH ?= amd64
SUPPORTED_PLATFORMS := linux-$(ARCH) darwin-$(GOARCH) windows-$(GOARCH).exe
GOBIN ?= $(shell go env GOBIN)

GIT_COMMIT ?= $(shell git rev-parse HEAD)
GIT_TREE_STATE ?= $(if $(shell git status --porcelain),dirty-$(shell git diff | shasum | head -c 10),clean)

VERSION_PACKAGE = $(REPOSITORY)/pkg/version
MAJOR_VERSION = 0
MINOR_VERSION = 0
BUILD_VERSION = 1
VERSION ?= $(MAJOR_VERSION).$(MINOR_VERSION).$(BUILD_VERSION)

PARSER_PKG = pkg/parser

GO_LDFLAGS :="
GO_LDFLAGS += -extldflags \"$(LDFLAGS)\"
GO_LDFLAGS += -X $(VERSION_PACKAGE).project=$(PROJECT)
GO_LDFLAGS += -X $(VERSION_PACKAGE).version=$(VERSION)
GO_LDFLAGS += -X $(VERSION_PACKAGE).buildDate=$(DATE)
GO_LDFLAGS += -X $(VERSION_PACKAGE).gitCommit=$(GIT_COMMIT)
GO_LDFLAGS += -X $(VERSION_PACKAGE).gitTreeState=$(GIT_TREE_STATE)
GO_LDFLAGS +="

GO_GCFLAGS := "all=-trimpath=$(PWD)"
GO_ASMFLAGS := "all=-trimpath=$(PWD)"
GO_CGO_ENABLED = 0
GO_BUILD_TAGS := ""

GO_FILES := $(PARSER_PKG)/lex.yy.go $(PARSER_PKG)/y.go $(shell find . -type f -name '*.go' -not -path "./vendor/*") 

$(GOBIN)/$(PROJECT): $(GO_FILES)
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(GO_CGO_ENABLED) go install -ldflags $(GO_LDFLAGS) -gcflags $(GO_GCFLAGS) -asmflags $(GO_ASMFLAGS) -tags $(GO_BUILD_TAGS) $(BUILD_PACKAGE)

$(PARSER_PKG)/y.go: $(PARSER_PKG)/parser.y
	goyacc -o $(PARSER_PKG)/y.go $(PARSER_PKG)/parser.y

$(PARSER_PKG)/lex.yy.go: $(PARSER_PKG)/lex.l
	golex -o $(PARSER_PKG)/lex.yy.go $(PARSER_PKG)/lex.l

%.sha256: %
	shasum -a 256 $< > $@

%.exe: %
	cp $< $@

.PHONY: release
release: $(BUILD_DIR) build-linux build-windows build-darwin
	mv $(BUILD_DIR)/virgo-windows $(BUILD_DIR)/virgo.exe

build-%:
	GOOS=$* GOARCH=$(GOARCH) CGO_ENABLED=$(GO_CGO_ENABLED) go build -o $(BUILD_DIR)/virgo-$* -ldflags $(GO_LDFLAGS) -gcflags $(GO_GCFLAGS) -asmflags $(GO_ASMFLAGS) -tags $(GO_BUILD_TAGS) $(BUILD_PACKAGE)

.PHONY: test
test: $(GOBIN)/$(PROJECT)
	@ hack/test.sh

.PHONY: coverage
coverage: $(BUILD_DIR)
	go test -coverprofile=$(BUILD_DIR)/coverage.txt -covermode=atomic ./...

.PHONY: clean
clean: $(BUILD_DIR)
	rm -r $(BUILD_DIR) $(PARSER_PKG)/y.go $(PARSER_PKG)/lex.yy.go

$(BUILD_DIR):
	mkdir $(BUILD_DIR)