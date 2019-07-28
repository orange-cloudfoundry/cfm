package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

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

func findAllTargets() []Target {
	b, err := ioutil.ReadFile(filepath.Join(cfmHome(), configFile))
	if err != nil {
		return []Target{}
	}
	var config Config

	err = json.Unmarshal(b, &config)
	if err != nil {
		return []Target{}
	}
	return config.Targets
}

func findTargets() []Target {
	targets := findAllTargets()
	group := getGroup()
	finalTargets := make([]Target, 0)
	for _, t := range targets {
		if t.Group == group {
			finalTargets = append(finalTargets, t)
		}
	}
	return finalTargets
}

func getGroup() string {
	b, err := ioutil.ReadFile(filepath.Join(cfmHome(), configFile))
	if err != nil {
		return ""
	}
	var config Config

	err = json.Unmarshal(b, &config)
	if err != nil {
		return ""
	}

	return config.CurrentGroup
}

func setGroup(group string) error {
	b, err := ioutil.ReadFile(filepath.Join(cfmHome(), configFile))
	if err != nil {
		return err
	}
	var config Config

	err = json.Unmarshal(b, &config)
	if err != nil {
		return err
	}

	b, err = json.MarshalIndent(Config{
		CurrentGroup: group,
		Targets:      config.Targets,
	}, "", "\t")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.Join(cfmHome(), configFile), b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func storeTargets(targets []Target) error {
	b, err := json.MarshalIndent(Config{
		CurrentGroup: getGroup(),
		Targets:      targets,
	}, "", "\t")
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
