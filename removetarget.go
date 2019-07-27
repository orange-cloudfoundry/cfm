package main

type RmTargetCmd struct {
	Target string `positional-args:"1" description:"Cloud foundry target"`
	Alias  string `long:"alias" short:"a" description:"set an alias to the target"`
}

var rmTargetCmd RmTargetCmd

func (c *RmTargetCmd) Execute(_ []string) error {
	targets := findTargets()
	finalTargets := make([]Target, 0)
	for _, t := range targets {
		if t.Api == c.Target || t.Alias == c.Alias {
			continue
		}
		finalTargets = append(finalTargets, t)
	}

	return storeTargets(finalTargets)
}

func init() {
	desc := `Add cloud foundry target`
	_, err := parser.AddCommand(
		"remove-target",
		desc,
		desc,
		&rmTargetCmd)
	if err != nil {
		panic(err)
	}
}
