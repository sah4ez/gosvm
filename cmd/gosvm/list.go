package main // import "github.com/sah4ez/gosvm/cmd/gosvm"

import (
	"io"
	"os"

	fmt "github.com/sah4ez/gosvm/pkg/formatting"
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

		fmt.Info.Fprintln(l.w, "Title:\t\t", root.Title)
		fmt.Info.Fprintln(l.w, "Description:\t", root.Description)
		fmt.Info.Fprintln(l.w, "Version:\t", root.Version)
		for i, sub := range root.SubProject {
			if i == 0 {
				fmt.Info.Fprintln(l.w, "\t")
				fmt.Info.Fprintln(l.w, "\tSubPackages:")
				fmt.Info.Fprintln(l.w, "\t")
			}
			fmt.Info.Fprint(l.w, "\t", sub.Title)
			if sub.Version != "" {
				fmt.Info.Fprint(l.w, "@", sub.Version)
			}
			fmt.Info.Fprint(l.w, "\n")
		}

	default:
		fmt.Info.Fprintf(l.w, "for args %s\n", args)
		fmt.Info.Fprintln(l.w, "stub for list.")
		fmt.Info.Fprintln(l.w, "Project Name: <NAME>")
		fmt.Info.Fprintln(l.w, "\t Sub: <NAME>@<VERSION>")
		fmt.Info.Fprintln(l.w, "\t Sub: <NAME>@<VERSION>")
	}
	return nil
}
