package glide

import (
	"os"

	"github.com/sah4ez/gosvm/internal/dependency"
	"github.com/sah4ez/gosvm/internal/structure"
)

type glideLoader struct {
	root *structure.Root
}

func (g *glideLoader) Load() (*dependency.Packages, error) {
	subs := g.root.SubProject
	deps := dependency.NewPackages()
	for _, sub := range subs {
		path := os.Getenv("GOPATH") + "/src/" + g.root.BasePath + "/" + sub.Title + "/glide.yaml"
		glide, err := LoadGlideFile(path)
		if err != nil {
			return deps, err
		}
		for _, gi := range glide.Import {
			pack := dependency.Package{
				Name:       sub.Title,
				LibVersion: gi.Version,
			}
			deps.Add(gi.Package, pack)
		}
	}
	return deps, nil
}

func (g *glideLoader) SetVersion(pack string, version string) error {
	panic("not implemented")
}

func (g *glideLoader) Version(pack string) (string, error) {
	panic("not implemented")
}

func (g *glideLoader) CompareVersion(source string, target string) (bool, error) {
	panic("not implemented")
}

func (g *glideLoader) Update() error {
	panic("not implemented")
}

func NewGlideLoader(root *structure.Root) dependency.Loader {
	return &glideLoader{root: root}
}
