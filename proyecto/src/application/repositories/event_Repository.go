package repositories

import (
    "proyecto/src/domain/entities"
    "database/sql"
)

// IEventRepository define los métodos que debe implementar el repositorio de eventos.
type IEventRepository interface {
    Create(event entities.Event) (int, error)   // Método para crear un evento
    FindAll() ([]entities.Event, error)          // Método para obtener todos los eventos
}

// EventRepository es la implementación concreta de IEventRepository
type EventRepository struct {
    db *sql.DB // Aquí puedes incluir la conexión a la base de datos
}

// NewEventRepository crea una nueva instancia de EventRepository
func NewEventRepository(db *sql.DB) *EventRepository {
    return &EventRepository{db: db}
}

// Create implementa el método para crear un evento
func (er *EventRepository) Create(event entities.Event) (int, error) {
    // Implementa la lógica para insertar el evento en la base de datos
    return 1, nil // Cambia esto por la lógica real
}

// FindAll implementa el método para obtener todos los eventos
func (er *EventRepository) FindAll() ([]entities.Event, error) {
    // Implementa la lógica para obtener todos los eventos de la base de datos
    return []entities.Event{}, nil // Cambia esto por la lógica real
}