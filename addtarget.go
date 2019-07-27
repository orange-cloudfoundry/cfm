package main

type AddTargetCmd struct {
	Target string `long:"target" short:"t" required:"1" description:"Cloud foundry target"`
	Alias  string `long:"alias" short:"a" description:"set an alias to the target"`
}

var addTargetCmd AddTargetCmd

func (c *AddTargetCmd) Execute(_ []string) error {
	targets := findTargets()
	targets = append(targets, Target{
		Api:   c.Target,
		Alias: c.Alias,
	})

	return storeTargets(targets)
}

func init() {
	desc := `Add cloud foundry target`
	_, err := parser.AddCommand(
		"add-target",
		desc,
		desc,
		&addTargetCmd)
	if err != nil {
		panic(err)
	}
}
