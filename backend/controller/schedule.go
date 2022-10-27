package controller

import (
	"github.com/nunpromporn/sa-65-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"

	"github.com/asaskevich/govalidator"
)

// POST /schedules

func CreateSchedule(c *gin.Context) {

	var user 	entity.User								//////////*******/////////

	// var authoritie entity.Authoritie
	var doctor entity.Doctor
	var location entity.Location
	var room entity.Room
	var department entity.Department
	var schedule entity.Schedule

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 9 จะถูก bind เข้าตัวแปร schedule
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 10: ค้นหา authorities ด้วย id
	// if tx := entity.DB().Where("id = ?", schedule.AuthoritieID).First(&authoritie); tx.RowsAffected == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "authoritie not found"})
	// 	return
	// }

	if tx := entity.DB().Where("id = ?", schedule.UserID).First(&user); tx.RowsAffected == 0 {							///////*****///////
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	// 11: ค้นหา doctor ด้วย id
	if tx := entity.DB().Where("id = ?", schedule.DoctorID).First(&doctor); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "doctor not found"})
		return
	}

	// 12: ค้นหา department ด้วย id
	if tx := entity.DB().Where("id = ?", schedule.DepartmentID).First(&department); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "department not found"})
		return
	}

	// 13: ค้นหา room ด้วย id
	if tx := entity.DB().Where("id = ?", schedule.RoomID).First(&room); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room not found"})
		return
	}


	// 14: ค้นหา location ด้วย id
	if tx := entity.DB().Where("id = ?", schedule.LocationID).First(&location); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "location not found"})
		return
	}

	// 15: สร้าง Schedule
	sd := entity.Schedule{
		User: 			user,					/////โยงความสัมพันธ์กับ Entity User
		// Authoritie:  authoritie,           // โยงความสัมพันธ์กับ Entity Authoritie
		Doctor:       doctor,                // โยงความสัมพันธ์กับ Entity Doctor
		Location:     location,              // โยงความสัมพันธ์กับ Entity Location
		Room:         room,                  // โยงความสัมพันธ์กับ Entity Room
		Department:   department,            // โยงความสัมพันธ์กับ Entity Department
		ScheduleTime: schedule.ScheduleTime, // ตั้งค่าฟิลด์ Schedule
	}

	// ขั้นตอนการ validate ที่นำมาจาก unit test
	if _, err := govalidator.ValidateStruct(sd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 13: บันทึก
	if err := entity.DB().Create(&sd).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": sd})
}

// GET /schedule/:id

func GetSchedule(c *gin.Context) {
	var schedule entity.Schedule
	id := c.Param("id")
	if err := entity.DB().Preload("Patient").Preload("Department").Preload("Doctor").Preload("Location").Preload("Room").Raw("SELECT * FROM schedules WHERE id = ?", id).Find(&schedule).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": schedule})
}

// GET /schedules

func ListSchedules(c *gin.Context) {
	var schedules []entity.Schedule
	if err := entity.DB().Preload("User").Preload("Department").Preload("Doctor").Preload("Location").Preload("Room").Raw("SELECT * FROM schedules").Find(&schedules).Error; err != nil {				//////////***********/
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": schedules})
}

// DELETE /schedules/:id

func DeleteSchedule(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM schedules WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "schedule not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /schedules

func UpdateSchedule(c *gin.Context) {

	var schedule entity.Schedule

	if err := c.ShouldBindJSON(&schedule); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", schedule.ID).First(&schedule); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "schedule not found"})

		return

	}

	if err := entity.DB().Save(&schedule).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": schedule})

}