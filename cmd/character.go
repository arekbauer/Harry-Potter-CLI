package cmd

import (
	"fmt"
	"hogwarts/api"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var characterCmd = &cobra.Command{
	Use:   "character",
	Short: "Get a Harry Potter character",
	Long: `Get a random character from the Wizarding World, or get info of a specific character.

For example: 
	
hogwarts character random 
hogwarts character lookup --name "Harry Potter"
howarts character list --house Ravenclaw --gender male`,
}

func characterPrinter(character api.Character) {
	// Creating colours to use
	fields := color.New(color.FgWhite, color.Bold)
	defaultText := color.New(color.FgCyan, color.Bold)
	falseText := color.New(color.FgRed, color.Bold)
	trueText := color.New(color.FgGreen, color.Bold)
	hufflepuff := color.New(color.FgYellow, color.Bold)
	ravenclaw := color.New(color.FgBlue, color.Bold)
	gryffindor := color.New(color.FgRed, color.Bold)
	slytherin := color.New(color.FgGreen, color.Bold)

	// Print name + gender
	fields.Printf("Name: ")
	defaultText.Printf("%s\n", character.Name)
	fields.Printf("Gender: ")
	defaultText.Printf("%s\n", cases.Title(language.English).String(character.Gender))

	// Print House
	fields.Printf("House: ")
	switch character.House {
	case "Gryffindor":
		gryffindor.Printf("%s\n", character.House)
	case "Slytherin":
		slytherin.Printf("%s\n", character.House)
	case "Ravenclaw":
		ravenclaw.Printf("%s\n", character.House)
	case "Hufflepuff":
		hufflepuff.Printf("%s\n", character.House)
	default:
		falseText.Printf("Unknown house\n")
	}

	// Print patronus
	fields.Printf("Patronus: ")
	if character.Patronus == "" {
		falseText.Println("Unknown patronus")
	} else {
		defaultText.Printf("%s\n", cases.Title(language.English).String(character.Patronus))
	}

	// Print alive status
	fields.Printf("Status: ")
	if character.Alive {
		trueText.Println("Alive")
	} else {
		falseText.Println("Deceased")
	}

	fmt.Printf("\n")
}

func init() {
	rootCmd.AddCommand(characterCmd)
	characterCmd.AddCommand(characterRandomCmd)
	characterCmd.AddCommand(charLookupCmd)
	characterCmd.AddCommand(charListCmd)
}
