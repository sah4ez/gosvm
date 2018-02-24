package main

import (
	"fmt"
	"os"

	structure "github.com/sah4ez/gosvm/internal/structure"
)

type listCmd struct{}

var shortHelp = "print list all services from project"

func (l *listCmd) Name() string      { return "list" }
func (l *listCmd) ShortHelp() string { return shortHelp }
func (l *listCmd) Enable() bool      { return true }
func (l *listCmd) Run(args []string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	switch len(args) {
	case 0:
		root, err := structure.LoadStructure(wd + "/svm.toml")
		if err != nil {
			return err
		}

		fmt.Fprintln(os.Stdout, "Title:\t\t", root.Title)
		fmt.Fprintln(os.Stdout, "Description:\t", root.Description)
		fmt.Fprintln(os.Stdout, "Version:\t", root.Description)
		for i, sub := range root.SubProject {
			if i == 0 {
				fmt.Fprintln(os.Stdout, "\t")
				fmt.Fprintln(os.Stdout, "\tSubPackages:")
				fmt.Fprintln(os.Stdout, "\t")
			}
			fmt.Fprintf(os.Stdout, "\t%s@%s\n", sub.Title, sub.Version)
		}

	default:
		fmt.Fprintln(os.Stdout, "stub for list.")
		fmt.Fprintln(os.Stdout, "Project Name: <NAME>")
		fmt.Fprintln(os.Stdout, "\t Sub: <NAME>@<VERSION>")
		fmt.Fprintln(os.Stdout, "\t Sub: <NAME>@<VERSION>")
	}
	return nil
}
