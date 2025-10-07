package coreutilitiesapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// replace with your own API key from data.gov.in
const apiKey = "579b464db66ec23bdd0000019cbce1e7cc0a473d74bceded45f67343"

// API endpoint from Data.gov.in (Town Directory dataset)
const baseURL = "https://api.data.gov.in/resource/97878f8a-97c3-4b52-b4bd-a40c530d0fe6"

type TownRecord struct {
	StateName string `json:"state_name"`
	TownName  string `json:"town_name"`
	District  string `json:"district_name"`
}

type ApiResponse struct {
	Records []TownRecord `json:"records"`
}

// GET /utils/get-allstate
func GetAllStateInIndia(c *gin.Context) {
	url := fmt.Sprintf("%s?api-key=%s&format=json&limit=10000", baseURL, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}
	defer resp.Body.Close()

	var data ApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}

	stateSet := make(map[string]bool)
	for _, record := range data.Records {
		stateSet[record.StateName] = true
	}

	states := make([]string, 0, len(stateSet))
	for state := range stateSet {
		states = append(states, state)
	}

	c.JSON(http.StatusOK, gin.H{"states": states})
}

// GET /utils/get-cities/:state
func GetCitiesByState(c *gin.Context) {
	state := c.Param("state")
	url := fmt.Sprintf("%s?api-key=%s&format=json&limit=10000", baseURL, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}
	defer resp.Body.Close()

	var data ApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}

	var cities []string
	for _, record := range data.Records {
		if strings.EqualFold(record.StateName, state) {
			cities = append(cities, record.TownName)
		}
	}

	if len(cities) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No cities found for this state"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"state": state, "cities": cities})
}
