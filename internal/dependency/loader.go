package dependency

// Loader provide methods for each dependecies manager (dep, glide and etc.)
type Loader interface {
	// Load loading slice Packages from config file or return error
	Load() ([]Package, error)
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

// Package provide internal data structure
type Package struct {
	// Name package
	Name string
	// Version package
	Version string
}
