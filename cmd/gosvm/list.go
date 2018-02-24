package main

import (
	"fmt"
	"os"
)

type listCmd struct{}

var shortHelp = "print list all services from project"

func (l *listCmd) Name() string      { return "list" }
func (l *listCmd) ShortHelp() string { return shortHelp }
func (l *listCmd) Enable() bool      { return true }
func (l *listCmd) Run(args []string) error {
	fmt.Fprintln(os.Stdout, "stub for list.")
	fmt.Fprintln(os.Stdout, "Project Name: <NAME>")
	fmt.Fprintln(os.Stdout, "\t Sub: <NAME>@<VERSION>")
	fmt.Fprintln(os.Stdout, "\t Sub: <NAME>@<VERSION>")
	return nil
}
