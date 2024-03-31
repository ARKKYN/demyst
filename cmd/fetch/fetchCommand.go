
package fetch

import (
	"demyst/clients"
	"demyst/services"
	"github.com/spf13/cobra"
)

func NewFetchCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fetch",
		Short: "fetches Todos",
		Long:  `Fetches todos from todo api`,
		RunE:  runFetchCommand,
	}
	return cmd
}

func runFetchCommand(cmd *cobra.Command, args []string) error {
	out := cmd.OutOrStdout()
	todoClient := clients.NewTodoClient()
	fetchService := services.NewFetchService(out, todoClient)
	return fetchService.FetchTodos()
}
