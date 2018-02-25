package glide

import (
	"os"
	"testing"

	"github.com/sah4ez/gosvm/internal/structure"
)

func TestLoadGlide(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Error("Could not get pwd")
	}
	root, err := structure.LoadStructure(pwd + "/testdata/test_project/svm.toml")
	if err != nil {
		t.Fatalf("Couldn't load config. Error %s", err)
	}

	gl := NewGlideLoader(root)

	packs, err := gl.Load()
	if err != nil {
		t.Error("could not load libs map. ", err)
	}
	t.Logf("%+v", packs)
}
