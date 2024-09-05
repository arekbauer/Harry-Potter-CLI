package cmd

import (
	"hogwarts/api"

	"github.com/spf13/cobra"
)

var characterRandomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random Harry Potter character",
	Long: `Get a random character from the Wizarding World.
	Using: 
	
	hogwarts character random`,
	Run: getRandomChar,
}

// Format wanted Expelliarmus: Disarms an opponent.
func getRandomChar(cmd *cobra.Command, args []string) {
	// Code that is executed when the command is ran
	character := api.GetRandomCharacter()

	// Call the character printer to do the outputting
	characterPrinter(character)
}
