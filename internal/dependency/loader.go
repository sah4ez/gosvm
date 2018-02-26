package dependency

import "sort"

var (
	GlideType = "glide"
	TomlType  = "toml"
)

// Loader provide methods for each dependecies manager (dep, glide and etc.)
type Loader interface {
	// Load loading slice Packages from config file or return error
	Load() (*Packages, error)
	// SetVersion change for pack Package on new or return error
	SetVersion(pack, version string) error
	// Version return version for pack Package or return error
	Version(pack string) (string, error)
	// CompareVersion comparation source and targer Package  version and
	// if equals return true, else false or return false and error
	CompareVersion(source, target string) (bool, error)
	// Update execute command update from dependencies manager (dep ensure --update or glide up, etc.)
	Update() error
}

type Packages struct {
	packages map[string]map[string][]string
}

// Package provide internal data structure
type Package struct {
	// Name package
	Name string
	// Version package
	LibVersion string
}

func NewPackages() *Packages {
	return &Packages{make(map[string]map[string][]string)}
}

func (p *Packages) Add(lib string, pack Package) {
	if _, ok := p.packages[lib]; !ok {
		p.packages[lib] = make(map[string][]string)
	}
	if _, ok := p.packages[lib][pack.LibVersion]; !ok {
		p.packages[lib][pack.LibVersion] = []string{}
	}
	p.packages[lib][pack.LibVersion] = append(p.packages[lib][pack.LibVersion], pack.Name)
}

func (p *Packages) Range(apply func(string, map[string][]string)) {
	keys := make([]string, 0, len(p.packages))
	for key := range p.packages {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		apply(key, p.packages[key])
	}
}
