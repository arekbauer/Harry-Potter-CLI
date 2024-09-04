package api

import (
	"io"
	"net/http"
	"os"
)

var URL string = "https://hp-api.onrender.com"

func getResponse(dataType string) ([]byte, error) {
	URL = URL + "/api/" + dataType

	response, err := http.Get(URL)
	// Check if we get err
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	// Check if we got a valid response
	if response.StatusCode != 200 {
		panic("Harry Potter API not available")
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	return responseBody, nil
}

// Returns True if file exists
func fileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}
