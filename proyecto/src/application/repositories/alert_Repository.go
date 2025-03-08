package repositories

import "proyecto/src/domain/entities"

// IEventRepository defines the methods for the event repository in the consumer.
type IEventRepository interface {
    ProcessEvent(event entities.Event) error // Procesa el evento recibido
    Save(event entities.Event) error         // Guarda el evento en la base de datos
}
