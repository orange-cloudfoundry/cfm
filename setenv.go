package main

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/orange-cloudfoundry/cfm/messages"
	"io/ioutil"
	"path/filepath"
)

type SetEnvCmd struct {
	Target string `long:"target" short:"t" description:"Cloud foundry target"`
	Alias  string `long:"alias" short:"a" description:"set an alias to the target"`
}

var setEnvCmd SetEnvCmd

func (c *SetEnvCmd) Execute(_ []string) error {
	targets := findTargets()
	var target Target
	for _, t := range targets {
		if t.Api == c.Target || t.Alias == c.Alias {
			target = t
			break
		}
	}

	if target.Api == "" {
		return fmt.Errorf("Target not found")
	}
	b, err := ioutil.ReadFile(filepath.Join(convertToFolder(target), ".cf", "config.json"))
	if err != nil {
		return err
	}
	return ioutil.WriteFile(cfConfFile(), b, 0644)
}

func init() {
	desc := `set your cf normal cli to targeted cloud foundry`
	_, err := parser.AddCommand(
		"set-cf-env",
		desc,
		desc,
		&setEnvCmd)
	if err != nil {
		panic(err)
	}
}

func cfConfFile() string {
	h, err := homedir.Dir()
	if err != nil {
		messages.Fatal(err.Error())
	}
	return filepath.Join(h, ".cf", "config.json")
}
