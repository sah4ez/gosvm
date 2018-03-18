package main // import "github.com/sah4ez/gosvm/cmd/gosvm"
import (
	"fmt"
	"io"
)

var (
	//Version current tools
	Version string
	// Hash revision number from git
	Hash string
	// BuildDate when building this utilitites
	BuildDate string
)

type versionCmd struct {
	w io.Writer
}

var shortHelpVersion = "version return current version gosvm"

func (v *versionCmd) Name() string      { return "version" }
func (v *versionCmd) ShortHelp() string { return shortHelpVersion }
func (v *versionCmd) Enable() bool      { return true }
func (v *versionCmd) Run(args []string) error {
	fmt.Fprintf(v.w, "Version %s \nRevision %s \nDate %s\n", Version, Hash, BuildDate)
	return nil
}
