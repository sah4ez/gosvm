package main // import "github.com/sah4ez/gosvm/cmd/gosvm"
import (
	"io"
	"io/ioutil"
	"os"

	stdfmt "fmt"

	toml "github.com/pelletier/go-toml"
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
		err = l.toTOML(wd)
	default:
		stdfmt.Fprintln(l.w, `
help:
  Command for generation documentation for all subproject 
  in the one document svm_doc.toml.
  
  If svm.toml not found in project folder, infomation will be 
  add from root svm.toml.
	
	Usage:

	gosvm doc 
	`)
	}
	return err
}

func (l *docCmd) toTOML(wd string) error {
	docSpec := wd + "/svm_doc.toml"
	if fs.Exists(docSpec) {
		err := os.Remove(docSpec)
		if err != nil {
			return err
		}
	}
	docs, err := os.OpenFile(docSpec, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer docs.Close()
	root, err := structure.LoadStructure(wd + "/svm.toml")
	if err != nil {
		return err
	}

	for _, sub := range root.SubProject {
		fullPath := fs.PathToProject(root.BasePath, sub.Title)
		if !fs.ExistsGoProject(root.BasePath, sub.Title) {
			return ErrProjectNotFound
		}
		specFile := fullPath + "/svm.toml"
		if !fs.Exists(specFile) {
			fmt.Warn.Fprintln(l.w, "[WARN] file not exist: "+specFile)
			b, err := toml.Marshal(sub)
			if err != nil {
				return nil
			}
			err = writeSubProject(docs, b)
			if err != nil {
				return err
			}
			continue
		}
		data, err := ioutil.ReadFile(specFile)
		if err != nil {
			return err
		}
		err = writeSubProject(docs, data)
		if err != nil {
			return err
		}
		stdfmt.Fprintln(l.w, specFile)
		stdfmt.Fprintln(l.w, string(data))
		stdfmt.Fprintln(l.w, "")
	}
	return nil
}

func writeSubProject(file *os.File, data []byte) error {
	_, err := file.WriteString("[SubPackage]\n")
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	_, err = file.WriteString("\n")
	return err
}
