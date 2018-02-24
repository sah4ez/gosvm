package structure

import (
	"os"
	"testing"
)

func TestLoadStructre(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Error("Could not get pwd")
	}
	root, err := LoadStructure(pwd + "/testdata/TestProject.toml")
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
