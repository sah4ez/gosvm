package main // import "github.com/sah4ez/gosvm/cmd/gosvm"

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/sah4ez/gosvm/fs"
	"github.com/sah4ez/gosvm/pkg/structure"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	gitssh "gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
)

type pullCmd struct {
	w io.Writer
}

var shortHelpFetch = "pull not exists repositories"

func (l *pullCmd) Name() string      { return "pull" }
func (l *pullCmd) ShortHelp() string { return shortHelpFetch }
func (l *pullCmd) Enable() bool      { return true }

func (l *pullCmd) Run(args []string) error {
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
		fmt.Fprintln(l.w, "Pull repositories:")
		fmt.Fprintln(l.w, "")
		fmt.Fprintln(l.w, "Base path: ", root.BasePath)
		for _, sub := range root.SubProject {
			fullPath := fs.PathToProject(root.BasePath, sub.Title)
			if !fs.ExistsGoProject(root.BasePath, sub.Title) {
				fmt.Fprintf(l.w, "%s...\t%s\n", fullPath, "Not exists. Please call:\n\tgosvm clone")
				continue
			}
			r, err := git.PlainOpen(fullPath)
			if err != nil {
				return err
			}

			w, err := r.Worktree()
			if err != nil {
				return err
			}
			pk := defaultKeyPath()
			auth, err := gitssh.NewPublicKeysFromFile("git", pk, "")
			if err != nil {
				return err
			}

			err = w.Pull(&git.PullOptions{
				ReferenceName: plumbing.Master,
				Progress:      os.Stdout,
				Auth:          auth,
				Force:         true,
			})
			if err != nil && err != git.NoErrAlreadyUpToDate {
				return err
			}
			ref, err := r.Head()
			if err != nil {
				return err
			}

			commit, err := r.CommitObject(ref.Hash())
			if err != nil {
				return err
			}

			fmt.Fprintf(l.w, "%s...\t%s\n", fullPath, commit)
		}
	default:
		fmt.Fprintln(l.w, "from args", args)
	}
	return nil
}

func defaultKeyPath() string {
	home := os.Getenv("HOME")
	if len(home) > 0 {
		return path.Join(home, "/.ssh/id_rsa")
	}
	return ""
}
