package dep // import "github.com/sah4ez/gosvm/pkg/dependency/dep"

import (
	"github.com/sah4ez/gosvm/fs"
	"github.com/sah4ez/gosvm/pkg/dependency"
	"github.com/sah4ez/gosvm/pkg/structure"
)

type depLoader struct {
	root *structure.Root
	deps *dependency.Packages
}

func (d *depLoader) Load() (*dependency.Packages, error) {
	subs := d.root.SubProject
	for _, sub := range subs {
		path, ok := fs.PathToDep(d.root.BasePath, sub.Title)
		if sub.Type != dependency.TomlType && !ok {
			//fmt.Fprintln(os.Stderr, "missing type Package and could not load file", path)
			continue
		}
		dep, err := LoadDepFile(path)
		if err != nil {
			return d.deps, err
		}
		for _, dc := range dep.Constraints {
			pack := dependency.Package{
				Name:       sub.Title,
				LibVersion: dc.Version,
			}
			d.deps.Add(dc.Name, pack)
		}
	}
	libs := d.root.Library
	for _, sub := range libs {
		path, ok := fs.PathToDep(d.root.BasePath, sub.Title)
		if sub.Type != dependency.TomlType && !ok {
			//fmt.Fprintln(os.Stderr, "missing type Package and could not load file", path)
			continue
		}
		dep, err := LoadDepFile(path)
		if err != nil {
			return d.deps, err
		}
		for _, dc := range dep.Constraints {
			pack := dependency.Package{
				Name:       sub.Title,
				LibVersion: dc.Version,
			}
			d.deps.Add(dc.Name, pack)
		}
	}
	return d.deps, nil
}

func (d *depLoader) SetVersion(pack string, version string) error {
	panic("not implemented")
}

func (d *depLoader) SetVersionAll(pack string, version string) error {
	panic("not implemented")
}

func (d *depLoader) Version(pack string) (string, error) {
	panic("not implemented")
}

func (d *depLoader) CompareVersion(source string, target string) (bool, error) {
	panic("not implemented")
}

func (d *depLoader) Update() error {
	panic("not implemented")
}

func NewDepLoader(root *structure.Root, deps *dependency.Packages) dependency.Loader {
	return &depLoader{
		root: root,
		deps: deps,
	}
}
