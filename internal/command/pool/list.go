package pool

import (
	"github.com/gdegiorgio/cognitools/internal/service"
	"github.com/gdegiorgio/cognitools/internal/utils"
	"github.com/spf13/cobra"
)

func newListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List User Pools",
		Long:  "List User Pools retrieves and displays all user pools in your application.",
		RunE:  runListCommand,
	}
}

func runListCommand(cmd *cobra.Command, args []string) error {
	return list(cmd, args, service.NewAWSService())
}

func list(cmd *cobra.Command, args []string, svc service.AWS) error {

	pools, err := svc.ListUsersPools()

	if err != nil {
		cmd.Printf("❌ could not list user pools: %v\n", err)
		return err
	}

	if len(pools) == 0 {
		cmd.Println("❌ No user pools found.")
		return nil
	}

	if outputJSON {
		json, _ := utils.FormatJSON(pools)
		cmd.Println(json)
	} else {
		for _, pool := range pools {
			cmd.Println("🏖️ User Pool:", *pool.Name, "-", *pool.Id)
		}
	}
	return nil
}
