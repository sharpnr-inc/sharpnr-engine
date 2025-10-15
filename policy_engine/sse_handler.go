package policyengine

import (
	"encoding/json"
	"fmt"
	"time"

	"sharpnr-engine/shared/middlewares"

	"github.com/gin-gonic/gin"
)

func PolicySSEHandler(policyRepo *organisationPolicyRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		employeeID := c.Param("employeeId")

		// ✅ Extract orgID from context (set by middleware)
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

		for {
			select {
			case <-c.Request.Context().Done():
				fmt.Println("Client disconnected:", employeeID)
				return

			case <-ticker.C:
				// ✅ Pass orgID + employeeID to repo
				policy, err := policyRepo.GetEmployeePolicy(orgID.(string), employeeID)
				if err != nil {
					fmt.Println("Error fetching policy:", err)
					continue
				}

				data, err := json.Marshal(policy)
				if err != nil {
					fmt.Println("Error marshalling policy:", err)
					continue
				}

				fmt.Fprintf(c.Writer, "data: %s\n\n", data)
				c.Writer.Flush()
			}
		}
	}
}
