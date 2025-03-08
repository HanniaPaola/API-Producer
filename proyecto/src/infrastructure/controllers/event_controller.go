package controllers

import (
    "proyecto/src/application/useCase"
    "proyecto/src/domain/entities"
    "github.com/gin-gonic/gin"
    "net/http"
)

type EventController struct {
    useCase *useCase.EventUseCase
}

// NewEventController crea una nueva instancia de EventController
func NewEventController(useCase *useCase.EventUseCase) *EventController {
    return &EventController{useCase: useCase}
}

// CreateEvent maneja la creación de un evento
func (ec *EventController) CreateEvent(c *gin.Context) {
    var event entities.Event
    if err := c.ShouldBindJSON(&event); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := ec.useCase.CreateEvent(event) // Asegúrate de capturar solo un error
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el evento"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"status": "Evento creado exitosamente"})
}

// GetAllEvents maneja la obtención de todos los eventos
func (ec *EventController) GetAllEvents(c *gin.Context) {
    events, err := ec.useCase.GetAllEvents() // Llama al caso de uso para obtener todos los eventos
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener eventos"})
        return
    }

    c.JSON(http.StatusOK, events) // Devuelve la lista de eventos
}
