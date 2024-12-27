package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"catproject/controllers"

	"github.com/beego/beego/v2/server/web"
	"github.com/stretchr/testify/assert"
)

func init() {

	_, file, _, _ := runtime.Caller(0)
	appPath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))

	// Initialize Beego
	web.TestBeegoInit(appPath)

	// Register your routes
	web.Router("/v1/fetchcat", &controllers.CatAPIController{}, "get:Get")
	web.Router("/v1/vote", &controllers.CatAPIController{}, "post:Vote")
	web.Router("/v1/favourite", &controllers.CatAPIController{}, "post:Favourite")
	web.Router("/v1/favorites", &controllers.CatAPIController{}, "get:GetFavorites")
}

// MockCatAPIResponse mocks the Cat API response for testing
var mockCatAPIResponse = `[{"id":"abc123","url":"https://example.com/cat.jpg"}]`

// TestCatAPIController_Get tests the Get method
func TestCatAPIController_Get(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/fetchcat", nil)
	w := httptest.NewRecorder()
	web.BeeApp.Handlers.ServeHTTP(w, r)

	assert.Equal(t, 200, w.Code, "Expected status code 200, got %d. Response: %s", w.Code, w.Body.String())
}

// TestCatAPIController_Vote tests the Vote method
func TestCatAPIController_Vote(t *testing.T) {
	tests := []struct {
		name       string
		payload    controllers.VoteRequest
		wantStatus int
		wantError  bool
	}{
		{
			name: "Valid Vote",
			payload: controllers.VoteRequest{
				ImageID: "test123",
				SubID:   "user123",
				Value:   1,
			},
			wantStatus: 200,
			wantError:  false,
		},
		{
			name: "Invalid Vote Value",
			payload: controllers.VoteRequest{
				ImageID: "test123",
				SubID:   "user123",
				Value:   3, // Invalid value
			},
			wantStatus: 400,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonPayload, _ := json.Marshal(tt.payload)
			r, _ := http.NewRequest("POST", "/v1/vote", bytes.NewBuffer(jsonPayload))
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			web.BeeApp.Handlers.ServeHTTP(w, r)

			assert.Equal(t, tt.wantStatus, w.Code, "Test '%s' failed. Expected status %d, got %d. Response: %s",
				tt.name, tt.wantStatus, w.Code, w.Body.String())
		})
	}
}

// TestCatAPIController_Favourite tests the Favourite method
func TestCatAPIController_Favourite(t *testing.T) {
	payload := controllers.FavourRequest{
		ImageID: "test123",
		SubID:   "user123",
	}

	jsonPayload, _ := json.Marshal(payload)
	r, _ := http.NewRequest("POST", "/v1/favourite", bytes.NewBuffer(jsonPayload))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	web.BeeApp.Handlers.ServeHTTP(w, r)

	assert.Equal(t, 200, w.Code, "Expected status code 200, got %d. Response: %s", w.Code, w.Body.String())
}

// TestFetchRandomCatImage tests the FetchRandomCatImage function
func TestFetchRandomCatImage(t *testing.T) {
	imageID, imageURL, err := controllers.FetchRandomCatImage()
	if err != nil {
		t.Logf("Error occurred: %v", err)
	}

	// We can't assert exact values since it's random, but we can check if they're not empty
	assert.NotEmpty(t, imageID, "ImageID should not be empty")
	assert.NotEmpty(t, imageURL, "ImageURL should not be empty")
}
