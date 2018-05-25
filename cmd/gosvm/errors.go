package main

import "fmt"

var (
	ErrProjectNotFound   = fmt.Errorf("project not found")
	ErrSpecProjectExists = fmt.Errorf("svm.toml spec exist for current project")
)
