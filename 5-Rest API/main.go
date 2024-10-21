package main

import (
	"net/http"

	"example.com/rest/db"
	"example.com/rest/models"
	"github.com/gin-gonic/gin"
)

func main() {
    // Veritabanı bağlantısını sağlar.
	db.Connect()
	// Yeni bir gin sunucusu oluşturur
	server := gin.Default() 

	// Get isteği ile events endpoint'ine erişildiğinde get events fonskiyonunu çağırır.
	server.GET("/events", getEvents) 
	server.POST("/events", createEvent)

	// Sunucuyu 8080 portunda çalıştır
	server.Run(":8080") 

}

func getEvents(context *gin.Context){

	// Tüm etkinlikleri alır
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later"})
		return
	}

	// JSON formatında yanıt döndürür.
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse"})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": event})
}