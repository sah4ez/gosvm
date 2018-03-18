package gomod

import (
	"fmt"
	"testing"
)

func TestLoadGoModFile(t *testing.T) {
	path := "testdata/go.mod"

	gomod, err := LoadGoModFile(path)
	if err != nil {
		t.Error("could not load file", err)
	}
	if len(gomod.Require) != 2 {
		t.Error("loaded wrong number requireds, ", len(gomod.Require))
	}

	if gomod.Require[0].Version != "v1.1.0" {
		t.Error("loaded wrong number version, ", gomod.Require[0].Version)
	}
	if gomod.Require[1].Version != "v1.1.1-gopkgin-v2.1.1" {
		t.Error("loaded wrong number version, ", gomod.Require[1].Version)
	}
	fmt.Printf("%+v\n", gomod)
}
