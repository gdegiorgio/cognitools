package command

import "github.com/spf13/cobra"

func NewRootCommand() *cobra.Command {
	root := &cobra.Command{
		Use:   "cognitools",
		Short: "Cognito Tools CLI",
	}

	root.AddCommand(NewGenerateCommand())

	return root
}
