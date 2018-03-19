package main // import "github.com/sah4ez/gosvm/cmd/gosvm"

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"
)

type command interface {
	Name() string
	Enable() bool
	ShortHelp() string
	Run(args []string) error
}

func main() {
	commands := [...]command{
		&listCmd{w: os.Stdout},
		&libsCmd{w: os.Stdout},
		&setCmd{w: os.Stdout},
		&versionCmd{w: os.Stdout},
		&cloneCmd{w: os.Stdout},
	}

	examples := [...][2]string{
		{
			"gosvm list",
			"show list all services for current project",
		},
		{
			"gosvm libs",
			"run checking all libs in list repositores and find different version",
		},
		{
			"gosvm set bbb.com/path/to/lib 1.2.3 pack.name",
			"set version of lib in specifict packages",
		},
		{
			"gosvm clone",
			"clone not exists repositories to basePath in svm.toml",
		},
	}

	usage := func(w io.Writer) {
		fmt.Fprintln(w, "gosvm is a tool for managing dependencies of services in")
		fmt.Fprintln(w, "microservices architecture approach.")
		fmt.Fprintln(w)
		fmt.Fprintln(w, "Usage: \"gosvm [command]\"")
		fmt.Fprintln(w)
		fmt.Fprintln(w, "Commands:")
		fmt.Fprintln(w)
		tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
		for _, cmd := range commands {
			if cmd.Enable() {
				fmt.Fprintf(tw, "\t%s\t%s\n", cmd.Name(), cmd.ShortHelp())
			}
		}
		tw.Flush()
		fmt.Fprintln(w)
		fmt.Fprintln(w, "Examples:")
		for _, example := range examples {
			fmt.Fprintf(tw, "\t%s\t%s\n", example[0], example[1])
		}
		tw.Flush()
		fmt.Fprintln(w)
		fmt.Fprintln(w, "Use \"gosvm help [command]\" for more information about a command.")
	}

	cmdName, _, exit := parseArgs(os.Args)
	if exit {
		usage(os.Stderr)
		os.Exit(1)
	}

	for _, cmd := range commands {
		if cmd.Name() == cmdName {
			cmdArgs := os.Args[1:]
			err := cmd.Run(cmdArgs)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Could not execute %s. Error: %s\n", cmdName, err)
				os.Exit(1)
			}
		}
	}
}

func parseArgs(args []string) (cmdName string, printHelpCmd bool, exit bool) {
	isHelpArg := func() bool {
		return strings.Contains(strings.ToLower(args[1]), "help") || strings.ToLower(args[1]) == "-h"
	}

	switch len(args) {
	case 0, 1:
		exit = true
	case 2:
		if isHelpArg() {
			exit = true
		} else {
			cmdName = args[1]
		}
	default:
		if isHelpArg() {
			cmdName = args[2]
			printHelpCmd = true
		} else {
			cmdName = args[1]
		}
	}
	return cmdName, printHelpCmd, exit
}
