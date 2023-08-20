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

	router.POST("/student", controllers.CreateStudents)
	router.GET("/students", controllers.GetAllStudents)
	router.GET("/student/:ID", controllers.GetStudentsById)
	router.PUT("/student/:ID", controllers.UpdateStudents)
	router.DELETE("/student/:ID", controllers.DeleteStudents)

	return router
}
