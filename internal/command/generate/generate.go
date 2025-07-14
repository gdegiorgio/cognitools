package generate

import (
	"fmt"

	"github.com/gdegiorgio/cognitools/internal/service"
	"github.com/gdegiorgio/cognitools/internal/ui"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "generate",
		Short: "Generate JWT",
		Long: `Generate JWT tokens for a selected AWS Cognito User Pool and Client.
	
	This command lists available Cognito User Pools, prompts you to select one,
	then lists its clients and scopes, and finally generates a JWT token that
	can be used to authenticate against your application.`,
		Run: runGenerate,
	}
}

func runGenerate(cmd *cobra.Command, args []string) {
	generate(cmd, args, service.NewAWSService(), ui.NewPrompt(), service.NewAuthService())
}

func generate(cmd *cobra.Command, args []string, svc service.AWS, prompt ui.Prompt, auth service.Auth) {

	pools, err := svc.ListPools()
	if err != nil {
		cmd.Printf("❌ could not list user pools: %v\n", err)
		return
	}

	selectInput := make([]string, len(pools))
	for i, p := range pools {
		selectInput[i] = fmt.Sprintf("%s - %s", p.Name, p.PoolId)
	}

	idx, err := prompt.SelectFromList("🏖️ Select Cognito User Pool", selectInput)

	if err != nil {
		cmd.Printf("❌ failed to select user pool: %v\n", err)
		return
	}

	pool := pools[idx]

	domain, err := svc.GetCognitoDomain(pool.PoolId)

	if err != nil {
		cmd.Printf("❌ failed to get Cognito domain: %v\n", err)
		return
	}

	clients, err := svc.ListClients(pool.PoolId)
	if err != nil {
		cmd.Printf("❌ failed to list clients: %v\n", err)
		return
	}

	clientInput := make([]string, len(clients))
	for i, c := range clients {
		clientInput[i] = fmt.Sprintf("%s - %s", c.Name, c.ClientId)
	}

	clientIdx, err := prompt.SelectFromList("👤 Select Cognito Client", clientInput)

	if err != nil {
		cmd.Printf("❌ failed to select client: %v\n", err)
		return
	}

	selectedClient := clients[clientIdx]

	secret, err := svc.GetCognitoClientSecret(pool.PoolId, selectedClient.ClientId)
	if err != nil {
		cmd.Printf("❌ failed to get client secret: %v\n", err)
		return
	}

	scope, err := prompt.PromptInput("🎯 Please enter the Cognito Scope Name to use for request")
	if err != nil {
		cmd.Printf("❌ failed to select scope: %v\n", err)
		return
	}

	jwt, err := auth.GenerateJWT(domain, selectedClient.ClientId, secret, scope)
	if err != nil {
		cmd.Printf("❌ could not generate JWT: %v\n", err)
		return
	}

	cmd.Printf("✅ JWT generated successfully.\n\n%s\n", jwt)
}
