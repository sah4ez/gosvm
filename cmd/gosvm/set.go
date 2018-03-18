package main // import "github.com/sah4ez/gosvm/cmd/gosvm"

import (
	"fmt"
	"io"
	"os"

	"github.com/sah4ez/gosvm/pkg/dependency"
	"github.com/sah4ez/gosvm/pkg/dependency/glide"
	"github.com/sah4ez/gosvm/pkg/structure"
)

type setCmd struct {
	w io.Writer
}

var shortHelpSet = "set version for all or cpecifict pacakge\n \tset <LIB_NAME><VERSION> [...<PACK_NAME>]"

func (s *setCmd) Name() string      { return "set" }
func (s *setCmd) ShortHelp() string { return shortHelpSet }
func (s *setCmd) Enable() bool      { return true }
func (s *setCmd) Run(args []string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	switch len(args) {
	case 0, 1, 2:
		return fmt.Errorf("wrong args: %s", args)
	case 3:
		libName := args[1]
		libVer := args[2]
		root, err := structure.LoadStructure(wd + "/svm.toml")
		if err != nil {
			return err
		}
		root.ParseType()
		packs := dependency.NewPackages()
		glideLoader := glide.NewGlideLoader(root, packs)
		//depLoader := dep.NewDepLoader(root, packs)
		err = glideLoader.SetVersionAll(libName, libVer)
		if err != nil {
			return err
		}

		fmt.Fprintf(s.w, "will be set %s for all libs %s\n", libVer, libVer)
	default:
		fmt.Fprintf(s.w, "will be set %s for all libs %s in packs: %s\n", args[2], args[1], args[3:])
	}
	return nil
}
