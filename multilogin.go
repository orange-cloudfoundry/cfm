package main

import (
	"bytes"
	"github.com/orange-cloudfoundry/cfm/messages"
)

type MultiLoginCmd struct {
	Apis              []string `short:"a" description:"api"`
	Username          string   `short:"u" description:"username"`
	Password          string   `short:"p" description:"password"`
	Org               string   `short:"o" description:"org"`
	Space             string   `short:"s" description:"s"`
	SkipSslValidation bool     `long:"skip-ssl-validation" description:"skip-ssl-validation"`
}

var multiLoginCmd MultiLoginCmd

func (c *MultiLoginCmd) Execute(_ []string) error {
	if len(c.Apis) > 0 {
		err := storeTargetsStr(c.Apis)
		if err != nil {
			return err
		}
	}
	targets := findTargets()
	if len(targets) == 0 {
		messages.Fatal("please set targets with command set-targets: cfm set-targets https://my.cf.1.com https://my.cf.2.com")
	}
	targetSet := false
	if c.Org != "" || c.Space != "" {
		targetSet = true
	}
	for _, target := range targets {
		var args []string
		if targetSet {
			args = []string{
				"login",
				"-a", target.Api,
				"-u", c.Username,
				"-p", c.Password,
			}
			if c.Org != "" {
				args = append(args, "-o", c.Org)
			}
			if c.Space != "" {
				args = append(args, "-s", c.Space)
			}
			if c.SkipSslValidation {
				args = append(args, "--skip-ssl-validation")
			}
			runCfCommand(args, convertToFolder(target), &bytes.Buffer{})
			continue
		}
		args = []string{"api", target.Api}
		if c.SkipSslValidation {
			args = append(args, "--skip-ssl-validation")
		}
		runCfCommand(args, convertToFolder(target), &bytes.Buffer{})
		args = []string{"auth", c.Username, c.Password}
		runCfCommand(args, convertToFolder(target), &bytes.Buffer{})
	}
	return nil
}

func init() {
	desc := `perform login on all targets`
	_, err := parser.AddCommand(
		"multi-login",
		desc,
		desc,
		&multiLoginCmd)
	if err != nil {
		panic(err)
	}
}
