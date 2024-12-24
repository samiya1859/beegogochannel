package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"net/http"
	"time"
)

// CatAPIController handles API requests for cat images
type CatAPIController struct {
	web.Controller
}

// FetchRandomCatImage fetches a random cat image URL from the API
func FetchRandomCatImage() (string, error) {
	const apiURL = "https://api.thecatapi.com/v1/images/search?size=med&mime_types=jpg&format=json&has_breeds=true&order=RANDOM&page=0&limit=1"

	// Set a timeout for the request
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(apiURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch cat image: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("unexpected response status: %v", resp.StatusCode)
	}

	// Decode the JSON response
	var result []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", fmt.Errorf("failed to parse response: %v", err)
	}

	// Extract the image URL from the response
	imageURL, ok := result[0]["url"].(string)
	if !ok {
		return "", fmt.Errorf("image URL not found in response")
	}

	return imageURL, nil
}

// Get handles GET requests for /fetchcat and renders the page
func (c *CatAPIController) Get() {
	// Create channels for the image URL and error handling
	catChannel := make(chan string, 1)
	errorChannel := make(chan error, 1)

	// Fetch the random cat image asynchronously using goroutines
	go func() {
		imageURL, err := FetchRandomCatImage()
		if err != nil {
			errorChannel <- err
		} else {
			catChannel <- imageURL
		}
	}()

	// Wait for the result from either the catChannel or errorChannel
	select {
	case imageURL := <-catChannel:
		// Successfully fetched the image, render the template
		c.Data["ImageURL"] = imageURL
		c.TplName = "home.tpl"
		close(catChannel) // Close the channel after use
		close(errorChannel)
	case err := <-errorChannel:
		// If there was an error, handle it gracefully
		c.Data["Error"] = "Error fetching cat image: " + err.Error()
		c.TplName = "error.tpl"
		close(catChannel)
		close(errorChannel)
	}
}
