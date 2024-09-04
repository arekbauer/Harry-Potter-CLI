package cmd

import (
	"fmt"

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

func getRandomSpell(cmd *cobra.Command, args []string) {
	// Code that is executed when the command is ran
	fmt.Println("Name: Accio")
	fmt.Println("Description: Moves things towards you")
}
