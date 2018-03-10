package structure

import (
	"os"
	"testing"
)

func TestLoadStructre(t *testing.T) {
	pwd := workDir(t)
	root, err := LoadStructure(pwd + "/testdata/simple_test/TestProject.toml")
	if err != nil {
		t.Fatalf("Couldn't load config. Error %s", err)
	}
	if root.BasePath != "github.com/sah4ez/gosvm/cmd/testdata/test_project" {
		t.Fatal("wrong base path: ", root.BasePath)
	}
	if root.Title != "test_project" {
		t.Fatal("wrong title project")
	}
	if root.Description != "Project with sub-project" {
		t.Fatal("wrong description")
	}
	if root.Version != "1.0.0" {
		t.Fatal("wrong version")
	}
	if len(root.SubProject) != 2 {
		t.Fatal("wrong number sub-projects")
	}
}

func TestParseType(t *testing.T) {
	pwd := workDir(t)
	root, err := LoadStructure(pwd + "/testdata/parse_type/svm.toml")
	if err != nil {
		t.Error("could not parse test svm.toml")
	}
	root.ParseType()
	for _, sub := range root.SubProject {
		if sub.Title == "glide" && sub.Type != Glide {
			t.Errorf("want %s have %s", Glide, sub.Type)
		}
		if sub.Title == "dep" && sub.Type != Dep {
			t.Errorf("want %s have %s", Dep, sub.Type)
		}
		if sub.Title == "go_mod" && sub.Type != GoMod {
			t.Errorf("want %s have %s", GoMod, sub.Type)
		}
	}
}

func workDir(t *testing.T) string {
	pwd, err := os.Getwd()
	if err != nil {
		t.Error("Could not get pwd")
		return ""
	}
	return pwd
}
