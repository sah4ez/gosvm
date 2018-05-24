package main // import "github.com/sah4ez/gosvm/cmd/gosvm"
import (
	"io"
	"io/ioutil"
	"os"

	"github.com/sah4ez/gosvm/fs"
	fmt "github.com/sah4ez/gosvm/pkg/formatting"
	"github.com/sah4ez/gosvm/pkg/structure"
)

type docCmd struct {
	w io.Writer
}

var shortHelpDoc = "collect all documents about repositories"

func (l *docCmd) Name() string      { return "doc" }
func (l *docCmd) ShortHelp() string { return shortHelpDoc }
func (l *docCmd) Enable() bool      { return true }

func (l *docCmd) Run(args []string) error {
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

		for _, sub := range root.SubProject {
			fullPath := fs.PathToProject(root.BasePath, sub.Title)
			specFile := fullPath + "/svm.toml"
			if !fs.Exists(specFile) {
				fmt.Warn.Fprintln(l.w, "[WARN] file not exist: "+specFile)
				continue
			}
			data, err := ioutil.ReadFile(specFile)
			if err != nil {
				return err
			}
			if !fs.ExistsGoProject(root.BasePath, sub.Title) {
				return ErrProjectNotFound
			}
			fmt.Info.Fprintln(l.w, "")
			fmt.Info.Fprintln(l.w, string(data))
			fmt.Info.Fprintln(l.w, "")
		}
	}
	return nil
}
