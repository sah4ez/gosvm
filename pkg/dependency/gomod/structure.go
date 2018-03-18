package gomod

import (
	"strings"

	"github.com/sah4ez/gosvm/fs"
)

const (
	space        = " "
	closeBracket = ")"
	doubleQoute  = "\""
)

type GoMod struct {
	Require []Require
}

type Require struct {
	Path    string
	Version string
}

func LoadGoModFile(path string) (*GoMod, error) {
	gomod := &GoMod{}

	reader, err, fClose := fs.ReadFile(path)
	if err != nil {
		return nil, err
	}
	defer fClose()

	required := false
	for {
		line, err := reader.ReadString('\n')

		if required && line != "" {
			line = strings.TrimSpace(line)
			words := strings.Split(line, space)
			if len(words) >= 2 {
				req := Require{
					Path:    strings.Trim(words[0], doubleQoute),
					Version: words[1],
				}
				gomod.Require = append(gomod.Require, req)
			}
		}

		if strings.Contains(line, "require") {
			required = true
		}

		if required && strings.TrimSpace(line) == closeBracket {
			required = false
			break
		}

		if err != nil {
			break
		}
	}
	return gomod, nil
}
