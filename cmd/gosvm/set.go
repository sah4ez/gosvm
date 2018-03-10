package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/sah4ez/gosvm/internal/dependency"
	"github.com/sah4ez/gosvm/internal/dependency/glide"
	structure "github.com/sah4ez/gosvm/internal/structure"
)

type setCmd struct{}

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
		return errors.New(fmt.Sprintf("wrong args: %s", args))
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

		fmt.Fprintf(os.Stdout, "will be set %s for all libs %s\n", libVer, libVer)
	default:
		fmt.Fprintf(os.Stdout, "will be set %s for all libs %s in packs: %s\n", args[2], args[1], args[3:])
	}
	return nil
}
