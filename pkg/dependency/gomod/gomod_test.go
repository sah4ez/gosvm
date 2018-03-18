package gomod

import (
	"os"
	"testing"

	"github.com/sah4ez/gosvm/pkg/dependency"
	"github.com/sah4ez/gosvm/pkg/structure"
)

func TestLoadGoMod(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Error("Could not get pwd")
	}
	root, err := structure.LoadStructure(pwd + "/testdata/test_project/svm.toml")
	if err != nil {
		t.Fatalf("Couldn't load config. Error %s", err)
	}
	packs := dependency.NewPackages()
	loader := NewGoModLoader(root, packs)

	packs, err = loader.Load()
	if err != nil {
		t.Error("could not load libs map. ", err)
	}
	t.Logf("%+v", packs)
}
