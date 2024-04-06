package fix

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func (f *Fix) action(_ *cli.Context) error {
	return f.fixPrompt()
}

func (f *Fix) before(_ *cli.Context) error {
	fmt.Println("fix -> before")
	return nil
}

func (f *Fix) Command() *cli.Command {
	return &cli.Command{
		Name:    "fix",
		Usage:   "Fixes a failed command",
		Aliases: []string{"f"},
		Action:  f.action,
		Before:  f.before,
	}
}
