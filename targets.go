package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Target struct {
	Api   string
	Alias string
}

func (t Target) String() string {
	if t.Alias != "" {
		return t.Alias
	}
	return t.Api
}

const configFile = "config.json"

type TargetCmd struct {
	Targets []string `positional-args:"true" description:"Cloud foundry target"`
}

var targetCmd TargetCmd

func (c *TargetCmd) Execute(_ []string) error {

	return storeTargetsStr(c.Targets)
}

func init() {
	desc := `Set cloud foundry targets`
	_, err := parser.AddCommand(
		"set-targets",
		desc,
		desc,
		&targetCmd)
	if err != nil {
		panic(err)
	}
}

func convertToFolder(target Target) string {
	targetApi := strings.Replace(target.Api, ":", "_", -1)
	targetApi = strings.Replace(targetApi, "/", "_", -1)
	targetApi = strings.Replace(targetApi, ".", "_", -1)
	return filepath.Join(cfmHome(), targetApi)
}

func findTargets() []Target {
	b, err := ioutil.ReadFile(filepath.Join(cfmHome(), configFile))
	if err != nil {
		return []Target{}
	}
	var targets []Target

	err = json.Unmarshal(b, &targets)
	if err != nil {
		return []Target{}
	}
	return targets
}

func storeTargets(targets []Target) error {
	b, err := json.MarshalIndent(targets, "", "\t")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.Join(cfmHome(), configFile), b, 0644)
	if err != nil {
		return err
	}
	for _, target := range targets {
		err = os.MkdirAll(convertToFolder(target), 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func storeTargetsStr(targetsStr []string) error {
	targets := make([]Target, len(targetsStr))
	for i, t := range targetsStr {
		targets[i] = Target{
			Api: t,
		}
	}
	return storeTargets(targets)
}
