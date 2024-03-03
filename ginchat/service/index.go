package service

import (
	"ginchat/models"
	"github.com/gin-gonic/gin"
)

// GetIndex
// @Tags 首页
// @Success 200 {string} welcome
// @Router /index [get]
func GetIndex(ctx *gin.Context) {
	ctx.JSONP(200, gin.H{
		"message": models.GetUserList(),
	})
}
