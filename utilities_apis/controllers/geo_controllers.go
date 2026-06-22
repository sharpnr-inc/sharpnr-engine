package utilitiesApi_controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// These handlers replicate the countrystatecity.in REST endpoints so the
// frontend can repoint its base URL at this service unchanged. Responses are
// raw JSON arrays/objects (no envelope), matching the upstream API.

// GetAllCountries -> GET /countries
func (c *utilitiesApiController) GetAllCountries(ctx *gin.Context) {
	countries, err := c.utilitiesApiService.ListCountries(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch countries"})
		return
	}
	ctx.JSON(http.StatusOK, countries)
}

// GetCountry -> GET /countries/:ciso
func (c *utilitiesApiController) GetCountry(ctx *gin.Context) {
	country, err := c.utilitiesApiService.GetCountry(ctx.Request.Context(), ctx.Param("ciso"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Country not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch country"})
		return
	}
	ctx.JSON(http.StatusOK, country)
}

// GetStatesByCountry -> GET /countries/:ciso/states
func (c *utilitiesApiController) GetStatesByCountry(ctx *gin.Context) {
	states, err := c.utilitiesApiService.ListStates(ctx.Request.Context(), ctx.Param("ciso"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch states"})
		return
	}
	ctx.JSON(http.StatusOK, states)
}

// GetState -> GET /countries/:ciso/states/:siso
func (c *utilitiesApiController) GetState(ctx *gin.Context) {
	state, err := c.utilitiesApiService.GetState(ctx.Request.Context(), ctx.Param("ciso"), ctx.Param("siso"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "State not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch state"})
		return
	}
	ctx.JSON(http.StatusOK, state)
}

// GetCitiesByState -> GET /countries/:ciso/states/:siso/cities
func (c *utilitiesApiController) GetCitiesByState(ctx *gin.Context) {
	cities, err := c.utilitiesApiService.ListCities(ctx.Request.Context(), ctx.Param("ciso"), ctx.Param("siso"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch cities"})
		return
	}
	ctx.JSON(http.StatusOK, cities)
}
