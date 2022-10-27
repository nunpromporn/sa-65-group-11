package controller

import (
	"net/http"

	// "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/nunpromporn/sa-65-example/entity"
	// "golang.org/x/crypto/bcrypt"
)

// POST /authorities

// func CreateAuthoritie(c *gin.Context) {

// 	var authoritie entity.Authoritie
// 	if err := c.ShouldBindJSON(&authoritie); err != nil {

// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	bytes, err := bcrypt.GenerateFromPassword([]byte(authoritie.Password), 14)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "error hashing password"})
// 		return
// 	}
// 	authoritie.Password = string(bytes)

// 	// แทรกการ validate ไว้ช่วงนี้ของ controller
// 	if _, err := govalidator.ValidateStruct(authoritie); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	if err := entity.DB().Create(&authoritie).Error; err != nil {

// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": authoritie})
// }



// POST /authoritie

func CreateAuthoritie(c *gin.Context) {

	var authoritie entity.Authoritie
	if err := c.ShouldBindJSON(&authoritie); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&authoritie).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	c.JSON(http.StatusOK, gin.H{"data": authoritie})
}



// GET /authorities/:id

func GetAuthoritie(c *gin.Context) {

	var authoritie entity.Authoritie
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM authorities WHERE id = ?", id).Scan(&authoritie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": authoritie})

}

// GET /authorities

func ListAuthorities(c *gin.Context) {

	var authorities []entity.Authoritie

	if err := entity.DB().Raw("SELECT * FROM authorities").Scan(&authorities).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": authorities})

}

// DELETE /authorities/:id

func DeleteAuthoritie(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM authorities WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "authorities not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /authorities

func UpdateAuthoritie(c *gin.Context) {

	var authoritie entity.Authoritie
	if err := c.ShouldBindJSON(&authoritie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", authoritie.ID).First(&authoritie); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "authoritie not found"})
		return
	}
	if err := entity.DB().Save(&authoritie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": authoritie})

}
