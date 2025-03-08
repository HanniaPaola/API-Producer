// src/application/services/event_service.go

package services

import (
    "proyecto/src/domain/entities"
    "proyecto/src/application/repositories"
)

type IEventService interface {
    CreateEvent(event entities.Event) error
    GetAllEvents() ([]entities.Event, error)
}

type EventService struct {
    eventRepo repositories.IEventRepository
}

// NewEventService crea una nueva instancia de EventService
func NewEventService(eventRepo repositories.IEventRepository) *EventService {
    return &EventService{eventRepo: eventRepo}
}

// CreateEvent envía el evento (no realiza almacenamiento)
func (s *EventService) CreateEvent(event entities.Event) error {
    // No se almacena el evento, por lo que solo puedes devolver nil
    return nil
}

// GetAllEvents devuelve una lista de eventos (puede devolver un slice vacío si no hay eventos)
func (s *EventService) GetAllEvents() ([]entities.Event, error) {
    // Retorna un slice vacío como no se almacenan eventos
    return []entities.Event{}, nil
}