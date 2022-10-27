package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nunpromporn/sa-65-example/entity"
)

// POST /doctors

func CreateDoctor(c *gin.Context) {

	var doctor entity.Doctor
	if err := c.ShouldBindJSON(&doctor); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&doctor).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	c.JSON(http.StatusOK, gin.H{"data": doctor})
}

// GET /doctor/:id
// เพื่อดึงข้อมูล doctor ออกมาตาม primary key ที่กำหนด ผ่าน func DB.Raw(...)
func GetDoctor(c *gin.Context) {

	var doctor entity.Doctor
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM doctors WHERE id = ?", id).Scan(&doctor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": doctor})

}

// GET /doctors
// เป็นการ list รายการของ Doctors ออกมา
func ListDoctors(c *gin.Context) {

	var doctors []entity.Doctor

	if err := entity.DB().Raw("SELECT * FROM doctors").Scan(&doctors).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}
	c.JSON(http.StatusOK, gin.H{"data": doctors})

}

// DELETE /doctors/:id
// เป็น function สำหรับลบ doctor ด้วย ID
func DeleteDoctor(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM doctors WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "doctor not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /doctors

func UpdateDoctor(c *gin.Context) {

	var doctor entity.Doctor
	if err := c.ShouldBindJSON(&doctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", doctor.ID).First(&doctor); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "doctor not found"})
		return
	}
	if err := entity.DB().Save(&doctor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": doctor})
}