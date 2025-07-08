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

	fmt.Printf("You selected %s\n", pools[idx].Name)

}
