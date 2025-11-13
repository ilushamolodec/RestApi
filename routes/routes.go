package routes

import (
	"RestApi/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	server.POST("/singup", singup)
	server.POST("/login", login)

	authenticate := server.Group("/")
	authenticate.Use(middlewares.Authenticate)
	{
		authenticate.POST("/events", createEvent)
		authenticate.PUT("/events/:id", updateEvent)
		authenticate.DELETE("/events/:id", deleteEvent)

		authenticate.POST("/events/:id/register", registerForEvents)
		authenticate.DELETE("/events/:id/register", cancelRegister)
	}

}
