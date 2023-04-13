package routers

import (
	"book-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.POST("/books", controllers.CreateBook)
	router.GET("/books", controllers.GetAllBook)
	router.GET("/books/:id", controllers.GetBook)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)
	return router
}
