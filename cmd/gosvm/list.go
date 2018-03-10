package main

import (
	"fmt"
	"os"

	structure "github.com/sah4ez/gosvm/internal/structure"
)

type listCmd struct{}

var shortHelpList = "print list all services from project"

func (l *listCmd) Name() string      { return "list" }
func (l *listCmd) ShortHelp() string { return shortHelpList }
func (l *listCmd) Enable() bool      { return true }
func (l *listCmd) Run(args []string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	switch len(args) {
	case 1:
		root, err := structure.LoadStructure(wd + "/svm.toml")
		if err != nil {
			return err
		}

		fmt.Fprintln(os.Stdout, "Title:\t\t", root.Title)
		fmt.Fprintln(os.Stdout, "Description:\t", root.Description)
		fmt.Fprintln(os.Stdout, "Version:\t", root.Version)
		for i, sub := range root.SubProject {
			if i == 0 {
				fmt.Fprintln(os.Stdout, "\t")
				fmt.Fprintln(os.Stdout, "\tSubPackages:")
				fmt.Fprintln(os.Stdout, "\t")
			}
			fmt.Fprint(os.Stdout, "\t", sub.Title)
			if sub.Version != "" {
				fmt.Fprint(os.Stdout, "@", sub.Version)
			}
			fmt.Fprint(os.Stdout, "\n")
		}

	default:
		fmt.Fprintf(os.Stdout, "for args %s\n", args)
		fmt.Fprintln(os.Stdout, "stub for list.")
		fmt.Fprintln(os.Stdout, "Project Name: <NAME>")
		fmt.Fprintln(os.Stdout, "\t Sub: <NAME>@<VERSION>")
		fmt.Fprintln(os.Stdout, "\t Sub: <NAME>@<VERSION>")
	}
	return nil
}
