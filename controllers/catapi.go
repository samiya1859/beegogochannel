package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"
	"time"

	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/server/web"
)

// CatAPIController handles API requests for cat images
type CatAPIController struct {
	web.Controller
}

// Define the VoteRequest struct
type VoteRequest struct {
	ImageID string `json:"image_id"`
	SubID   string `json:"sub_id"`
	Value   int    `json:"value"`
}
type FavourRequest struct {
	ImageID string `json:"image_id"`
	SubID   string `json:"sub_id"`
}
type FavoriteResponse struct {
	Image struct {
		ID  string `json:"id"`
		URL string `json:"url"`
	} `json:"image"`
}

// FetchRandomCatImage fetches a random cat image URL from the API
func FetchRandomCatImage() (string, string, error) {
	const apiURL = "https://api.thecatapi.com/v1/images/search?size=med&mime_types=jpg&format=json&has_breeds=true&order=RANDOM&page=0&limit=1"

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(apiURL)
	if err != nil {
		return "", "", fmt.Errorf("failed to fetch cat image: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", "", fmt.Errorf("unexpected response status: %v", resp.StatusCode)
	}

	var result []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse response: %v", err)
	}

	// Extract image ID and URL from the response
	imageID, ok1 := result[0]["id"].(string)
	imageURL, ok2 := result[0]["url"].(string)
	if !ok1 || !ok2 {
		return "", "", fmt.Errorf("image ID or URL not found in response")
	}

	return imageID, imageURL, nil
}

// Get handles GET requests for /fetchcat and renders the page
func (c *CatAPIController) Get() {
	catChannel := make(chan struct {
		ImageID  string
		ImageURL string
	}, 1)
	errorChannel := make(chan error, 1)

	go func() {
		imageID, imageURL, err := FetchRandomCatImage()
		if err != nil {
			errorChannel <- err
		} else {
			catChannel <- struct {
				ImageID  string
				ImageURL string
			}{ImageID: imageID, ImageURL: imageURL}
		}
	}()

	select {
	case catData := <-catChannel:
		c.Data["ImageID"] = catData.ImageID
		c.Data["ImageURL"] = catData.ImageURL
		c.TplName = "home.tpl"
		close(catChannel)
		close(errorChannel)
	case err := <-errorChannel:
		c.Data["Error"] = "Error fetching cat image: " + err.Error()
		c.TplName = "error.tpl"
		close(catChannel)
		close(errorChannel)
	}
}

func (c *CatAPIController) Vote() {
	body, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.WriteString("Error reading request body")
		return
	}

	// Print the raw body to check if it's received
	fmt.Println("Raw Request Body:", string(body))

	var voteRequest VoteRequest

	// Parse the incoming JSON payload
	err = json.Unmarshal(body, &voteRequest)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		c.Ctx.WriteString(fmt.Sprintf("Error parsing JSON: %s", err.Error()))
		return
	}
	// Log the parsed request
	fmt.Println("Received Vote:", voteRequest)
	// Validate required fields
	if voteRequest.ImageID == "" || voteRequest.SubID == "" {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		c.Ctx.WriteString("Missing required fields")
		return
	}

	// Validate the vote value
	if voteRequest.Value != 1 && voteRequest.Value != 2 {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		c.Ctx.WriteString("Invalid vote value")
		return
	}

	// Channel to handle the result of the vote submission
	resultChannel := make(chan string)

	// Use Go routine to handle the vote submission asynchronously
	go func() {

		voteAPIURL := "https://api.thecatapi.com/v1/votes"
		voteAPIKey := "live_aEIcWNMCOCKINdpArjjNnm54ivWft2t1E2ZiBWTPsjuhWXPjq5Ih8NhFzZUqzwHW"

		// Create the vote data in the correct format
		voteData := map[string]interface{}{
			"image_id": voteRequest.ImageID,
			"sub_id":   voteRequest.SubID,
			"value":    voteRequest.Value,
		}

		// Convert the vote data to JSON
		requestBody, err := json.Marshal(voteData)
		if err != nil {
			resultChannel <- fmt.Sprintf("Error marshaling vote data: %s", err.Error())
			return
		}

		// Send the request to the Cat API to register the vote
		client := &http.Client{}
		req, err := http.NewRequest("POST", voteAPIURL, bytes.NewBuffer(requestBody))
		if err != nil {
			resultChannel <- fmt.Sprintf("Error creating HTTP request: %s", err.Error())
			return
		}

		// Add the authorization header with the API key
		req.Header.Set("x-api-key", voteAPIKey)
		req.Header.Set("Content-Type", "application/json")

		// Perform the request
		resp, err := client.Do(req)
		if err != nil {
			resultChannel <- fmt.Sprintf("Error sending request to Cat API: %s", err.Error())
			return
		}
		defer resp.Body.Close()

		// Check if the response is successful
		if resp.StatusCode != 200 && resp.StatusCode != 201 {
			resultChannel <- fmt.Sprintf("Error: Cat API returned status %d", resp.StatusCode)
			return
		}

		// Send success message to the channel
		resultChannel <- "Vote successfully submitted to the Cat API!"
	}()

	// Wait for the result from the channel
	result := <-resultChannel

	// Respond with the result message
	c.Ctx.WriteString(result)

	// Close the channel after processing
	close(resultChannel)
}

// Favourite handles the POST request for adding a cat image to favourites
func (c *CatAPIController) Favourite() {
	favbody, faverr := ioutil.ReadAll(c.Ctx.Request.Body)
	if faverr != nil {
		fmt.Println("Error reading body:", faverr)
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.WriteString("Error reading request body")
		return
	}

	// Print the raw body to check if it's received
	fmt.Println("Raw Request Body:", string(favbody))

	var favourRequest FavourRequest

	// Parse the incoming JSON payload
	err := json.Unmarshal(favbody, &favourRequest)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		c.Ctx.WriteString(fmt.Sprintf("Error parsing JSON: %s", err.Error()))
		return
	}

	// Log the parsed request
	fmt.Println("Received Favourite Request:", favourRequest)

	// Validate required fields
	if favourRequest.ImageID == "" || favourRequest.SubID == "" {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		c.Ctx.WriteString("Missing required fields")
		return
	}

	resultChannel := make(chan string)

	go func() {
		favouriteAPIURL := "https://api.thecatapi.com/v1/favourites"
		favouriteAPIKey := "live_aEIcWNMCOCKINdpArjjNnm54ivWft2t1E2ZiBWTPsjuhWXPjq5Ih8NhFzZUqzwHW"

		// Create the favourite data in the correct format
		favouriteData := map[string]interface{}{
			"image_id": favourRequest.ImageID,
			"sub_id":   favourRequest.SubID,
		}

		// Convert the favourite data to JSON
		requestBody, err := json.Marshal(favouriteData)
		if err != nil {
			resultChannel <- fmt.Sprintf("Error marshaling favourite data: %s", err.Error())
			return
		}

		client := &http.Client{}
		req, err := http.NewRequest("POST", favouriteAPIURL, bytes.NewBuffer(requestBody))
		if err != nil {
			resultChannel <- fmt.Sprintf("Error creating HTTP request: %s", err.Error())
			return
		}

		req.Header.Set("x-api-key", favouriteAPIKey)
		req.Header.Set("Content-Type", "application/json")

		// Perform the request
		resp, err := client.Do(req)
		if err != nil {
			resultChannel <- fmt.Sprintf("Error sending request to Cat API: %s", err.Error())
			return
		}
		defer resp.Body.Close()

		// Check if the response is successful
		if resp.StatusCode != 200 {
			resultChannel <- fmt.Sprintf("Error: Cat API returned status %d", resp.StatusCode)
			return
		}

		// Send success message to the channel
		resultChannel <- "Image successfully added to favourites!"
	}()

	// Wait for the result from the channel
	result := <-resultChannel

	// Send structured JSON response to the client
	c.Data["json"] = map[string]interface{}{"status": "success", "message": result}
	c.ServeJSON()

	// Close the channel after processing
	close(resultChannel)
}

func (c *CatAPIController) GetFavorites() {
	// Load the API key from configuration
	apiKey, err := config.String("CAT_api")
	if err != nil || apiKey == "" {
		c.Ctx.WriteString("API Key is missing")
		return
	}

	// Prepare the API request to get the user's favorites
	req, err := http.NewRequest("GET", "https://api.thecatapi.com/v1/favourites", nil)
	if err != nil {
		c.Ctx.WriteString("Error preparing the request")
		return
	}

	// Set the API Key header
	req.Header.Set("x-api-key", apiKey)

	// Make the API call
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Ctx.WriteString("Error fetching favorites")
		return
	}
	defer resp.Body.Close()

	// Read and parse the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Ctx.WriteString("Error reading response")
		return
	}

	var favorites []struct {
		Image struct {
			ID  string `json:"id"`
			URL string `json:"url"`
		} `json:"image"`
	}
	err = json.Unmarshal(body, &favorites)
	if err != nil {
		c.Ctx.WriteString("Error unmarshalling JSON")
		return
	}

	// Pass the favorite images to the template
	if len(favorites) > 0 {
		c.Data["Favorites"] = favorites
	} else {
		c.Data["Favorites"] = nil
	}

	c.TplName = "home.tpl"
}
