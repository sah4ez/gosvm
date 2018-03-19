package fs

import (
	"os"
	"testing"
)

func TestPathToProject(t *testing.T) {
	basePath := "github.com/sah4ez/gosvm/fs/testdata"
	projectExists := "exists"
	exp := PathToProject(basePath, projectExists)
	if exp != os.Getenv("GOPATH")+"/src/"+basePath+"/"+projectExists {
		t.Error("expected path not found", exp)
	}
}

func TestExistGoProject(t *testing.T) {
	basePath := "github.com/sah4ez/gosvm/fs/testdata"
	projectExists := "exists"
	projectNotExists := "notExists"
	exists := ExistsGoProject(basePath, projectExists)
	if !exists {
		t.Error("expected path not found", PathToProject(basePath, projectExists))
	}
	notExists := ExistsGoProject(basePath, projectNotExists)
	if notExists {
		t.Error("not expected path exists", PathToProject(basePath, projectNotExists))
	}
}
