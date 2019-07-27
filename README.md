# cfm cli

cfm is the small utility which wrap [cloud foundry cli](https://github.com/cloudfoundry/cli) to be able to run same command on 
multiple cloud foundry.

## Installation

### On *nix system

You can install this via the command-line with either `curl` or `wget`.

#### via curl

```bash
$ sh -c "$(curl -fsSL https://raw.github.com/orange-cloudfoundry/cfm/master/bin/install.sh)"
```

#### via wget

```bash
$ sh -c "$(wget https://raw.github.com/orange-cloudfoundry/cfm/master/bin/install.sh -O -)"
```

### On windows

You can install it by downloading the `.exe` corresponding to your cpu from releases page: https://github.com/orange-cloudfoundry/cfm/releases .
Alternatively, if you have terminal interpreting shell you can also use command line script above, it will download file in your current working dir.

### From go command line

Simply run in terminal:

```bash
$ go get github.com/orange-cloudfoundry/cfm
```

## Usage

### prerequisite

You must set target with `add-target` for each cloud foundry:

```
cfm [OPTIONS] add-target [add-target-OPTIONS]

Add cloud foundry target

Application Options:
  -h, --help

Help Options:
  -h, --help        Show this help message

[add-target command options]
      -t, --target= Cloud foundry target
      -a, --alias=  set an alias to the target
```

Example: `cfm add-target -t https://api.my.cloudfoundry.com -a mycf`

### Run cf command

simply call cfm with cf valid arguments and the command will be called on each targets, example for see each orgs on targets:

```
$ cfm orgs
```

### Target a particular cloud foundry

You may want to do commands for only one targets and not re-login through normal cli, to do so you can use this command:


```
cfm [OPTIONS] set-cf-env [set-cf-env-OPTIONS]

set your cf normal cli to targeted cloud foundry

Application Options:
  -h, --help

Help Options:
  -h, --help        Show this help message

[set-cf-env command options]
      -t, --target= Cloud foundry target
      -a, --alias=  set an alias to the target
```

Example: `cfm set-cf-env -a mycf`

You can now perform command with normal `cf` cli.