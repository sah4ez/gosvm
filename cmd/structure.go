package main

import (
	"github.com/BurntSushi/toml"
)

type Project struct {
	Title       string
	Description string
	Path        string `toml:,omitempty`
	Version     string `toml:,omitempty`
}

type Root struct {
	Project
	BasePath   string
	SubProject []SubProject
}

type SubProject struct {
	Project
	Type string `toml:,omitempty`
}

func LoadStructure(path string) (*Root, error) {
	var root Root
	if _, err := toml.DecodeFile(path, &root); err != nil {
		return nil, err
	}
	return &root, nil
}
