package cmd

import (
	"hogwarts/api"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var charListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get info on specific Harry Potter characters",
	Long: `Get info of specific characters from the Wizarding World.
	Using: 
	
	hogwarts character list --house Ravenclaw`,
	Run: getListChar,
}

func init() {
	rootCmd.AddCommand(charListCmd)

	charListCmd.Flags().StringP("house", "f", "", "House of character you wish you search")
	charListCmd.Flags().StringP("gender", "g", "", "Gender of character you wish you search")
	charListCmd.Flags().StringP("patronus", "p", "", "Patronus of character you wish you search")
	charListCmd.Flags().StringP("student", "s", "", "Student status of the character you wish you search")
	charListCmd.Flags().StringP("alive", "a", "", "Life status of the character you wish you search")
}

// Format wanted Expelliarmus: Disarms an opponent.
func getListChar(cmd *cobra.Command, args []string) {
	// Code that is executed when the command is ran
	house, _ := cmd.Flags().GetString("house")
	gender, _ := cmd.Flags().GetString("gender")
	patronus, _ := cmd.Flags().GetString("patronus")
	student, _ := cmd.Flags().GetString("student")
	alive, _ := cmd.Flags().GetString("alive")

	// Look up character index of name given
	characterIndexList := api.ListCharacters(house, gender, patronus, student, alive)

	// If character does not exist, print red err msg
	if len(characterIndexList) == 0 {
		color.New(color.FgRed, color.Bold).Println("No characters with those filters exist!")
	} else {
		for i := 0; i < len(characterIndexList); i++ {
			characterItem := api.GetCharacter(characterIndexList[i])
			characterPrinter(characterItem) //call the printer function
		}

	}
}
