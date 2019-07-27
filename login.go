package main

import (
	"github.com/orange-cloudfoundry/cfm/messages"
	"os"
)

type LoginCmd struct {
}

var loginCmd LoginCmd

func (c *LoginCmd) Execute(_ []string) error {
	targets := findTargets()
	if len(targets) == 0 {
		messages.Fatal("please set targets with command set-targets: cfm set-targets https://my.cf.1.com https://my.cf.2.com")
	}

	for _, target := range targets {
		runCfCommand(os.Args[1:], convertToFolder(target), os.Stdin)
	}
	return nil
}

func init() {
	desc := `Login`
	_, err := parser.AddCommand(
		"login",
		desc,
		desc,
		&multiLoginCmd)
	if err != nil {
		panic(err)
	}
}
