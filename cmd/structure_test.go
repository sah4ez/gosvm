package main

import (
	"fmt"
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
	fmt.Printf("%+v", *root)
}
