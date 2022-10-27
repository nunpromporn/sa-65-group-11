package controller

import (
	"github.com/nunpromporn/sa-65-example/entity"
	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /location

func CreateLocation(c *gin.Context) {

	var location entity.Location
	if err := c.ShouldBindJSON(&location); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&location).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	c.JSON(http.StatusOK, gin.H{"data": location})
}

// GET /location/:id

func GetLocation(c *gin.Context) {

	var location entity.Location

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM locations WHERE id = ?", id).Scan(&location).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": location})

}

// GET /locations
//9.ดึงข้อมูลทั้งหมด
func ListLocations(c *gin.Context) {

	var locations []entity.Location

	if err := entity.DB().Raw("SELECT * FROM locations").Scan(&locations).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": locations})

}

// DELETE /locations/:id

func DeleteLocation(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM locations WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "location not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /locations

func UpdateLocation(c *gin.Context) {

	var location entity.Location

	if err := c.ShouldBindJSON(&location); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", location.ID).First(&location); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "location not found"})

		return

	}

	if err := entity.DB().Save(&location).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": location})

}
