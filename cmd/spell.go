package cmd

import "github.com/spf13/cobra"

var spellCmd = &cobra.Command{
	Use:   "spell",
	Short: "Get a Harry Potter spell",
	Long: `Get a random spell from the Wizarding World, or find the use of a certain one.

For example: 
	
potter spell --random 
potter spell lookup`,
}

func init() {
	rootCmd.AddCommand(spellCmd)
	spellCmd.AddCommand(spellRandomCmd)
}
