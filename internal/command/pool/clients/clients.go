package clients

import "github.com/spf13/cobra"

func NewCommand() *cobra.Command {
	clientsCmd := &cobra.Command{
		Use:   "clients",
		Short: "Manage User Pool Clients",
		Long:  "Manage User Pool Clients allows you to list and describe clients in a user pool.",
	}

	clientsCmd.AddCommand(newListCommand())
	clientsCmd.AddCommand(newDescribeCommand())
	
	return clientsCmd
}