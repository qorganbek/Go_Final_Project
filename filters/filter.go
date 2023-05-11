package filters

import (
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/initializers"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOfAdvertisementsByPrice(c *gin.Context) {
	var advertisements []models.Advertisement

	initializers.DB.Find(&advertisements)
	if len(advertisements) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't retrieve data!"})
		return
	}

	min := c.Param("min")
	max := c.Param("max")
	initializers.DB.Where("? <= price AND price <= ?", min, max).Find(&advertisements)
	// SELECT * FROM users WHERE Price_Min <= age <= Price_Max;

	c.JSON(http.StatusOK, gin.H{"data": advertisements})
}

func ListOfAdvertisementsByYear(c *gin.Context) {
	var advertisements []models.Advertisement

	initializers.DB.Find(&advertisements)
	if len(advertisements) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't retrieve data!"})
		return
	}

	min := c.Param("min")
	max := c.Param("max")
	initializers.DB.Where("? <= year AND year <= ?", min, max).Find(&advertisements)
	// SELECT * FROM users WHERE Age_Min <= age <= Age_Max;

	c.JSON(http.StatusOK, gin.H{"data": advertisements})
}
