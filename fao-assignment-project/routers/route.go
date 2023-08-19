package routers

import (
	"assignment-project/controllers"
	"assignment-project/database"

	"github.com/gin-gonic/gin"
)

func init() {
	database.StartDB()
}

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/Student", controllers.CreateStudents)
	router.GET("/Students/", controllers.GetAllStudents)
	router.GET("/Student/:ID", controllers.GetStudentsById)
	router.PUT("/Student/:ID", controllers.UpdateStudents)
	router.DELETE("/Student/:ID", controllers.DeleteStudents)

	return router
}
