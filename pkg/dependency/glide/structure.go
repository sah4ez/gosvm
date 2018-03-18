package glide

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Glide struct {
	Package string
	Import  []Import
}

type Import struct {
	Package string
	Version string
}

func LoadGlideFile(path string) (*Glide, error) {
	glide := &Glide{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, glide)
	if err != nil {
		return nil, err
	}
	return glide, nil
}
