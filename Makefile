_ != mkdir -p .make

GO     ?= go
DEVCTL ?= $(GO) tool devctl
GINKGO ?= $(GO) tool ginkgo
GOLINT ?= $(GO) tool golangci-lint

GO_SRC != $(DEVCTL) list --go

##@ Primary Targets

build: bin/kubebuilder
test: .make/ginkgo-run
format fmt: .make/go-fmt
lint: .make/go-vet .make/golangci-lint
tidy: go.sum

##@ Artifacts

bin/kubebuilder: go.mod ${GO_SRC}
	$(GO) build -o $@

go.sum: go.mod ${GO_SRC}
	$(GO) mod tidy

##@ Development environment

%_suite_test.go: ## Bootstrap a Ginkgo test suite
	cd $(dir $@) && $(GINKGO) bootstrap
%_test.go: ## Generate a Ginkgo test
	cd $(dir $@) && $(GINKGO) generate $(notdir $@)

ENVRC ?= example

.envrc: hack/${ENVRC}.envrc
	cp $< $@ && chmod a=,u+r $@

##@ Sentinels

.make/ginkgo-run: ${GO_SRC}
	$(GINKGO) $(sort $(dir $?))
	@touch $@

.make/go-fmt: ${GO_SRC}
	$(GO) fmt $(addprefix ./,$(sort $(dir $?)))
	@touch $@

.make/go-vet: ${GO_SRC}
	$(GO) vet $(addprefix ./,$(sort $(dir $?)))
	@touch $@

.make/golangci-lint: ${GO_SRC}
	$(GOLINT) run
	@touch $@

##@ Local tool binaries

export GOBIN := ${CURDIR}/bin

bin/devctl: go.mod ## Optional bin install
	$(GO) install github.com/unmango/devctl

bin/ginkgo: go.mod ## Optional bin install
	$(GO) install github.com/onsi/ginkgo/v2/ginkgo

bin/golangci-lint: go.mod ## Optional bin install
	$(GO) install github.com/golangci/golangci-lint/v2/cmd/golangci-lint
