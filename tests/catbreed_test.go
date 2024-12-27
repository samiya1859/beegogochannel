package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"catproject/controllers" // Replace with your actual project name
	"github.com/beego/beego/v2/server/web"
	"github.com/stretchr/testify/assert"
)

// Mock responses
var mockBreedListResponse = `[
    {"id": "abys", "name": "Abyssinian"},
    {"id": "beng", "name": "Bengal"}
]`

var mockBreedDetailsResponse = `[{
    "id": "abys",
    "name": "Abyssinian",
    "description": "The Abyssinian is easy to care for",
    "origin": "Egypt"
}]`

var mockBreedImagesResponse = `[
    {"id": "img1", "url": "https://example.com/cat1.jpg"},
    {"id": "img2", "url": "https://example.com/cat2.jpg"}
]`

func init() {
	_, file, _, _ := runtime.Caller(0)
	appPath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	web.TestBeegoInit(appPath)

	// Register routes for testing
	web.Router("/v1/breeds", &controllers.BreedController{}, "get:GetAllBreeds")
	web.Router("/v1/breeds/:breed_id", &controllers.BreedController{}, "get:GetBreedDetails")
	web.Router("/v1/breeds/:breed_id/images", &controllers.BreedController{}, "get:GetBreedImages")
}

// TestGetAllBreeds tests the GetAllBreeds endpoint
func TestGetAllBreeds(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/breeds", nil)
	w := httptest.NewRecorder()
	web.BeeApp.Handlers.ServeHTTP(w, r)

	assert.Equal(t, 200, w.Code, "Expected status code 200")

	var response []controllers.BreedName
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Should be able to parse response")
	assert.NotEmpty(t, response, "Response should not be empty")
}

// TestGetBreedDetails tests the GetBreedDetails endpoint
func TestGetBreedDetails(t *testing.T) {
	tests := []struct {
		name       string
		breedID    string
		wantStatus int
		wantError  bool
	}{
		{
			name:       "Valid Breed ID",
			breedID:    "abys",
			wantStatus: 200,
			wantError:  false,
		},
		{
			name:       "Invalid Breed ID",
			breedID:    "",
			wantStatus: 400,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := http.NewRequest("GET", "/v1/breeds/"+tt.breedID, nil)
			w := httptest.NewRecorder()
			web.BeeApp.Handlers.ServeHTTP(w, r)

			assert.Equal(t, tt.wantStatus, w.Code, "Test '%s' failed: wrong status code", tt.name)

			if !tt.wantError {
				var response controllers.Breed
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err, "Should be able to parse response")
				assert.Equal(t, tt.breedID, response.ID, "Should return correct breed ID")
			}
		})
	}
}

// TestGetBreedImages tests the GetBreedImages endpoint
func TestGetBreedImages(t *testing.T) {
	tests := []struct {
		name       string
		breedID    string
		wantStatus int
		wantError  bool
	}{
		{
			name:       "Valid Breed ID",
			breedID:    "abys",
			wantStatus: 200,
			wantError:  false,
		},
		{
			name:       "Invalid Breed ID",
			breedID:    "",
			wantStatus: 400,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := http.NewRequest("GET", "/v1/breeds/"+tt.breedID+"/images", nil)
			w := httptest.NewRecorder()
			web.BeeApp.Handlers.ServeHTTP(w, r)

			assert.Equal(t, tt.wantStatus, w.Code, "Test '%s' failed: wrong status code", tt.name)

			if !tt.wantError {
				var response []controllers.BreedImage
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err, "Should be able to parse response")
				assert.NotEmpty(t, response, "Should return breed images")
			}
		})
	}
}

// TestFetchBreedImages tests the FetchBreedImages function
func TestFetchBreedImages(t *testing.T) {
	imageChannel := make(chan []controllers.BreedImage)
	errorChannel := make(chan error)

	go controllers.FetchBreedImages("abys", "test-api-key", imageChannel, errorChannel)

	select {
	case images := <-imageChannel:
		assert.NotEmpty(t, images, "Should receive breed images")
	case err := <-errorChannel:
		t.Errorf("Unexpected error: %v", err)
	}
}

// TestFetchBreedDetails tests the FetchBreedDetails function
func TestFetchBreedDetails(t *testing.T) {
	breedChannel := make(chan controllers.Breed)
	errorChannel := make(chan error)

	go controllers.FetchBreedDetails("abys", "test-api-key", breedChannel, errorChannel)

	select {
	case breed := <-breedChannel:
		assert.Equal(t, "abys", breed.ID, "Should receive correct breed details")
	case err := <-errorChannel:
		t.Errorf("Unexpected error: %v", err)
	}
}

// TestFetchAllBreeds tests the FetchAllBreeds function
func TestFetchAllBreeds(t *testing.T) {
	breeds, err := controllers.FetchAllBreeds()
	assert.NoError(t, err, "Should fetch breeds without error")
	assert.NotEmpty(t, breeds, "Should return non-empty breed list")
}
