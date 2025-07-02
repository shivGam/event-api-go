package handlers

import(
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/shivGam/event-api-go/models"
)

func GetEvents(context *gin.Context) {
	events,err:=models.GetAllEvents()
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),"events":events})
		return
	}
	context.JSON(http.StatusOK,events)
}

func GetEventsById(context *gin.Context){
	eventId, err:= strconv.ParseInt(context.Param("id"),10,64)
	if err!=nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	event,err:=models.GetEventById(eventId)
	if err!=nil {
		context.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	context.JSON(http.StatusOK,event)
}

func UpdateEvent(context * gin.Context){
	eventId, err:= strconv.ParseInt(context.Param("id"),10,64)
	if err!=nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	var event models.Event
	err = context.ShouldBindJSON(&event)
	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	event.ID = eventId
	event,err = event.UpdateEvent()
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	context.JSON(http.StatusOK,event)
}

func CreateEvent(context *gin.Context){
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

func DeleteEvent(context *gin.Context){
	eventId,err:=strconv.ParseInt(context.Param("id"),10,64)
	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	}
	err = models.DeleteEvent(eventId)
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	context.JSON(http.StatusOK,gin.H{"message":"Event deleted successfully"})
}