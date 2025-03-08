// src/application/services/event_service.go (CONSUMIDOR)
package services

import (
    "log"
    "proyecto/src/domain/entities"
    "proyecto/src/application/repositories"
)

type EventService struct {
    eventRepository repositories.IEventRepository // Repositorio para guardar eventos
}

// NewEventService crea un nuevo servicio de eventos en el consumidor
func NewEventService(eventRepository repositories.IEventRepository) *EventService {
    return &EventService{eventRepository: eventRepository}
}

// ProcessEvent procesa un evento recibido desde RabbitMQ
func (es *EventService) ProcessEvent(event entities.Event) error {
    log.Println("📥 Procesando evento:", event)
    return es.eventRepository.Save(event) // Guarda el evento en la base de datos
}
