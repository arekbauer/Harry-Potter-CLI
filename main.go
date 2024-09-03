package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Spell struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func main() {
	var q string = "spells"
	const url string = "https://hp-api.onrender.com"

	response, err := http.Get(url + "/api/" + q)
	// Check if we get err
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	// Check if we got a valid response
	if response.StatusCode != 200 {
		panic("Harry Potter Spells API not available")
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// Create a instance of our Spell struct
	var spells []Spell
	err = json.Unmarshal(body, &spells) //convert body into the Spell format
	if err != nil {
		panic(err)
	}

	// Can save info to variables, but need to specify which spell
	spellName, spellDesc := spells[1].Name, spells[1].Description

	fmt.Printf(
		"%s - %s\n",
		spellName,
		spellDesc,
	)
}
