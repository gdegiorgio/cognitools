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
		cmd.Printf("could not generate jwt: %v", err)
	}

	selectInput := []string{}

	for _, p := range pools {
		selectInput = append(selectInput, fmt.Sprintf("%s - %s", p.Name, p.PoolId))
	}

	idx, err := pkg.SelectUserPool(selectInput)

	if err != nil {
		cmd.Printf("failed to select user pool: %v", err)
	}

	domain, err := svc.GetCognitoDomain(pools[idx].PoolId)

	clients, err := svc.ListClients(pools[idx].PoolId)

	if err != nil {
		cmd.Printf("failed to list clients: %v", err)
	}

	for _, c := range clients {
		selectInput = append(selectInput, fmt.Sprintf("%s - %s", c.Name, c.ClientId))
	}

	_, err = pkg.SelectClients(selectInput)

	if err != nil {
		cmd.Printf("failed to select client: %v", err)
	}

	secret, err := svc.GetCognitoClientSecret(pools[idx].PoolId, clients[idx].ClientId)

	scope, err := pkg.SelectScope(selectInput)

	if err != nil {
		cmd.Printf("failed to select scope: %v", err)
	}

	j, err := service.GenerateJWT(domain, clients[idx].ClientId, secret, scope)

	if err != nil {
		cmd.Printf("could not generate jwt: %v", err)
	}

	cmd.Printf("💡 JWT generated successfully. You can now use it in your application.\n\n\n")
	cmd.Println(j)
}
