package main // import "github.com/sah4ez/gosvm/cmd/gosvm"

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/fatih/color"
)

type command interface {
	Name() string
	Enable() bool
	ShortHelp() string
	Run(args []string) error
}

var (
	noColorized = flag.Bool("no-color", false, "Disable color output")
)

func main() {
	flag.Parse()
	if *noColorized {
		color.NoColor = true
	}

	commands := [...]command{
		&listCmd{w: os.Stdout},
		&libsCmd{w: os.Stdout},
		&setCmd{w: os.Stdout},
		&versionCmd{w: os.Stdout},
		&cloneCmd{w: os.Stdout},
		&docCmd{w: os.Stdout},
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
		{
			"gosvm doc",
			"get from each repos svm.toml and add it to root svm.toml project",
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

	cmdName, _, exit, args := parseArgs(os.Args)
	if exit {
		usage(os.Stderr)
		os.Exit(1)
	}

	for _, cmd := range commands {
		if cmd.Name() == cmdName {
			cmdArgs := args[1:]
			err := cmd.Run(cmdArgs)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Could not execute %s. Error: %s\n", cmdName, err)
				os.Exit(1)
			}
		}
	}
}

func parseArgs(args []string) (cmdName string, printHelpCmd bool, exit bool, rargs []string) {
	isHelpArg := func() bool {
		return strings.Contains(strings.ToLower(args[1]), "help") || strings.ToLower(args[1]) == "-h"
	}
	for i, arg := range args {
		if arg[0] == '-' {
			args = append(args[:i], args[i+1:]...)
		}
	}
	rargs = make([]string, len(args), len(args))
	copy(rargs, args)

	switch len(rargs) {
	case 0, 1:
		exit = true
	case 2:
		if isHelpArg() {
			exit = true
		} else {
			cmdName = rargs[1]
		}
	default:
		if isHelpArg() {
			cmdName = rargs[2]
			printHelpCmd = true
		} else {
			cmdName = rargs[1]
		}
	}
	return cmdName, printHelpCmd, exit, rargs
}
