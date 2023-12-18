package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wtran29/event-booking/models"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func getEvent(ctx *gin.Context) {
	eid, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	event, err := models.GetEventByID(eid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}
	ctx.JSON(http.StatusOK, event)
}
func createEvent(ctx *gin.Context) {

	var event models.Event
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	userId := ctx.GetInt64("userId")
	event.UserID = userId
	if err := event.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func updateEvent(ctx *gin.Context) {
	eid, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	userId := ctx.GetInt64("userId")
	event, err := models.GetEventByID(eid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	if event.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to update event!"})
		return
	}

	var updatedEvent models.Event
	err = ctx.ShouldBindJSON(&updatedEvent)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request data event."})
		return
	}
	updatedEvent.ID = eid
	err = updatedEvent.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!"})
}

func deleteEvent(ctx *gin.Context) {
	eid, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	userId := ctx.GetInt64("userId")
	event, err := models.GetEventByID(eid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}
	if event.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to delete event!"})
		return
	}

	err = event.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully!"})
}
