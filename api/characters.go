package api

import (
	"encoding/json"
	"os"
	"strings"
	"time"

	"golang.org/x/exp/rand"
)

const cacheCharFile = "cache/character.json" // CHANGE TO ABSOLUTE

type Character struct {
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	House    string `json:"house"`
	Patronus string `json:"patronus"`
	Alive    bool   `json:"alive"`
}

// Function that grabs character data from the Harry Potter API
func getCharacters() {
	// Call the request
	responseBody, _ := getResponse("characters")

	// Write data to json file
	err := os.WriteFile(cacheCharFile, responseBody, 0644)
	if err != nil {
		panic(err)
	}

}

// Function that checks cached files, creates cached file, and returns the character list
func readCharacters() []Character {
	// Check if cache character file exists
	if !fileExists(cacheCharFile) {
		getCharacters()
	}

	// Read json file contents
	file, err := os.ReadFile(cacheCharFile)
	if err != nil {
		panic(err)
	}

	// Create a instance of our Character struct
	var characters []Character
	err = json.Unmarshal(file, &characters) //convert file into the Character format
	if err != nil {
		panic(err)
	}

	return characters
}

// Function that returns a random character info
func GetRandomCharacter() Character {
	rand.Seed(uint64(time.Now().UnixNano()))

	// Grab the whole characters list
	characters := readCharacters()
	numOfChars := len(characters)

	rndNum := rand.Intn(numOfChars)

	return characters[rndNum]
}

// Function that returns a certain character and description based on index
func GetCharacter(index int) Character {
	// Grab the whole characters list
	character := readCharacters()

	return character[index]
}

// Function that returns a character index, if it exists, and 0 if not
func LookupCharacter(search string) int {
	// Grab the whole characters list
	characters := readCharacters()

	// Convert search string to lowercase for case-insensitive matching
	search = strings.ToLower(search)

	//Loop through the characters list until we find a match
	for i := 0; i < len(characters); i++ {
		if strings.ToLower(characters[i].Name) == search {
			return i
		}
	}
	return -1
}

// Function that returns a list of all characters that fit the input parameters
func ListCharacters(house string, gender string, patronus string, alive bool) []int {
	var listOfValidChars []int

	// Grab the whole characters list
	characters := readCharacters()

	for i := 0; i < len(characters); i++ {
		if (house == "" || characters[i].House == house) &&
			(gender == "" || characters[i].Gender == gender) &&
			(patronus == "" || characters[i].Patronus == patronus) &&
			(characters[i].Alive == alive) {
			listOfValidChars = append(listOfValidChars, i)
		}
	}

	return listOfValidChars
}
