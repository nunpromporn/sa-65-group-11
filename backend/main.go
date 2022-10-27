package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nunpromporn/sa-65-example/controller"
	"github.com/nunpromporn/sa-65-example/entity"
	"github.com/nunpromporn/sa-65-example/middlewares"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{

			// User Routes
			protected.GET("/users", controller.ListUsers)
			protected.GET("/user/:id", controller.GetUser)
			// protected.POST("/users", controller.CreateUser)
			// protected.PATCH("/users", controller.UpdateUser)
			// protected.DELETE("/users/:id", controller.DeleteUser)

			
			// Authoritie Routes
			protected.GET("/authorities", controller.ListAuthorities)
			protected.GET("/authoritie/:id", controller.GetAuthoritie)
			protected.POST("/authorities", controller.CreateAuthoritie)				///******
			protected.PATCH("/authorities", controller.UpdateAuthoritie)
			protected.DELETE("/authorities/:id", controller.DeleteAuthoritie)

			// Department Routes
			protected.GET("/departments", controller.ListDepartments)
			protected.GET("/department/:id", controller.GetDepartment)
			protected.POST("/departments", controller.CreateDepartment)
			protected.PATCH("/departments", controller.UpdateDepartment)
			protected.DELETE("/departments/:id", controller.DeleteDepartment)

			// doctor Routes
			protected.GET("/doctors", controller.ListDoctors)
			protected.GET("/doctor/:id", controller.GetDoctor)
			protected.POST("/doctors", controller.CreateDoctor)
			protected.PATCH("/doctors", controller.UpdateDoctor)
			protected.DELETE("/doctors/:id", controller.DeleteDoctor)

			// location Routes
			protected.GET("/locations", controller.ListLocations)
			protected.GET("/location/:id", controller.GetLocation)
			protected.POST("/locations", controller.CreateLocation)
			protected.PATCH("/locations", controller.UpdateLocation)
			protected.DELETE("/locations/:id", controller.DeleteLocation)

			// room Routes
			protected.GET("/rooms", controller.ListRooms)
			protected.GET("/room/:id", controller.GetRoom)
			protected.POST("/rooms", controller.CreateRoom)
			protected.PATCH("/rooms", controller.UpdateRoom)
			protected.DELETE("/rooms/:id", controller.DeleteRoom)

			// schedule Routes
			protected.GET("/schedules", controller.ListSchedules)
			protected.GET("/schedule/:id", controller.GetSchedule)
			protected.POST("/schedules", controller.CreateSchedule)
			protected.PATCH("/schedules", controller.UpdateSchedule)
			protected.DELETE("/schedules/:id", controller.DeleteSchedule)

		}

	}


	r.POST("/users", controller.CreateUser)										////*****



	// authoritie Routes
	// r.POST("/authorities", controller.CreateAuthoritie)

	// Authentication Routes
	r.POST("/login", controller.Login)

	// Run the server
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}

}
