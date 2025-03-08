package useCase

import (
    "proyecto/src/domain/entities"
    "proyecto/src/application/services"
)

type EventUseCase struct {
    eventService services.IEventService
}

// NewEventUseCase crea una nueva instancia de EventUseCase
func NewEventUseCase(eventService services.IEventService) *EventUseCase {
    return &EventUseCase{eventService: eventService}
}

// CreateEvent llama al servicio para crear un evento
func (euc *EventUseCase) CreateEvent(event entities.Event) error {
    return euc.eventService.CreateEvent(event) // Solo devuelve un error
}

// GetAllEvents obtiene todos los eventos
func (euc *EventUseCase) GetAllEvents() ([]entities.Event, error) {
    return euc.eventService.GetAllEvents() // Llama al método del servicio
}
