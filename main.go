package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const usage = `Team tools for the Shoald GitHub org.

Usage:
  gh shoald <command> [flags]

Commands:
  new [--template=base|go] <name>   Create a repo from a template (default: base)
  help                              Show this help

Examples:
  gh shoald new repo-name
  gh shoald new --template=go repo-name
`

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 0 {
		fmt.Print(usage)
		return nil
	}

	switch args[0] {
	case "new":
		return runNew(args[1:])
	case "help", "h", "-h", "--help":
		fmt.Print(usage)
		return nil
	default:
		return fmt.Errorf("unknown command %q\n\n%s", args[0], usage)
	}
}

func runNew(args []string) error {
	fs := flag.NewFlagSet("new", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	template := fs.String("template", "base", "template to use: base or go")

	if err := fs.Parse(args); err != nil {
		return err
	}
	if *template != "base" && *template != "go" {
		return fmt.Errorf("unknown template %q, use base or go", *template)
	}
	names := fs.Args()
	if len(names) != 1 {
		return fmt.Errorf("new needs exactly one repo name, after any flags\n\n%s", usage)
	}

	fmt.Printf("new: would create %s from template-%s\n", names[0], *template)
	return nil
}

/*
func main() {
	fmt.Println("hi world, this is the gh-shoald extension!")
	client, err := api.DefaultRESTClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	response := struct {Login string}{}
	err = client.Get("user", &response)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("running as %s\n", response.Login)
}
*/

// For more examples of using go-gh, see:
// https://github.com/cli/go-gh/blob/trunk/example_gh_test.go
