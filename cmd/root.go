package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/unmango/kubebuilder/pkg/cli"
)

func init() {
	log.SetReportTimestamp(false)
}

func Run() {
	cli, err := cli.New()
	if err != nil {
		log.Fatal(err)
	}
	if err := cli.Run(); err != nil {
		log.Fatal(err)
	}
}
