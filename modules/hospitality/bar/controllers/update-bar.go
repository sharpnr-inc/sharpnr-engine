package barcontrollers

import (
	"fmt"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

	"github.com/gin-gonic/gin"
)

// Update godoc
// @Summary Update Bar Details
// @Description
// @Tags Hospitality - Bar
func (c *barController) UpdateBarDetails(ctx *gin.Context) {
	/*
	   TODO Complete this Update Bar details controller

	   *1 Create the DTO for REQ and RES
	   *2 Create its Service to Update Bar Details
	*/

	// Step 1 --> Get the Request Data//
	var req hospitalitydtos.BAR_UpdateBarDetails_Request
	fmt.Println(req)

	// Step 2 --> find the Bar Based on the OrgId in the Params

	// Step 3 -->

}
