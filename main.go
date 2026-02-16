package main

import (
	"fmt"
	"os"

	"github.com/devx-base-setup/utils/cmd/mkissue"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	subcommand := os.Args[1]

	switch subcommand {
	case "mkissue":
		mkissue.Run(os.Args[2:])
	case "help", "--help", "-h":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "Unknown subcommand: %s\n\n", subcommand)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: utils <subcommand> [args]")
	fmt.Println("\nAvailable subcommands:")
	fmt.Println("  mkissue <file.issue.md>  Create a GitHub issue from a markdown file")
	fmt.Println("  help                     Show this help message")
}
