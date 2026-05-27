package main

import (
	"flag"
	"log"
	"os"

	"github.com/foomo/sesamy-cli/internal/cli"
	cmdx "github.com/foomo/sesamy-cli/pkg/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	out := flag.String("out", "docs/reference/cli", "output directory")
	flag.Parse()

	if err := os.MkdirAll(*out, 0o755); err != nil {
		log.Fatal(err)
	}

	root := cli.NewCommand(cmdx.NewLogger())
	root.DisableAutoGenTag = true

	if err := doc.GenMarkdownTree(root, *out); err != nil {
		log.Fatal(err)
	}
}
