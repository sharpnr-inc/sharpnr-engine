package hotelcontrollers

import (
	"fmt"
	"net/http"
	"sharpnr-engine/sere"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func (s *HotelController) HotelWSHandler(hub *sere.Engine) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		orgUUID, exist := ctx.Get("OrgUUID")
		if !exist {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Organization UUID missing"})
			return
		}

		orgID, ok := orgUUID.(uuid.UUID)
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid organization UUID format"})
			return
		}
		fmt.Println(orgID)
		upgrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		}

		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer conn.Close()

		// subID, _ := hub.Subscribe("ws.hotel."+orgID.String(), sere.And(sere.MatchOrg(orgID), sere.Or()))

	}
}
