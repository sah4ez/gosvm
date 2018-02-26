package dep

import "testing"

func TestLoadDepFile(t *testing.T) {
	path := "testdata/Gopkg.toml"

	dep, err := LoadDepFile(path)

	if err != nil {
		t.Error("could not load file", err)
	}
	if len(dep.Constraints) != 2 {
		t.Error("wrong count constraints")
	}

	if dep.Constraints[0].Name != "github.com/pelletier/go-toml" {
		t.Error("wrong name")
	}
	if dep.Constraints[0].Version != "1.1.0" {
		t.Error("wrong version")
	}

	if dep.Constraints[1].Name != "gopkg.in/yaml.v2" {
		t.Error("wrong name")
	}
	if dep.Constraints[1].Version != "2.1.1" {
		t.Error("wrong version")
	}
}
