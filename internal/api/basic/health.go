package basic

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthApi struct{}

// @Tags basic
// @Summary 心跳检测
// @Description 心跳检测aaa
// @Success 200 {string} health
// @Router /health [get]
func (h *HealthApi) Health(c *gin.Context) {
	c.JSON(http.StatusOK, "health")
}
