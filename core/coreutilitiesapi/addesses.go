package coreutilitiesapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

// makeRequest performs HTTP request to CSC API
func makeRequest(endpoint string) ([]byte, error) {
	apiKey := os.Getenv("CITY_STATE_API_KEY")
	baseURL := os.Getenv("CITY_STATE_BASE_URL")

	if apiKey == "" || baseURL == "" {
		return nil, fmt.Errorf("CSC_API_KEY or CSC_BASE_URL not set in environment")
	}

	url := strings.TrimRight(baseURL, "/") + "/" + strings.TrimLeft(endpoint, "/")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-CSCAPI-KEY", apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", string(body))
	}

	return body, nil
}

// Gin Handlers

func GetCountriesHandler(c *gin.Context) {
	body, err := makeRequest("/countries")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func GetStatesHandler(c *gin.Context) {
	countryCode := c.Param("countryCode")

	endpoint := fmt.Sprintf("/countries/%s/states", countryCode)
	body, err := makeRequest(endpoint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func GetCitiesHandler(c *gin.Context) {
	countryCode := c.Param("countryCode")
	stateCode := c.Param("stateCode")

	endpoint := fmt.Sprintf("/countries/%s/states/%s/cities", countryCode, stateCode)
	body, err := makeRequest(endpoint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}
