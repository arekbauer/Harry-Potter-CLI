package cmd

import "github.com/spf13/cobra"

var spellCmd = &cobra.Command{
	Use:   "spell",
	Short: "Get a Harry Potter spell",
	Long: `Get a random spell from the Wizarding World, or get info of a specific spell.

For example: 
	
hogwarts spell random 
hogwarts spell lookup --name Accio`,
}

func init() {
	rootCmd.AddCommand(spellCmd)
	spellCmd.AddCommand(spellRandomCmd)
	spellCmd.AddCommand(spellLookupCmd)
}
