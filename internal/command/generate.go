package command

import (
	"fmt"

	"github.com/gdegiorgio/cognitools/internal/pkg"
	"github.com/gdegiorgio/cognitools/internal/service"
	"github.com/spf13/cobra"
)

func NewGenerateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "generate",
		Short: "Generate JWT",
		Run:   runGenerate,
	}
}

func runGenerate(cmd *cobra.Command, args []string) {

	svc := service.NewAWSService()

	pools, err := svc.ListPools()

	if err != nil {
		fmt.Printf("could not generate jwt: %v", err)
	}

	selectInput := []string{}

	for _, p := range pools {
		selectInput = append(selectInput, fmt.Sprintf("%s - %s", p.Name, p.PoolId))
	}

	idx, err := pkg.SelectUserPool(selectInput)

	if err != nil {
		fmt.Printf("failed to select user pool: %v", err)
	}

	clients, err := svc.ListClients(pools[idx].PoolId)

	if err != nil {
		fmt.Printf("failed to list clients: %v", err)
	}

	for _, c := range clients {
		selectInput = append(selectInput, fmt.Sprintf("%s - %s", c.Name, c.ClientId))
	}

	_, err = pkg.SelectClients(selectInput)

	if err != nil {
		fmt.Printf("failed to select client: %v", err)
	}

	scope, err := pkg.SelectScope(selectInput)

	if err != nil {
		fmt.Printf("failed to select scope: %v", err)
	}

	fmt.Printf("Selected scope: %s\n", scope)
}
