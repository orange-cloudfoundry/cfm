package main

import (
	"bytes"
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/mitchellh/go-homedir"
	"github.com/orange-cloudfoundry/cfm/messages"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

const cfmHomeDir = ".cfm"

type Options struct {
	Help func() `long:"help" short:"h" description:""`
}

var options Options

var parser = flags.NewParser(&options, flags.HelpFlag|flags.PassDoubleDash|flags.IgnoreUnknown)

func help() {
	cmd := exec.Command("cf", os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func main() {
	options.Help = help
	os.MkdirAll(cfmHome(), 0755)
	_, err := exec.LookPath("cf")
	if err != nil {
		messages.Fatal("Can't found cf cli please install before: https://github.com/cloudfoundry/cli")
	}
	_, err = parser.ParseArgs(os.Args[1:])
	if err == nil {
		return
	}
	if err != nil {
		errFlag, ok := err.(*flags.Error)
		if !ok || errFlag.Type != flags.ErrUnknownCommand && errFlag.Type != flags.ErrUnknownFlag {
			if os.Args[1] == "--help" || os.Args[1] == "-h" {
				help()
				fmt.Println(err.Error())
				return
			}
			messages.Fatal(err.Error())
		}
	}

	if os.Args[1] == "--help" || os.Args[1] == "-h" {
		return
	}
	targets := findTargets()
	if len(targets) == 0 {
		messages.Fatal("please set targets with command set-targets: cfm set-targets https://my.cf.1.com https://my.cf.2.com")
	}
	stdin := createStdin()
	inError := false
	for _, target := range targets {
		messages.Printf("Runnning command on cf api '%s':\n", messages.C.Cyan(target))
		err := runCfCommand(os.Args[1:], convertToFolder(target), stdin)
		if err != nil {
			inError = true
		}
		messages.Println("-------")
		messages.Println("")
		if stdinReader, ok := stdin.(*bytes.Reader); ok {
			stdinReader.Seek(0, 0)
		}
	}
	if inError {
		os.Exit(1)
	}
}

func cfmHome() string {
	h, err := homedir.Dir()
	if err != nil {
		messages.Fatal(err.Error())
	}
	return filepath.Join(h, cfmHomeDir)
}

func createStdin() io.Reader {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return &bytes.Buffer{}
	}
	if stat.Size() == 0 {
		return &bytes.Buffer{}
	}
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return &bytes.Buffer{}
	}
	return bytes.NewReader(b)
}

func runCfCommand(args []string, configDir string, stdin io.Reader) error {
	cmd := exec.Command("cf", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = stdin
	env := append(os.Environ(), fmt.Sprintf("CF_HOME=%s", configDir))
	cmd.Env = env

	return cmd.Run()
}
