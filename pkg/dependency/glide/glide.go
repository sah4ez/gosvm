package glide

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/sah4ez/gosvm/fs"
	"github.com/sah4ez/gosvm/pkg/dependency"
	"github.com/sah4ez/gosvm/pkg/structure"
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

func (g *glideLoader) SetVersionAll(pack string, version string) error {
	for _, sub := range g.root.SubProject {
		if sub.Type != structure.Glide {
			continue
		}

		path, ok := fs.PathToGlide(g.root.BasePath, sub.Title)
		if !ok {
			return fmt.Errorf("could not find path to glide.yaml: %s", sub.Title)
		}

		reader := fs.ReadFile(path)

		result := []string{}
		var line string
		var found bool
		var pos int
		for {
			line, err = reader.ReadString('\n')
			if strings.Contains(line, pack) {
				found = true
			}
			if found && strings.Contains(line, "version") {
				found = false
				v := strings.Split(line, ":")
				line = strings.Join([]string{v[0], version}, ": ")
				line += "\n"
			}
			pos += len(line)
			result = append(result, line)

			if err != nil {
				break
			}
		}
		result = append(result, "")
		err = ioutil.WriteFile(path, []byte(strings.Join(result, "")), 0644)
		if err != nil {
			return err
		}

		fmt.Println("loaded: \n", strings.Join(result, ""))
		file.Close()
	}
	return nil
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
