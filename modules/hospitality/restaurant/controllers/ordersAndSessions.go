package restaurantcontrollers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *restaurantController) CreateSession(ctx *gin.Context) {

	// ✅ Get OrgUUID from context (set by middleware)
	// orgUUID, exists := ctx.Get("OrgUUID")
	// if !exists {
	// 	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Organization UUID missing"})
	// 	return
	// }
	// orgID, ok := orgUUID.(uuid.UUID)
	// if !ok {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid organization UUID format"})
	// 	return
	// }

	// fmt.Println(orgID)

	// var req hospitalitydtos.Hospitality_Restaurant_CreateSession_REQUEST

	// if err := ctx.ShouldBindJSON(&req); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
	// 	return
	// }

	// if err := c.resSvc.CreateTableBookingSession(ctx, orgID, req); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
	// 	return
	// }

	// ctx.JSON(http.StatusOK, gin.H{"success": "Booking session Created"})

}

//---------------------------------------------------------------------------------------------------------------//

func (c *restaurantController) AddOrdersInSession(ctx *gin.Context) {

	// orgUUID, exists := ctx.Get("OrgUUID")
	// if !exists {
	// 	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Organization UUID missing"})
	// 	return
	// }
	// orgID, ok := orgUUID.(uuid.UUID)
	// if !ok {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid organization UUID format"})
	// 	return
	// }

	// fmt.Println(orgID)

	// sessionIdParam := ctx.Param("id")
	// if sessionIdParam == "" {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "Table ID not provided",
	// 	})
	// 	return
	// }

	// sessionID, err := uuid.Parse(sessionIdParam)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "Invalid session ID format",
	// 	})
	// 	return
	// }

	// var req hospitalitydtos.Hospitality_Restaurant_AddOrderItemINTheSession_Request
	// if err := ctx.ShouldBindJSON(&req); err != nil {
	// 	res := dto.Response[any]{
	// 		StatusCode: http.StatusOK,
	// 		Message:    "Error Creating the Order",
	// 		Error: dto.ErrorResponse{
	// 			Message: err.Error(),
	// 			IsError: true,
	// 		},
	// 		Data:      nil,
	// 		Timestamp: time.Now(),
	// 	}
	// 	ctx.JSON(http.StatusOK, res)
	// 	return
	// }

	// // it need sessionId and ReqBody
	// if err := c.resSvc.AddItemInBookingSession(ctx, orgID, sessionID, req); err != nil {
	// 	response := dto.Response[any]{
	// 		StatusCode: http.StatusOK,
	// 		Message:    "There was an Error",
	// 		Error: dto.ErrorResponse{
	// 			Message: err.Error(),
	// 			IsError: true,
	// 		},
	// 		Data:      nil,
	// 		Timestamp: time.Now(),
	// 	}
	// 	ctx.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	// response := dto.Response[any]{
	// 	StatusCode: http.StatusOK,
	// 	Message:    "Order Sent to Kitchen successfully",
	// 	Error: dto.ErrorResponse{
	// 		Message: "",
	// 		IsError: false,
	// 	},
	// 	Data:      nil,
	// 	Timestamp: time.Now(),
	// }
	// ctx.JSON(http.StatusCreated, response)
}

func (c *restaurantController) CreateInvoiceForSession(ctx *gin.Context) {}

func (c *restaurantController) GetBookingSessionById(ctx *gin.Context) {

	// Here we have the OrgId and we can take the OrgId and Fetch the Restaurant
	orgUUID, exists := ctx.Get("OrgUUID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Organization UUID missing"})
		return
	}
	orgID, ok := orgUUID.(uuid.UUID)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid organization UUID format"})
		return
	}

	// --------------------------------
	// 2️⃣ Read tableId from URL param
	// --------------------------------
	tableIdParam := ctx.Param("id")
	if tableIdParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Table ID not provided",
		})
		return
	}

	tableID, err := uuid.Parse(tableIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid table ID format",
		})
		return
	}

	session, err := c.resSvc.GetResTableBookingSessionByTableId(ctx, orgID, tableID)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": "something bad happen"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"res": session})

}

func (c *restaurantController) GetOrderStatus(ctx *gin.Context) {
	orgUUID, exists := ctx.Get("OrgUUID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Organization UUID missing"})
		return
	}
	orgID, ok := orgUUID.(uuid.UUID)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid organization UUID format"})
		return
	}

	kitchenTicketIdParams := ctx.Param("id")
	if kitchenTicketIdParams == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id not provided"})
		return
	}

	kitchenTicketId, err := uuid.Parse(kitchenTicketIdParams)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "KitchenId is in wrong Format"})
		return
	}

	// res, err := c.resSvc.GetOrderStatus(ctx, orgID, kitchenTicketId)
	// if err != nil {

	// }

	fmt.Println(kitchenTicketId)
	fmt.Println(orgID)

}
