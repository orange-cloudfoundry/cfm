package main

type HelpCmd struct {
}

var helpCmd HelpCmd

func (c *HelpCmd) Execute(_ []string) error {
	help()
	return nil
}

func init() {
	desc := `Help`
	_, err := parser.AddCommand(
		"help",
		desc,
		desc,
		&HelpCmd{})
	if err != nil {
		panic(err)
	}
}
