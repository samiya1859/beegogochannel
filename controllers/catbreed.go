package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/beego/beego/config"
	"github.com/beego/beego/v2/server/web"
)

type BreedController struct {
	web.Controller
}

// Breed represents the structure for the breed name and its ID
type BreedName struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// Breed struct for breed details
type Breed struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Origin      string `json:"origin"`
}

// BreedImage represents the breed image response
type BreedImage struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

// FetchBreedImages fetches images for a specific breed asynchronously
func FetchBreedImages(breedID, apiKey string, imageChannel chan<- []BreedImage, errorChannel chan<- error) {
	url := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?breed_ids=%s&limit=8", breedID)

	// Create HTTP client and request
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		errorChannel <- err
		return
	}

	// Set API key header
	req.Header.Set("x-api-key", apiKey)

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		errorChannel <- err
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errorChannel <- err
		return
	}

	// Parse JSON response
	var images []BreedImage
	err = json.Unmarshal(body, &images)
	if err != nil {
		errorChannel <- err
		return
	}

	// Send images to channel
	imageChannel <- images
}

// FetchBreedDetails fetches details for a specific breed asynchronously
func FetchBreedDetails(breedID, apiKey string, breedChannel chan<- Breed, errorChannel chan<- error) {
	url := fmt.Sprintf("https://api.thecatapi.com/v1/breeds/search?q=%s&attach_image=1", breedID)

	// Create HTTP client and request
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		errorChannel <- err
		return
	}

	// Set API key header
	req.Header.Set("x-api-key", apiKey)

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		errorChannel <- err
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errorChannel <- err
		return
	}

	// Parse JSON response
	var breeds []Breed
	err = json.Unmarshal(body, &breeds)
	if err != nil {
		errorChannel <- err
		return
	}

	if len(breeds) == 0 {
		errorChannel <- fmt.Errorf("no breed found for ID: %s", breedID)
		return
	}

	// Send breed data to channel
	breedChannel <- breeds[0]
}

// GetBreedImages handles fetching images for a specific breed
func (c *BreedController) GetBreedImages() {
	// Load configuration from app.conf
	cfg, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		log.Fatal("Error loading app.conf file")
	}

	// Get the API key from app.conf
	apiKey := cfg.String("CAT_api")
	if apiKey == "" {
		c.Ctx.WriteString("API Key is missing")
		return
	}

	// Fetch breed ID from the URL path
	breedID := c.Ctx.Input.Param(":breed_id") // Use Param(":breed_id") for URL path parameters
	if breedID == "" {
		c.Ctx.WriteString("Breed ID is required")
		return
	}

	// Create channels for breed images and errors
	imageChannel := make(chan []BreedImage)
	errorChannel := make(chan error)

	// Start goroutine to fetch breed images
	go FetchBreedImages(breedID, apiKey, imageChannel, errorChannel)

	// Wait for the result or an error
	select {
	case images := <-imageChannel:
		// Return breed images as JSON
		c.Data["json"] = images
		c.ServeJSON()
	case err := <-errorChannel:
		// Return error as JSON response
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
	}
}

// GetBreedDetails handles fetching details for a specific breed
func (c *BreedController) GetBreedDetails() {
	// Load configuration from app.conf
	cfg, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		log.Fatal("Error loading app.conf file")
	}

	apiKey := cfg.String("CAT_api")
	if apiKey == "" {
		c.Ctx.WriteString("API Key is missing")
		return
	}

	breedID := c.GetString(":breed_id")
	if breedID == "" {
		c.Ctx.WriteString("Breed ID is required")
		return
	}

	// Create channels for breed details and errors
	breedChannel := make(chan Breed)
	errorChannel := make(chan error)

	// Start a goroutine to fetch breed details
	go FetchBreedDetails(breedID, apiKey, breedChannel, errorChannel)

	// Wait for the result or an error
	select {
	case breed := <-breedChannel:
		// Return breed details as JSON
		c.Data["json"] = breed
		c.ServeJSON()
	case err := <-errorChannel:
		// Return error as JSON response
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
	}
}

// FetchAllBreeds fetches the list of all breeds from TheCatAPI
func FetchAllBreeds() ([]BreedName, error) {
	const apiURL = "https://api.thecatapi.com/v1/breeds"

	// Set a timeout for the request
	client := &http.Client{Timeout: 10 * time.Second}

	// Fetch all breeds
	resp, err := client.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch breeds: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected response status: %v", resp.StatusCode)
	}

	// Decode breeds response
	var breedList []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&breedList)
	if err != nil {
		return nil, fmt.Errorf("failed to parse breeds data: %v", err)
	}

	// Create a slice of breed names and IDs
	var breeds []BreedName
	for _, breed := range breedList {
		if name, ok := breed["name"].(string); ok {
			breeds = append(breeds, BreedName{
				Name: name,
				ID:   breed["id"].(string),
			})
		}
	}

	return breeds, nil
}

func (c *BreedController) GetAllBreeds() {
	// Fetch all breeds from the API or database
	breeds, err := FetchAllBreeds()
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	// Return the list of breeds as JSON
	c.Data["json"] = breeds
	c.ServeJSON()
}
