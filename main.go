package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// AppName is the name of the built binary.
const AppName = "templatefile"

// LicenseInfo is the license information for the app.
const LicenseInfo = "License MPL-2.0 <https://www.mozilla.org/en-US/MPL/2.0/>"

// Version is the application version. This is filled in by the compiler.
var Version = "0.0.0"

// GitCommit is the commit that was build. This is filled in by the compiler.
var GitCommit = "HEAD"

// UsageText is the help text for the root command.
var UsageText = fmt.Sprintf("Usage: %s [-v] path vars", AppName)

func main() {
	argv := os.Args[1:]

	help(argv)
	version(argv)
	showUsage(argv)

	if argv[1] == "-" {
		argv[1] = os.Stdin.Name()
	}

	vars, err := ioutil.ReadFile(argv[1])
	pp(err)

	out, err := templatefile(".", argv[0], string(vars))
	pp(err)

	_, err = fmt.Fprintf(os.Stdout, "%s", out)
	pp(err)
	os.Exit(0)
}

func sliceContains(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}

	return false
}

func help(argv []string) {
	if !sliceContains(argv, "--help") {
		return
	}

	_, err := fmt.Fprintf(os.Stdout, "%s\n\n%s\n", "Render a Terraform template file", UsageText)
	pp(err)
	os.Exit(0)
}

func version(argv []string) {
	if !sliceContains(argv, "--version") {
		return
	}

	v := fmt.Sprintf("%s %s-%s", AppName, Version, GitCommit)
	_, err := fmt.Fprintf(os.Stdout, "%s\n%s\n", v, LicenseInfo)
	pp(err)
	os.Exit(0)
}

func showUsage(argv []string) {
	if len(argv) == 2 {
		return
	}

	_, err := fmt.Fprintf(os.Stderr, "%s\n", UsageText)
	pp(err)
	os.Exit(1)
}

func pp(err error) {
	if err == nil {
		return
	}

	_, err = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	if err != nil {
		panic(err)
	}

	os.Exit(1)
}
