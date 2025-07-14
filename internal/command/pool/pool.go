package pool

import "github.com/spf13/cobra"

var outputJSON bool

func NewCommand() *cobra.Command {
	poolCmd := &cobra.Command{
		Use:   "pool",
		Short: "Manage User Pools",
		Long:  "Manage User Pools allows you to create, update, delete, and list user pools in your application.",
	}

	poolCmd.PersistentFlags().BoolVar(&outputJSON, "json", false, "Output in JSON format")

	poolCmd.AddCommand(newListCommand())
	return poolCmd
}
