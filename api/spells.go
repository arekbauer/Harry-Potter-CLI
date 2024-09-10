package api

import (
	"encoding/json"
	"os"
	"strings"
	"time"

	"golang.org/x/exp/rand"
)

const cacheSpellFile = "cache/spells.json" // CHANGE TO ABSOLUTE

type Spell struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Function that grabs spells data from the Harry Potter API
func getSpells() {
	// Call the request
	responseBody, _ := getResponse("spells")

	// Write data to json file
	err := os.WriteFile(cacheSpellFile, responseBody, 0644)
	if err != nil {
		panic(err)
	}

}

// Function that checks cached files, creates cached file, and returns the spell list
func readSpells() []Spell {
	// Check if cache spells exists
	if !fileExists(cacheSpellFile) {
		getSpells()
	}

	// Read json file contents
	file, err := os.ReadFile(cacheSpellFile)
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

// Function that returns a certain spell and description based on index
func GetSpell(index int) (string, string) {
	// Grab the whole spell list
	spells := readSpells()

	return spells[index].Name, spells[index].Description
}

// Function that returns a spell index, if it exists, and 0 if not
func LookupSpell(search string) int {
	// Grab the whole spell list
	spells := readSpells()

	// Convert search string to lowercase for case-insensitive matching
	search = strings.ToLower(search)

	//Loop through the spells list until we find a match
	for i := 0; i < len(spells); i++ {
		if strings.ToLower(spells[i].Name) == search {
			return i
		}
	}
	return -1
}
