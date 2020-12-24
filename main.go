package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const UsageText = "Usage: templatefile [-v] path vars"

func main() {
	argv := os.Args[1:]
	baseDir, _ := os.Getwd()

	help(argv)
	version(argv)
	showUsage(argv)

	if argv[1] == "-" {
		argv[1] = os.Stdin.Name()
	}

	vars, err := ioutil.ReadFile(argv[1])
	if err != nil {
		panic(err)
	}

	out, err := templatefile(baseDir, argv[0], string(vars))
	if err != nil {
		panic(err)
	}

	_, err = fmt.Fprintf(os.Stdout, "%s\n", out)
	if err != nil {
		panic(err)
	}

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
	if err != nil {
		panic(err)
	}

	os.Exit(0)
}

func version(argv []string) {
	if !sliceContains(argv, "--version") {
		return
	}

	version := "templatefile version 1.0.0"
	license := "License MPL-2.0 <https://www.mozilla.org/en-US/MPL/2.0/>"
	_, err := fmt.Fprintf(os.Stdout, "%s\n%s\n", version, license)
	if err != nil {
		panic(err)
	}

	os.Exit(0)
}

func showUsage(argv []string) {
	if len(argv) == 2 {
		return
	}

	_, err := fmt.Fprintf(os.Stderr, "%s\n", UsageText)
	if err != nil {
		panic(err)
	}

	os.Exit(1)
}
