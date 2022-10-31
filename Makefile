BIN_DIR = .
PROTO_DIR = grpc-common
GRPC_OUT = grpc
HELP_CMD = grep -E '^[a-zA-Z_-]+:.*?\#\# .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?\#\# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


.DEFAULT_GOAL := help
.PHONY: openhms
project := openhms

all: $(project) ## Generate Pbs and build
openhms: $@ ## Build the api binary

$(project):
	go build


test: all ## Launch tests
	go test ./...

bump: all ## Update packages version
	go get -u ./...

help: ## Show this help
	@${HELP_CMD}
