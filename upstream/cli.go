package upstream

import (
	"github.com/charmbracelet/log"
	"sigs.k8s.io/kubebuilder/v4/pkg/cli"
)

func NewCli() (*cli.CLI, error) {
	if options, err := CliOptions(); err != nil {
		return nil, err
	} else {
		return cli.New(options)
	}
}

func Run() {
	cli, err := NewCli()
	if err != nil {
		log.Fatal(err)
	}
	if err = cli.Run(); err != nil {
		log.Fatal(err)
	}
}
