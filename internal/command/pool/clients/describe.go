package clients

import (
	"strings"

	"github.com/gdegiorgio/cognitools/internal/service"
	"github.com/gdegiorgio/cognitools/internal/utils"
	"github.com/spf13/cobra"
)

func newDescribeCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "describe <pool-id> <client-id>",
		Short: "Describe User Pool Client",
		Long:  "Describe User Pool Client retrieves and displays detailed information about a specific client.",
		Args:  cobra.ExactArgs(2),
		RunE:  runDescribeCommand,
	}
}

func runDescribeCommand(cmd *cobra.Command, args []string) error {
	return describe(cmd, args, service.NewAWSService())
}

func describe(cmd *cobra.Command, args []string, svc service.AWS) error {
	poolId := args[0]
	clientId := args[1]

	client, err := svc.DescribeUserPoolClient(poolId, clientId)
	if err != nil {
		cmd.Printf("❌ %s\n", formatAWSError(err))
		return err
	}

	// Check if parent command has JSON flag
	outputJSON, _ := cmd.Flags().GetBool("json")
	if !cmd.Flags().Changed("json") {
		// Check parent command for JSON flag
		if parent := cmd.Parent(); parent != nil {
			if parent.Flags().Lookup("json") != nil {
				outputJSON, _ = parent.Flags().GetBool("json")
			}
		}
	}

	if outputJSON {
		json, _ := utils.FormatJSON(client)
		cmd.Println(json)
	} else {
		cmd.Printf("👤 User Pool Client Details:\n")
		cmd.Printf("Client Name: %s\n", *client.ClientName)
		cmd.Printf("Client ID: %s\n", *client.ClientId)
		cmd.Printf("User Pool ID: %s\n", *client.UserPoolId)
		
		if client.ClientSecret != nil {
			cmd.Printf("Client Secret: %s\n", *client.ClientSecret)
		}
		
		if client.CreationDate != nil {
			cmd.Printf("Creation Date: %s\n", client.CreationDate.Format("2006-01-02 15:04:05"))
		}
		
		if client.LastModifiedDate != nil {
			cmd.Printf("Last Modified: %s\n", client.LastModifiedDate.Format("2006-01-02 15:04:05"))
		}
		
		if len(client.ExplicitAuthFlows) > 0 {
			authFlows := make([]string, len(client.ExplicitAuthFlows))
			for i, flow := range client.ExplicitAuthFlows {
				authFlows[i] = string(flow)
			}
			cmd.Printf("Auth Flows: %s\n", strings.Join(authFlows, ", "))
		}
		
		if len(client.SupportedIdentityProviders) > 0 {
			cmd.Printf("Identity Providers: %s\n", strings.Join(client.SupportedIdentityProviders, ", "))
		}
		
		if len(client.CallbackURLs) > 0 {
			cmd.Printf("Callback URLs: %s\n", strings.Join(client.CallbackURLs, ", "))
		}
		
		if len(client.LogoutURLs) > 0 {
			cmd.Printf("Logout URLs: %s\n", strings.Join(client.LogoutURLs, ", "))
		}
		
		if client.RefreshTokenValidity != 0 {
			cmd.Printf("Refresh Token Validity: %d days\n", client.RefreshTokenValidity)
		}
		
		if client.AccessTokenValidity != nil && *client.AccessTokenValidity != 0 {
			cmd.Printf("Access Token Validity: %d minutes\n", *client.AccessTokenValidity)
		}
		
		if client.IdTokenValidity != nil && *client.IdTokenValidity != 0 {
			cmd.Printf("ID Token Validity: %d minutes\n", *client.IdTokenValidity)
		}
	}
	
	return nil
}

