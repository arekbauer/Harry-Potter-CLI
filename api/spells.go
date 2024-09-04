package api

import (
	"encoding/json"
	"os"
	"time"

	"golang.org/x/exp/rand"
)

const cacheFile = "cache/spells.json"

type Spell struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Function that grabs spells data from the Harry Potter API
func getSpells() {
	// Call the request
	responseBody, _ := getResponse("spells")

	// Create a instance of our Spell struct
	var spells []Spell
	err := json.Unmarshal(responseBody, &spells) //convert responseBody into the Spell format
	if err != nil {
		panic(err)
	}

	// Write data to json file
	err = os.WriteFile(cacheFile, responseBody, 0644)
	if err != nil {
		panic(err)
	}

}

// Function that checks cached files, creates cached file, and returns the spell list
func readSpells() []Spell {
	// Check if cache spells exists
	if !fileExists(cacheFile) {
		getSpells()
	}

	// Read json file contents
	file, err := os.ReadFile(cacheFile)
	if err != nil {
		panic(err)
	}

	// Create a instance of our Spell struct
	var spells []Spell
	err = json.Unmarshal(file, &spells) //convert file into the Spell format
	if err != nil {
		panic(err)
	}

	return spells
}

// Function that returns a random spell name and description
func GetRandomSpell() (string, string) {
	rand.Seed(uint64(time.Now().UnixNano()))

	// Grab the whole spell list
	spells := readSpells()
	numOfSpells := len(spells)

	rndNum := rand.Intn(numOfSpells)

	return spells[rndNum].Name, spells[rndNum].Description
}

func LookupSpell() {

}
