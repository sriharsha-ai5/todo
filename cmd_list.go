package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
)

const (
	doneMark1 = "\u2610"
	doneMark2 = "\u2611"
)

func makeCmdList(filename string) *commander.Command {
	cmdList := func(cmd *commander.Command, args []string) error {
		nflag := cmd.Flag.Lookup("n").Value.Get().(bool)
		byt, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println("Error in reading whole file at once")
			return err

		}
		sstrspl := strings.Split(string(byt), "\n")
		sstrspl = sstrspl[:len(sstrspl)-1]
		for i, v := range sstrspl {
			n := i + 1
			line := v
			if strings.HasPrefix(line, "-") {
				if !nflag {
					fmt.Printf("%s %03d: %s\n", doneMark2, n, strings.TrimSpace(line[1:]))
				}
			} else {
				fmt.Printf("%s %03d: %s\n", doneMark1, n, strings.TrimSpace(line))
			}

		}
		return nil
	}

	flg := *flag.NewFlagSet("list", flag.ExitOnError)
	flg.Bool("n", false, "only not done")

	return &commander.Command{
		Run:       cmdList,
		UsageLine: "list [options]",
		Short:     "show list index",
		Flag:      flg,
	}
}
