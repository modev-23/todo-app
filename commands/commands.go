package commands

import "flag"

// Command struct with updated names to export
type Command struct {
	Name        string
	Description string
	flags       *flag.FlagSet
	Execute     func(cmd *Command, args []string)
}

func (c *Command) Init(args []string) error {
	return c.flags.Parse(args)
}

func (c *Command) Called() bool {
	return c.flags.Parsed()
}

func (c *Command) Run() {
	c.Execute(c, c.flags.Args())
}
