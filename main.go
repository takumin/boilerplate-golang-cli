package main

import (
	"os"

	"github.com/takumin/boilerplate-golang-cli/internal/command"
)

func main() {
	os.Exit(command.Main(
		os.Stdout,
		os.Stderr,
		os.Stdin,
		os.Args,
	))
}
