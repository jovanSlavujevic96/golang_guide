package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request event id.", "error": err.Error()})
		return
	}
	event, err := models.GetEventByID(id)
	if err != nil || event == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event. Try again later."})
		return
	}
	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err.Error()})
		return
	}

	event.ID = 1
	event.UserID = 1
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event.", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request event id.", "error": err.Error()})
		return
	}
	event, err := models.GetEventByID(id)
	if err != nil || event == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event.", "error": err.Error()})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBind(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err.Error()})
		return
	}

	updatedEvent.ID = event.ID
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated succesfully!"})
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request event id.", "error": err.Error()})
		return
	}
	event, err := models.GetEventByID(id)
	if err != nil || event == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event.", "error": err.Error()})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully!"})

}
