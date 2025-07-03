package main

import (
	"github.com/shivGam/event-api-go/db"
	"github.com/shivGam/event-api-go/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server:= gin.Default()

	server.GET("/events", handlers.GetEvents)
	server.GET("events/:id",handlers.GetEventsById)
	server.POST("/events",handlers.CreateEvent)
	server.PUT("/events/:id",handlers.UpdateEvent)
	server.DELETE("/events/:id",handlers.DeleteEvent)
	
	server.POST("/signup",handlers.CreateUser)
	server.POST("/login",handlers.LoginUser)
	server.Run(":8080")
}