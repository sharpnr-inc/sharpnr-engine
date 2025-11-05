package policyengine

import (
	"encoding/json"
	"fmt"
	"time"

	"sharpnr-engine/middlewares"

	"github.com/gin-gonic/gin"
)

func PolicySSEHandler(policyRepo *organizationPolicyRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		employeeID := c.Param("employeeId")

		orgID, exists := c.Get(middlewares.OrgIDContextKey)
		if !exists {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden: org not found in context"})
			return
		}

		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")

		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		var lastPolicyJSON []byte

		for {
			select {
			case <-c.Request.Context().Done():
				fmt.Println("🧹 Client disconnected:", employeeID)
				return

			case <-ticker.C:
				policy, err := policyRepo.GetEmployeePolicy(orgID.(string), employeeID)
				if err != nil {
					fmt.Println("❌ Error fetching policy:", err)
					continue
				}

				data, err := json.Marshal(policy)
				if err != nil {
					fmt.Println("❌ Error marshalling policy:", err)
					continue
				}

				// ✅ Send only if policy changed
				if string(data) != string(lastPolicyJSON) {
					fmt.Println("🔄 Sending updated policy for employee:", employeeID)
					fmt.Fprintf(c.Writer, "data: %s\n\n", data)
					c.Writer.Flush()
					lastPolicyJSON = data
				}
			}
		}
	}
}
