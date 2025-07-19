package command

import (
	"github.com/gdegiorgio/cognitools/internal/command/generate"
	"github.com/gdegiorgio/cognitools/internal/command/pool"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {

	root := &cobra.Command{
		Use:   "cognitools",
		Short: "Cognito Tools CLI",
	}

	root.AddCommand(generate.NewCommand())
	root.AddCommand(pool.NewCommand())

	return root
}
