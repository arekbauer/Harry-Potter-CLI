package cmd

import (
	"hogwarts/api"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var charLookupCmd = &cobra.Command{
	Use:   "lookup",
	Short: "Get info on a specific Harry Potter character",
	Long: `Get info of a specific character from the Wizarding World.
Using: 
	
hogwarts character lookup --name Harry Potter`,
	Run: getLookupChar,
}

func init() {
	rootCmd.AddCommand(charLookupCmd)

	charLookupCmd.Flags().StringP("name", "n", "", "Name of character you wish you search")
}

// Format wanted Expelliarmus: Disarms an opponent.
func getLookupChar(cmd *cobra.Command, args []string) {
	// Code that is executed when the command is ran
	name, _ := cmd.Flags().GetString("name")

	// Look up character index of name given
	characterIndex := api.LookupCharacter(name)

	// If character does not exist, print red err msg
	if characterIndex <= -1 {
		color.New(color.FgRed, color.Bold).Println("That character does not exist!")
	} else {
		characterItem := api.GetCharacter(characterIndex)
		characterPrinter(characterItem) //call the printer function
	}
}
