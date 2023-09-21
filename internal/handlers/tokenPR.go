package handlers

import (
	"demerzel-events/internal/db"
	"demerzel-events/internal/models"
	"demerzel-events/pkg/jwt"
	"demerzel-events/pkg/response"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetTokenTest(c *gin.Context) {
	var user models.User
	db.DB.First(&user)
	token, _ := jwt.CreateToken(map[string]any{"id": user.Id}, os.Getenv("JWT_SECRET"), 3)
	response.Success(c, http.StatusOK, "User retrieved successfully", map[string]any{"token": token})
}

func GetUserGroups(c *gin.Context) {
	rawUser, exists := c.Get("user")

	if !exists {
		response.Error(c, http.StatusConflict, "error: unable to retrieve user from context")
		return
	}
	user, ok := rawUser.(*models.User)
	if !ok {
		response.Error(c, http.StatusConflict, "error: invalid user type in context")
		return
	}

	var groups []models.Group
	result := db.DB.Joins("join user_groups on groups.id=user_groups.group_id").Where("user_groups.user_id=?", user.Id).Find(&groups)
	// result := db.DB.Where(&models.UserGroup{
	// 	UserID: user.Id,
	// }).Find(&groups)

	if result.Error != nil {
		response.Error(c, http.StatusConflict, "error: something went wrong")
		return // Return the actual error for other errors
	}

	response.Success(c, http.StatusOK, "User retrieved successfully", map[string]any{"groups": groups})
}
