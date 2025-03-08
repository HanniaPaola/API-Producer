package db

import (
    "database/sql"
    "proyecto/src/domain/entities"
)

// EventRepository implements IEventRepository for MySQL
type EventRepository struct {
    db *sql.DB // Conexión a la base de datos
}

// NewEventRepository crea una nueva instancia de EventRepository
func NewEventRepository(db *sql.DB) *EventRepository {
    return &EventRepository{db: db}
}

// Create guarda un evento en la base de datos
func (r *EventRepository) Create(event entities.Event) (int, error) {
    // Aquí se implementaría la lógica para guardar el evento en la base de datos
    query := "INSERT INTO events (type, created_at, unit) VALUES (?, ?, ?)"
    result, err := r.db.Exec(query, event.Type, event.CreatedAt, event.Unit)
    if err != nil {
        return 0, err
    }

    // Obtén el ID del evento creado
    eventID, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    return int(eventID), nil
}

// GetAll recupera todos los eventos de la base de datos
func (r *EventRepository) GetAll() ([]entities.Event, error) {
    // Aquí se implementaría la lógica para recuperar eventos de la base de datos
    rows, err := r.db.Query("SELECT event_id, type, created_at, unit FROM events")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var events []entities.Event
    for rows.Next() {
        var event entities.Event
        if err := rows.Scan(&event.EventID, &event.Type, &event.CreatedAt, &event.Unit); err != nil {
            return nil, err
        }
        events = append(events, event)
    }

    return events, nil
}
