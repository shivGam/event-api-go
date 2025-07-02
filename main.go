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
	events,err:=models.GetAllEvents()
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),"events":events})
		return
	}
	context.JSON(http.StatusOK,events)
}

func createEvent(context *gin.Context){
	var NewEvent models.Event
	err:=context.ShouldBindJSON(&NewEvent)
	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	err=NewEvent.Save()
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	context.JSON(http.StatusCreated,NewEvent)
}