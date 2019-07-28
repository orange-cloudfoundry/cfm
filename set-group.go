package main

type SetGroupCmd struct {
	Group string `long:"group" short:"g" description:"set the group of targets to use"`
}

var setGroupCmd SetGroupCmd

func (c *SetGroupCmd) Execute(_ []string) error {
	return setGroup(c.Group)
}

func init() {
	desc := `Add cloud foundry target`
	_, err := parser.AddCommand(
		"set-group",
		desc,
		desc,
		&setGroupCmd)
	if err != nil {
		panic(err)
	}
}
