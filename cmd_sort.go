package main

import (
	"os"
	"strings"

	"github.com/gonuts/commander"
	fp "github.com/repeale/fp-go"
)

func makeCmdSort(filename string) *commander.Command {
	cmdSort := func(cmd *commander.Command, args []string) error {
		if len(args) != 0 {
			cmd.Usage()
			return nil
		}
		byt, err := os.ReadFile(filename)
		strsplt := strings.Split(string(byt), "\n")
		err = os.WriteFile(filename, []byte(strings.Join(append(fp.Filter(func(v string) bool {
			return strings.HasPrefix(v, "-")
		})(strsplt), fp.Filter(func(v string) bool {
			return !strings.HasPrefix(v, "-")
		})(strsplt)...), "\n")), 0644)
		return err
	}

	return &commander.Command{
		Run:       cmdSort,
		UsageLine: "sort",
		Short:     "sorts done to the top and undone to the bottom",
	}
}
