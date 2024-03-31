package cmd

import (
	"demyst/cmd/fetch"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{Use: "demyst"}
	rootCmd.AddCommand(fetch.NewFetchCommand())
	return rootCmd
}