package clients

import (
	"github.com/gdegiorgio/cognitools/internal/service"
	"github.com/gdegiorgio/cognitools/internal/utils"
	"github.com/spf13/cobra"
)

func newListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list <pool-id>",
		Short: "List User Pool Clients",
		Long:  "List User Pool Clients retrieves and displays all clients for a specific user pool.",
		Args:  cobra.ExactArgs(1),
		RunE:  runListCommand,
	}
}

func runListCommand(cmd *cobra.Command, args []string) error {
	return list(cmd, args, service.NewAWSService())
}

func list(cmd *cobra.Command, args []string, svc service.AWS) error {
	poolID := args[0]

	clients, err := svc.ListUserPoolClients(poolID)
	if err != nil {
		cmd.Printf("❌ %s\n", formatAWSError(err))
		return err
	}

	if len(clients) == 0 {
		cmd.Println("❌ No clients found for this user pool.")
		return nil
	}

	outputJSON := getJSONFlag(cmd)
	if outputJSON {
		json, _ := utils.FormatJSON(clients)
		cmd.Println(json)
		return nil
	}

	cmd.Println("👤 User Pool Clients:")
	for _, client := range clients {
		cmd.Printf("  %s - %s\n", *client.ClientName, *client.ClientId)
	}

	return nil
}