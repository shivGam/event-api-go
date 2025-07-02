package main

import (
	"net/http"
	"github.com/shivGam/event-api-go/db"
	"github.com/shivGam/event-api-go/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server:= gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events",createEvent)

	server.Run(":8080")

}	

func getEvents(context *gin.Context) {
	events:=models.GetAllEvents()

	context.JSON(http.StatusOK,events)
}

func createEvent(context *gin.Context){
	var NewEvent models.Event
	err:=context.ShouldBindJSON(&NewEvent)
	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	NewEvent.UserID=1
	NewEvent.ID=1
	NewEvent.Save()
	context.JSON(http.StatusCreated,NewEvent)
}