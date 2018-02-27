package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/sah4ez/gosvm/internal/dependency"
	"github.com/sah4ez/gosvm/internal/dependency/dep"
	"github.com/sah4ez/gosvm/internal/dependency/glide"
	"github.com/sah4ez/gosvm/internal/structure"
)

type libsCmd struct{}

var shortHelpLibs = "list using libs in SubProject and hist version"

func (l *libsCmd) Name() string      { return "libs" }
func (l *libsCmd) ShortHelp() string { return shortHelpLibs }
func (l *libsCmd) Enable() bool      { return true }
func (l *libsCmd) Run(args []string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	stdout := bytes.NewBuffer([]byte{})

	switch len(args) {
	case 0:
		root, err := structure.LoadStructure(wd + "/svm.toml")
		if err != nil {
			return err
		}
		packs := dependency.NewPackages()

		// TODO: need implement loader information per package instead of using all subpackages.
		glideLoader := glide.NewGlideLoader(root, packs)
		packs, err = glideLoader.Load()
		if err != nil {
			return err
		}
		depLoader := dep.NewDepLoader(root, packs)
		packs, err = depLoader.Load()
		if err != nil {
			return err
		}

		fmt.Fprintln(stdout, "Title:\t\t", root.Title)
		fmt.Fprintln(stdout, "Description:\t", root.Description)
		fmt.Fprintln(stdout, "Version:\t", root.Version)
		fmt.Fprintln(stdout, "\n====================Libs====================\n")

		packs.Range(func(libName string, packVersion map[string][]string) {
			fmt.Fprintln(stdout, "\t", libName)
			if len(packVersion) > 1 {
				fmt.Fprintf(stdout, "\t !!!%d differnt vesrion are used!!!\n", len(packVersion))
			}
			for ver, pack := range packVersion {
				if ver == "" {
					fmt.Fprintf(stdout, "\t\tlatest :")
				} else {
					fmt.Fprintf(stdout, "\t\t%.12s :", ver)
				}

				fmt.Fprintln(stdout)
				for _, name := range pack {
					fmt.Fprintf(stdout, "\t\t\t%s\n", name)
				}
			}
			fmt.Fprintln(stdout)
		})
	default:
		fmt.Fprintln(os.Stderr, "wrong args")
	}
	defer func(out bytes.Buffer) {
		fmt.Fprintln(os.Stdout, out.String())
	}(*stdout)
	less(stdout)
	return nil
}

func less(stdin io.Reader) {
	less := exec.Command("less")
	less.Stdin = stdin
	less.Stdout = os.Stdout
	less.Start()
	less.Wait()
}