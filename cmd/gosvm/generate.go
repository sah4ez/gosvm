package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	toml "github.com/pelletier/go-toml"
	"github.com/sah4ez/gosvm/fs"
	st "github.com/sah4ez/gosvm/pkg/structure"
)

type generateCmd struct {
	w io.Writer
}

var shortHelpGenerate = `[OPTION] generate toml template for specification of project
		OPTION:
			custom - generate spec in interactive mode
`

func (l *generateCmd) Name() string      { return "generate" }
func (l *generateCmd) ShortHelp() string { return shortHelpGenerate }
func (l *generateCmd) Enable() bool      { return true }

func (l *generateCmd) Run(args []string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	specFile := wd + "/svm.toml"
	if fs.Exists(specFile) {
		fmt.Fprintln(l.w, "for generate spec remove: ", specFile)
		return ErrSpecProjectExists
	}
	switch len(args) {
	case 1:
		err = l.automationMode(specFile)
	case 2:
		err = l.interactiveMode(specFile, args)
	default:
		fmt.Fprintln(l.w, `
help:
  generate template specifation of project
			`)
	}
	return err
}

func (l *generateCmd) automationMode(specFile string) error {
	docs, err := os.OpenFile(specFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	defer docs.Close()
	tmp := st.SubProject{
		Title:       "Title project, use for seach foler project",
		Description: "Description of project ",
		Path:        "Custom path to project",
		Version:     "0.0.1",
		Type:        "(" + st.Glide + "|" + st.Dep + "|" + st.GoMod + ")",
	}

	doc, err := toml.Marshal(tmp)
	if err != nil {
		return err
	}
	_, err = docs.Write(doc)
	if err != nil {
		return err
	}
	err = l.saveDoc(tmp, docs)
	return err
}

func (l *generateCmd) interactiveMode(specFile string, args []string) error {
	switch args[1] {
	case "custom":
		docs, err := os.OpenFile(specFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
		if err != nil {
			return err
		}
		defer docs.Close()

		tmp := st.SubProject{}

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Title: ")
		if scanner.Scan() {
			tmp.Title = scanner.Text()
		}
		fmt.Print("Description: ")
		if scanner.Scan() {
			tmp.Description = scanner.Text()
		}
		fmt.Print("Path: ")
		if scanner.Scan() {
			tmp.Path = scanner.Text()
		}
		fmt.Print("Version: ")
		if scanner.Scan() {
			tmp.Version = scanner.Text()
		}
		fmt.Printf("Type(%s|%s|%s): ", st.Glide, st.Dep, st.GoMod)
		if scanner.Scan() {
			tmp.Type = scanner.Text()
		}

		fmt.Println("Add custom field <KEY>:<VALUE> (empty for skip and continue): ")

		err = l.saveDoc(tmp, docs)
		if err != nil {
			return err
		}

		for scanner.Scan() {
			str := scanner.Text()
			if str == "" {
				break
			}
			pair := strings.Split(str, ":")
			if len(pair) < 2 {
				fmt.Println("Ivalid key:value. Skiped ':'", str)
				continue
			}
			line := strings.TrimSpace(pair[0]) + " = " + strings.Join(pair[1:], ":") + "\n"
			docs.WriteString(line)
		}
		return scanner.Err()
	}
	return nil
}

func (l *generateCmd) saveDoc(tmp st.SubProject, docs *os.File) error {
	doc, err := toml.Marshal(tmp)
	if err != nil {
		return err
	}
	_, err = docs.Write(doc)
	if err != nil {
		return err
	}
	fmt.Fprintln(l.w, "Generate template:")
	fmt.Fprintln(l.w)
	fmt.Fprintln(l.w, string(doc))
	return nil
}
