package main // import "github.com/sah4ez/gosvm/cmd/gosvm"

import (
	"fmt"
	"io"
	"os"

	"github.com/sah4ez/gosvm/pkg/structure"
)

type listCmd struct {
	w io.Writer
}

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

		fmt.Fprintln(l.w, "Title:\t\t", root.Title)
		fmt.Fprintln(l.w, "Description:\t", root.Description)
		fmt.Fprintln(l.w, "Version:\t", root.Version)
		for i, sub := range root.SubProject {
			if i == 0 {
				fmt.Fprintln(l.w, "\t")
				fmt.Fprintln(l.w, "\tSubPackages:")
				fmt.Fprintln(l.w, "\t")
			}
			fmt.Fprint(l.w, "\t", sub.Title)
			if sub.Version != "" {
				fmt.Fprint(l.w, "@", sub.Version)
			}
			fmt.Fprint(l.w, "\n")
		}

	default:
		fmt.Fprintf(l.w, "for args %s\n", args)
		fmt.Fprintln(l.w, "stub for list.")
		fmt.Fprintln(l.w, "Project Name: <NAME>")
		fmt.Fprintln(l.w, "\t Sub: <NAME>@<VERSION>")
		fmt.Fprintln(l.w, "\t Sub: <NAME>@<VERSION>")
	}
	return nil
}
