package organization

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler to get all organisations
func GetOrganisations(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "List of organisations",
        "data":    []string{"Org A", "Org B"},
    })
}

// Handler to get organisation by ID
func GetOrganisationByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "message": "Organisation found",
        "id":      id,
        "name":    "Org " + idParam,
    })
}
