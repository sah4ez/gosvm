package gomod

import (
	"github.com/sah4ez/gosvm/fs"
	"github.com/sah4ez/gosvm/pkg/dependency"
	"github.com/sah4ez/gosvm/pkg/structure"
)

type goModLoader struct {
	root *structure.Root
	deps *dependency.Packages
}

func (g *goModLoader) Load() (*dependency.Packages, error) {
	subs := g.root.SubProject
	for _, sub := range subs {
		path, ok := fs.PathToGoMod(g.root.BasePath, sub.Title)
		if sub.Type != dependency.GoModType && !ok {
			continue
		}
		gomod, err := LoadGoModFile(path)
		if err != nil {
			return g.deps, err
		}
		for _, req := range gomod.Require {
			pack := dependency.Package{
				Name:       sub.Title,
				LibVersion: req.Version,
			}
			g.deps.Add(req.Path, pack)
		}
	}
	return g.deps, nil
}

func (g *goModLoader) SetVersion(pack string, version string) error {
	panic("not implemented")
}

func (g *goModLoader) SetVersionAll(pack string, version string) error {
	panic("not implemented")
}

func (g *goModLoader) Version(pack string) (string, error) {
	panic("not implemented")
}

func (g *goModLoader) CompareVersion(source string, target string) (bool, error) {
	panic("not implemented")
}

func (g *goModLoader) Update() error {
	panic("not implemented")
}

func NewGoModLoader(root *structure.Root, deps *dependency.Packages) dependency.Loader {
	return &goModLoader{
		root: root,
		deps: deps,
	}
}
