package dep

import (
	"io/ioutil"

	toml "github.com/pelletier/go-toml"
)

type Dep struct {
	Constraints []Constraint `toml:"constraint,omitempty"`
	Overrides   []Constraint `toml:"override,omitempty"`
	Ignored     []string     `toml:"ignored,omitempty"`
	Required    []string     `toml:"required,omitempty"`
}

type Constraint struct {
	Name     string `toml:"name"`
	Branch   string `toml:"branch,omitempty"`
	Revision string `toml:"revision,omitempty"`
	Version  string `toml:"version,omitempty"`
	Source   string `toml:"source,omitempty"`
}

func LoadDepFile(path string) (*Dep, error) {
	dep := &Dep{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = toml.Unmarshal(data, dep)
	if err != nil {
		return nil, err
	}
	return dep, nil
}
