package hotelcontrollers

import (
	"github.com/gin-gonic/gin"
)

func (c *HotelController) AddPayments(ctx *gin.Context) {
	// 1️⃣ Org context
	// orgUUIDVal, exists := ctx.Get("OrgUUID")
	// if !exists {
	// 	utils.ErrorResponse(
	// 		ctx,
	// 		http.StatusUnauthorized,
	// 		"Organization UUID missing",
	// 		nil,
	// 	)
	// 	return
	// }
	// orgID, ok := orgUUIDVal.(uuid.UUID)
	// if !ok {
	// 	utils.ErrorResponse(
	// 		ctx,
	// 		http.StatusBadRequest,
	// 		"Invalid organization UUID",
	// 		nil,
	// 	)
	// 	return
	// }

	// var req hoteldtos.AddPaymentsForBookingRequest_dto
	// if err := ctx.ShouldBindJSON(&req); err != nil {
	// 	utils.ErrorResponse(
	// 		ctx,
	// 		http.StatusBadRequest,
	// 		"Invalid request payload",
	// 		err,
	// 	)
	// 	return
	// }

	// // 3️⃣ Call service
	// if err := c.hotelService.AddPaymentForBooking(ctx, orgID, req); err != nil {
	// 	utils.ErrorResponse(
	// 		ctx,
	// 		http.StatusBadRequest,
	// 		"Failed to add payment",
	// 		err,
	// 	)
	// 	return
	// }

	// // 4️⃣ Success
	// var data = "SUCCESS"
	// utils.SuccessResponse(
	// 	ctx,
	// 	http.StatusCreated,
	// 	"Payment added successfully",
	// 	data,
	// )

}
