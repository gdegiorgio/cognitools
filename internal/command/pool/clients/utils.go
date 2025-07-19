package clients

import (
	"github.com/spf13/cobra"
)

// getJSONFlag checks for JSON output flag, looking at current command and parent commands
func getJSONFlag(cmd *cobra.Command) bool {
	// First check if the current command has the flag set
	if cmd.Flags().Changed("json") {
		outputJSON, _ := cmd.Flags().GetBool("json")
		return outputJSON
	}
	
	// Check parent commands for the flag
	current := cmd
	for current != nil {
		if current.Flags().Lookup("json") != nil {
			outputJSON, _ := current.Flags().GetBool("json")
			return outputJSON
		}
		current = current.Parent()
	}
	
	return false
}