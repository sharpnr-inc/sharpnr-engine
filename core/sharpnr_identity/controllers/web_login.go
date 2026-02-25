package sharpnr_identity_auth_controllers

import (
	"net/http"
	"sharpnr-engine/cookieManager"
	sharpnr_identity_auth_dtos "sharpnr-engine/core/sharpnr_identity/dtos"

	"github.com/gin-gonic/gin"
)

func (s *sharpnrIdentityAuthControllers) WebLogin(ctx *gin.Context) {
	var req sharpnr_identity_auth_dtos.LoginRequest_dto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	res, err := s.IdSvc.LoginUser(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"Error": err.Error(),
		})
		return
	}

	token, err := s.jwtMgr.IssueWebAuthCookieToken(res.EmployeeID, res.OrgId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"Error": err.Error(),
		})
		return
	}
	s.cookieMgr.Set(ctx, cookieManager.CSRF, token)
	s.cookieMgr.Set(ctx, cookieManager.Auth, token)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login Successfully",
		"data":    res,
	})

}
