package controller

import (
	"github.com/nunpromporn/sa-65-example/entity"
	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /departments

func CreateDepartment(c *gin.Context) {

	var department entity.Department
	if err := c.ShouldBindJSON(&department); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&department).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	c.JSON(http.StatusOK, gin.H{"data": department})
}

// GET /department/:id
// เพื่อดึงข้อมูล department ออกมาตาม primary key ที่กำหนด ผ่าน func DB.Raw(...)
func GetDepartment(c *gin.Context) {

	var department entity.Department
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM departments WHERE id = ?", id).Scan(&department).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": department})

}

// GET /departments
// เป็นการ list รายการของ Departments ออกมา
func ListDepartments(c *gin.Context) {

	var departments []entity.Department

	if err := entity.DB().Raw("SELECT * FROM departments").Scan(&departments).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}
	c.JSON(http.StatusOK, gin.H{"data": departments})

}

// DELETE /departments/:id
// เป็น function สำหรับลบ department ด้วย ID
func DeleteDepartment(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM departments WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "department not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /departments

func UpdateDepartment(c *gin.Context) {

	var department entity.Department
	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", department.ID).First(&department); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "department not found"})
		return
	}
	if err := entity.DB().Save(&department).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": department})
}