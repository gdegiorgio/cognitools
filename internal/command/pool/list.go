package pool

import "github.com/spf13/cobra"

func newListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List User Pools",
		Long:  "List User Pools retrieves and displays all user pools in your application.",
		Run:   runListCommand,
	}
}

func runListCommand(cmd *cobra.Command, args []string) {
}
