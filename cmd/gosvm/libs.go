package main // import "github.com/sah4ez/gosvm/cmd/gosvm"

import (
	"bytes"
	"io"
	"os"
	"os/exec"

	"github.com/sah4ez/gosvm/pkg/dependency"
	"github.com/sah4ez/gosvm/pkg/dependency/dep"
	"github.com/sah4ez/gosvm/pkg/dependency/glide"
	"github.com/sah4ez/gosvm/pkg/dependency/gomod"
	fmt "github.com/sah4ez/gosvm/pkg/formatting"
	"github.com/sah4ez/gosvm/pkg/structure"
)

type libsCmd struct {
	w io.Writer
}

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
	case 1:
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
		goModLoader := gomod.NewGoModLoader(root, packs)
		packs, err = goModLoader.Load()
		if err != nil {
			return err
		}

		fmt.Info.Fprintln(stdout, "Title:\t\t", root.Title)
		fmt.Info.Fprintln(stdout, "Description:\t", root.Description)
		fmt.Info.Fprintln(stdout, "Version:\t", root.Version)
		fmt.Info.Fprintln(stdout, "====================Libs====================")
		fmt.Info.Fprintln(stdout, "")

		packs.Range(func(libName string, packVersion map[string][]string) {
			fmt.Info.Fprintln(stdout, "\t", libName)
			if len(packVersion) > 1 {
				fmt.Info.Fprintf(stdout, "\t !!!%d differnt vesrion are used!!!\n", len(packVersion))
			}
			for ver, pack := range packVersion {
				if ver == "" {
					fmt.Info.Fprintf(stdout, "\t\tlatest :")
				} else {
					fmt.Info.Fprintf(stdout, "\t\t%.12s :", ver)
				}

				fmt.Info.Fprintln(stdout)
				for _, name := range pack {
					fmt.Info.Fprintf(stdout, "\t\t\t%s\n", name)
				}
			}
			fmt.Info.Fprintln(stdout)
		})
	default:
		fmt.Info.Fprintln(os.Stderr, "wrong args", args)
	}
	defer func(out bytes.Buffer) {
		fmt.Info.Fprintln(os.Stdout, out.String())
	}(*stdout)
	l.less(stdout)
	return nil
}

func (l *libsCmd) less(stdin io.Reader) {
	less := exec.Command("less")
	less.Stdin = stdin
	less.Stdout = l.w
	less.Start()
	less.Wait()
}
