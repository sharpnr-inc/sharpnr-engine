package restaurantcontrollers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (c *restaurantController) CreateDineSession(ctx *gin.Context) {

	orgID, empID, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	// 2. Resolve Restaurant ID from params
	restaurantID, ok := c.resolveRestaurantID(ctx)
	if !ok {
		return // Response already handled by helper
	}

	fmt.Println(restaurantID)
	fmt.Println(orgID)
	fmt.Println(empID)

	

}

/* ----------------------------------- <*> ---------------------------------- */

