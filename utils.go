package final_project

import (
	"fmt"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/middleware"
	"github.com/gin-gonic/gin"
)

func IsAdmin(c *gin.Context) bool {
	userRole := middleware.GetUserDetailsFromToken(c)["userRole"]
	fmt.Println(userRole)
	if userRole != "Admin" {
		return false
	}
	return true
}

func IsAuthorizedOrReadOnly(c *gin.Context) bool {
	loggedUser := middleware.GetUserDetailsFromToken(c)
	if loggedUser == nil {
		return false
	}

	return true
}
