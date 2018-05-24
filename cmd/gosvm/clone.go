package main // import "github.com/sah4ez/gosvm/cmd/gosvm"

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sah4ez/gosvm/fs"
	"github.com/sah4ez/gosvm/pkg/structure"
	git "gopkg.in/src-d/go-git.v4"
)

type cloneCmd struct {
	w io.Writer
}

var shortHelpPull = "clone not exists repositories"

func (l *cloneCmd) Name() string      { return "clone" }
func (l *cloneCmd) ShortHelp() string { return shortHelpPull }
func (l *cloneCmd) Enable() bool      { return true }

func (l *cloneCmd) Run(args []string) error {
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
		fmt.Fprintln(l.w, "Check and clone not exists:")
		fmt.Fprintln(l.w, "")
		fmt.Fprintln(l.w, "Base path: ", root.BasePath)
		for _, sub := range root.SubProject {
			fullPath := fs.PathToProject(root.BasePath, sub.Title)
			if !fs.ExistsGoProject(root.BasePath, sub.Title) {
				httpsURL := "https://" + root.BasePath + "/" + sub.Title + ".git"
				_, err := git.PlainClone(
					fullPath,
					false,
					&git.CloneOptions{
						URL:      httpsURL,
						Progress: l.w,
					},
				)
				if err != nil {
					err = os.RemoveAll(fullPath)
					if err != nil {
						return err
					}
					sshURL := strings.Replace(httpsURL, "https://", "git@", 1)
					sshURL = strings.Replace(sshURL, "/", ":", 1)
					_, err = git.PlainClone(
						fullPath,
						false,
						&git.CloneOptions{
							URL:      sshURL,
							Progress: l.w,
						},
					)
				}
				if err != nil {
					return err
				}
				fmt.Fprintf(l.w, "%s...\t%s\n", fullPath, "Cloned")
				continue
			}
			fmt.Fprintf(l.w, "%s...\t%s\n", fullPath, "Exists")
		}
	default:
		fmt.Fprintln(l.w, "from args", args)
	}
	return nil
}
