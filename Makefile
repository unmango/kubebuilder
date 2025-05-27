GO     ?= go
DEVCTL ?= $(GO) tool devctl

GO_SRC != $(DEVCTL) list --go

build: bin/kubebuilder
tidy: go.sum

bin/kubebuilder: go.mod ${GO_SRC}
	$(GO) build -o $@

go.sum: go.mod ${GO_SRC}
	$(GO) mod tidy

ENVRC ?= example

.envrc: hack/${ENVRC}.envrc
	cp $< $@ && chmod a=,u+r $@
