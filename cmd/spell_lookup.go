package cmd

import (
	"hogwarts/api"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var spellLookupCmd = &cobra.Command{
	Use:   "lookup",
	Short: "Get info on a specific Harry Potter spell",
	Long: `Get info of a specific spell from the Wizarding World.
Using: 
	
hogwarts spell lookup --name Accio`,
	Run: getLookupSpell,
}

func init() {
	rootCmd.AddCommand(spellLookupCmd)

	spellLookupCmd.Flags().StringP("name", "n", "", "Name of spell you wish you search")
}

// Format wanted Expelliarmus: Disarms an opponent.
func getLookupSpell(cmd *cobra.Command, args []string) {
	// Code that is executed when the command is ran
	name, _ := cmd.Flags().GetString("name")

	// Look up spell index of name given
	spellIndex := api.LookupSpell(name)

	// If spell does not exist, print red err msg
	if spellIndex <= -1 {
		color.New(color.FgRed, color.Bold).Println("A spell with that name does not exist!")
	} else {
		spellName, spellDesc := api.GetSpell(spellIndex)
		color.New(color.FgWhite, color.Bold).Printf("%s: %s", spellName, spellDesc)
	}
}
