package structure

import (
	"io/ioutil"

	toml "github.com/pelletier/go-toml"
)

type Root struct {
	Title       string
	Description string
	Path        string `toml:,omitempty`
	Version     string `toml:,omitempty`
	BasePath    string
	SubProject  []SubProject
}

type SubProject struct {
	Title       string
	Description string
	Path        string `toml:,omitempty`
	Version     string `toml:,omitempty`
	Type        string `toml:,omitempty`
}

func LoadStructure(path string) (*Root, error) {
	var root Root
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err = toml.Unmarshal(data, &root); err != nil {
		return nil, err
	}
	return &root, nil
}
