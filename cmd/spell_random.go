package cmd

import (
	"hogwarts/api"

	"github.com/fatih/color"

	"github.com/spf13/cobra"
)

var spellRandomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random Harry Potter spell",
	Long: `Get a random spell from the Wizarding World.
Using: 
	
potter spell random`,
	Run: getRandomSpell,
}

// Format wanted Expelliarmus: Disarms an opponent.
func getRandomSpell(cmd *cobra.Command, args []string) {
	// Code that is executed when the command is ran
	spellName, spellDesc := api.GetRandomSpell()

	boldWhite := color.New(color.FgWhite, color.Bold)
	boldWhite.Printf("%s: %s", spellName, spellDesc)
}
