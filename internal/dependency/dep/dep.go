package dep

import "github.com/sah4ez/gosvm/internal/dependency"

type depLoader struct{}

func (d *depLoader) Load() (*dependency.Packages, error) {
	panic("not implemented")
}

func (d *depLoader) SetVersion(pack string, version string) error {
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

func New() dependency.Loader {
	return &depLoader{}
}
