package main

import (
	"errors"
	"fmt"
	"os"
)

type setCmd struct{}

var shortHelpSet = "set version for all or cpecifict pacakge\n \tset <LIB_NAME><VERSION> [...<PACK_NAME>]"

func (s *setCmd) Name() string      { return "set" }
func (s *setCmd) ShortHelp() string { return shortHelpSet }
func (s *setCmd) Enable() bool      { return true }
func (s *setCmd) Run(args []string) error {
	switch len(args) {
	case 0, 1, 2:
		return errors.New(fmt.Sprintf("wrong nubmer args: %d", len(args)))
	case 3:
		fmt.Fprintf(os.Stdout, "will be set %s for all libs %s\n", args[2], args[1])
	default:
		fmt.Fprintf(os.Stdout, "will be set %s for all libs %s in packs: %s\n", args[2], args[1], args[3:])
	}
	return nil
}
