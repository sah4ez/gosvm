package structure

import (
	"io/ioutil"
	"os"
	"strings"

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

const (
	Glide = "glide.yaml"
	Dep   = "Gopkg.toml"
	GoMod = "go.mod"
)

func (r *Root) ParseType() {
	files := []string{Glide, Dep, GoMod}
	for i := range r.SubProject {
		for _, file := range files {
			path := strings.Join([]string{os.Getenv("GOPATH"), "src", r.BasePath, r.SubProject[i].Title, file}, "/")
			if _, err := os.Stat(path); err == nil {
				r.SubProject[i].Type = file
			}
		}
	}
}
