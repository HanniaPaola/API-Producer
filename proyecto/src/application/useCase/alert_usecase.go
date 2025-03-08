// src/application/useCase/event_usecase.go (CONSUMIDOR)
package useCase

import (
    "proyecto/src/domain/entities"
    "proyecto/src/application/services"
)

type EventUseCase struct {
    eventService *services.EventService
}

// NewEventUseCase crea un nuevo caso de uso de eventos en el consumidor
func NewEventUseCase(eventService *services.EventService) *EventUseCase {
    return &EventUseCase{eventService: eventService}
}

// ProcessEvent procesa un evento recibido desde RabbitMQ y lo guarda en MySQL
func (euc *EventUseCase) ProcessEvent(event entities.Event) error {
    return euc.eventService.ProcessEvent(event) // Llama al método del servicio
}

// GetAllEvents obtiene todos los eventos almacenados en MySQL
func (euc *EventUseCase) GetAllEvents() ([]entities.Event, error) {
    return euc.eventService.GetAllEvents() // Solo si quieres listar eventos guardados
}
