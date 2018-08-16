GOCMD=go
BUILDDIR=build
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
SRV_BIN_NAME=telnetgod
SRV_BIN_LINUX=$(SRV_BIN_NAME)_linux
SRV_BIN_DARWIN=$(SRV_BIN_NAME)_darwin

GREEN=\033[32;01m
NOCO=\033[0m

help:
	@echo "$(GREEN)Commands: $(NOCO)"
	@echo "    deps $(NOCO) [OPTIONS]"
	@echo "    build $(NOCO) [OPTIONS]"
	@echo "    all $(NOCO) [OPTIONS]"

deps:
	@echo "$(GREEN)==>$(NOCO) Installing Dependencies"
	@make deps-server

deps-server:
	@echo "$(GREEN)==>$(NOCO) Installing Server Dependencies"
	@$(GOGET) github.com/reiver/go-oi
	@$(GOGET) github.com/reiver/go-telnet
	@$(GOGET) github.com/reiver/go-telnet/telsh

serve:
	@make build
	@echo "$(GREEN)==>$(NOCO) Running Server"
	@./$(BUILDDIR)/$(SRV_BIN_NAME)

build:
	@make clean
	@make build-server

build-server:
	@echo "$(GREEN)==>$(NOCO) Building Server"
	@$(GOBUILD) -o $(BUILDDIR)/$(SRV_BIN_NAME) server/server.go

build-server-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BUILDDIR)/$(SRV_BIN_LINUX) server/server.go

build-server-darwin:
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BUILDDIR)/$(SRV_BIN_DARWIN) server/server.go

clean:
	@$(GOCLEAN)
	@make clean-server

clean-server:
	@echo "$(GREEN)==>$(NOCO) Cleaning Go"
	@$(GOCLEAN)
	@echo "$(GREEN)==>$(NOCO) Cleaning $(BUILDDIR)/$(SRV_BIN_NAME)* "
	@[ -z "$(ls -A $(BUILDDIR))" ] && rm -rf $(BUILDDIR)/$(SRV_BIN_NAME)* || true

all:
	@make deps
	@make build

.PHONY: all deps deps-server build build-server clean clean-server
