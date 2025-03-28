package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Lấy danh sách người dùng (chỉ dành cho admin)
func GetUsers(c *gin.Context) {
	// var users []database.User
	// if err := database.DB.Find(&users).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lấy danh sách người dùng"})
	// 	return
	// }
	var users []string
	users = append(users, "user1", "user2", "user3")
	c.JSON(http.StatusOK, users)
}
