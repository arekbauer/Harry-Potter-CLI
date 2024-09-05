package api

import (
	"encoding/json"
	"os"
	"time"

	"golang.org/x/exp/rand"
)

const cacheCharFile = "cache/character.json"

type Character struct {
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	House    string `json:"house"`
	Patronus string `json:"patronus"`
	Student  bool   `json:"hogwartsStudent"`
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
	// Check if cache spells exists
	if !fileExists(cacheCharFile) {
		getCharacters()
	}

	// Read json file contents
	file, err := os.ReadFile(cacheCharFile)
	if err != nil {
		panic(err)
	}

	// Create a instance of our Spell struct
	var characters []Character
	err = json.Unmarshal(file, &characters) //convert file into the Spell format
	if err != nil {
		panic(err)
	}

	return characters
}

// Function that returns a random character info
func GetRandomCharacter() Character {
	rand.Seed(uint64(time.Now().UnixNano()))

	// Grab the whole spell list
	characters := readCharacters()
	numOfChars := len(characters)

	rndNum := rand.Intn(numOfChars)

	return characters[rndNum]
}
