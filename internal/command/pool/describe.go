package pool

import (
	"github.com/gdegiorgio/cognitools/internal/service"
	"github.com/gdegiorgio/cognitools/internal/utils"
	"github.com/spf13/cobra"
)

func newDescribeCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "describe <pool-id>",
		Short: "Describe User Pool",
		Long:  "Describe User Pool retrieves and displays detailed information about a specific user pool.",
		Args:  cobra.ExactArgs(1),
		RunE:  runDescribeCommand,
	}
}

func runDescribeCommand(cmd *cobra.Command, args []string) error {
	return describe(cmd, args, service.NewAWSService())
}

func describe(cmd *cobra.Command, args []string, svc service.AWS) error {
	poolID := args[0]

	pool, err := svc.DescribeUserPool(poolID)
	if err != nil {
		cmd.Printf("❌ %s\n", formatAWSError(err))
		return err
	}

	if outputJSON {
		json, _ := utils.FormatJSON(pool)
		cmd.Println(json)
		return nil
	}

	cmd.Printf("🏖️ User Pool Details:\n")
	cmd.Printf("Name: %s\n", *pool.Name)
	cmd.Printf("ID: %s\n", *pool.Id)
	cmd.Printf("Creation Date: %s\n", pool.CreationDate.Format("2006-01-02 15:04:05"))
	cmd.Printf("Last Modified: %s\n", pool.LastModifiedDate.Format("2006-01-02 15:04:05"))

	if pool.Domain != nil {
		cmd.Printf("Domain: %s\n", *pool.Domain)
	}

	if pool.Policies != nil && pool.Policies.PasswordPolicy != nil {
		cmd.Printf("Password Policy:\n")
		if pool.Policies.PasswordPolicy.MinimumLength != nil {
			cmd.Printf("  Minimum Length: %d\n", *pool.Policies.PasswordPolicy.MinimumLength)
		}
		cmd.Printf("  Require Uppercase: %t\n", pool.Policies.PasswordPolicy.RequireUppercase)
		cmd.Printf("  Require Lowercase: %t\n", pool.Policies.PasswordPolicy.RequireLowercase)
		cmd.Printf("  Require Numbers: %t\n", pool.Policies.PasswordPolicy.RequireNumbers)
		cmd.Printf("  Require Symbols: %t\n", pool.Policies.PasswordPolicy.RequireSymbols)
	}

	if len(pool.AliasAttributes) > 0 {
		cmd.Printf("Alias Attributes: %v\n", pool.AliasAttributes)
	}

	if len(pool.UsernameAttributes) > 0 {
		cmd.Printf("Username Attributes: %v\n", pool.UsernameAttributes)
	}

	return nil
}