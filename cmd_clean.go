package main

import (
	"os"
	"strings"

	"github.com/gonuts/commander"
	fp "github.com/repeale/fp-go"
)

func makeCmdClean(filename string) *commander.Command {
	cmdClean := func(cmd *commander.Command, args []string) error {
		if len(args) != 0 {
			cmd.Usage()
			return nil
		}
		byt, err := os.ReadFile(filename)
		strsplt := strings.Split(string(byt), "\n")
		err = os.WriteFile(filename, []byte(strings.Join(fp.Filter(func(v string) bool {
			return !strings.HasPrefix(v, "-")
		})(strsplt), "\n")), 0644)
		return err
	}

	return &commander.Command{
		Run:       cmdClean,
		UsageLine: "clean",
		Short:     "remove all done items",
	}
}
