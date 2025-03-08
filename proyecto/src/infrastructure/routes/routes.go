// src/infrastructure/routes/routes.go

package routes

import (
    "github.com/gin-gonic/gin"
    "proyecto/src/infrastructure/controllers"
    "proyecto/src/application/useCase"
)

// RegisterRoutes registra las rutas en el router de Gin
func RegisterRoutes(r *gin.Engine, eventUseCase *useCase.EventUseCase) {
    eventController := controllers.NewEventController(eventUseCase)

    r.POST("/events", eventController.CreateEvent)
    r.GET("/events", eventController.GetAllEvents) // Asegúrate de que este método esté implementado en el controlador
}
