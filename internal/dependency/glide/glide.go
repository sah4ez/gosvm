package glide

import (
	"github.com/sah4ez/gosvm/fs"
	"github.com/sah4ez/gosvm/internal/dependency"
	"github.com/sah4ez/gosvm/internal/structure"
)

type glideLoader struct {
	root *structure.Root
	deps *dependency.Packages
}

func (g *glideLoader) Load() (*dependency.Packages, error) {
	subs := g.root.SubProject
	for _, sub := range subs {
		path, ok := fs.PathToGlide(g.root.BasePath, sub.Title)
		if sub.Type != dependency.GlideType && !ok {
			//			fmt.Fprintln(os.Stderr, "missing type Package and could not load file", path)
			continue
		}
		glide, err := LoadGlideFile(path)
		if err != nil {
			return g.deps, err
		}
		for _, gi := range glide.Import {
			pack := dependency.Package{
				Name:       sub.Title,
				LibVersion: gi.Version,
			}
			g.deps.Add(gi.Package, pack)
		}
	}
	return g.deps, nil
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

func NewGlideLoader(root *structure.Root, deps *dependency.Packages) dependency.Loader {
	return &glideLoader{
		root: root,
		deps: deps,
	}
}
