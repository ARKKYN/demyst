package main

import (
	"demyst/cmd"
	"os"
)

func main() {
	rootCmd := cmd.NewRootCommand()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
