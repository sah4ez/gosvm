package glide

import (
	"os"
	"testing"

	"github.com/sah4ez/gosvm/pkg/dependency"
	"github.com/sah4ez/gosvm/pkg/structure"
)

func TestLoadGlide(t *testing.T) {
	//exp := make(map[string]map[string][]string)

	//exp := {
	//	"import/first/package":{
	//		"1.0.0":{"sub_project1", "sub_project2"},
	//	},
	//	"import/second/package":{
	//		"2.0.0":{"sub_project1", "sub_project2"},
	//	}, //} exp := {""} pwd, err := os.Getwd()
	pwd, err := os.Getwd()
	if err != nil {
		t.Error("Could not get pwd")
	}
	root, err := structure.LoadStructure(pwd + "/testdata/test_project/svm.toml")
	if err != nil {
		t.Fatalf("Couldn't load config. Error %s", err)
	}
	packs := dependency.NewPackages()
	gl := NewGlideLoader(root, packs)

	packs, err = gl.Load()
	if err != nil {
		t.Error("could not load libs map. ", err)
	}
	t.Logf("%+v", packs)
}
