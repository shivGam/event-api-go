package handlers

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/shivGam/event-api-go/models"
)

func CreateUser(context *gin.Context){
	var user models.User
	err:=context.ShouldBindJSON(&user)
	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	err = user.Save()
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	context.JSON(http.StatusCreated,gin.H{"message":"User created successfully"})
}

func LoginUser(context *gin.Context){
	var user models.User
	err:=context.ShouldBindJSON(&user)
	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	err=user.ValidateCredentials()
	if err!=nil{
		context.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid credentials"})
		return
	}
	context.JSON(http.StatusOK,gin.H{"message":"Login successful"})
}
