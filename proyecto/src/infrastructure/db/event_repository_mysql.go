package db

import (
    "database/sql"
    "log"
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

// Create inserta un nuevo evento en la base de datos y devuelve su ID
func (r *EventRepository) Create(event entities.Event) (int, error) {
    query := "INSERT INTO events (type, created_at, unit) VALUES (?, ?, ?)"
    
    // Ejecutar la consulta de inserción
    result, err := r.db.Exec(query, event.Type, event.CreatedAt, event.Unit)
    if err != nil {
        log.Printf("Error al insertar el evento: %s", err)
        return 0, err // Retorna 0 y el error si la inserción falla
    }
    
    // Obtener el ID del evento insertado
    eventID, err := result.LastInsertId()
    if err != nil {
        log.Printf("Error al obtener el ID del evento: %s", err)
        return 0, err // Retorna 0 y el error si no se puede obtener el ID
    }
    
    return int(eventID), nil // Retorna el ID del evento creado y nil si todo fue bien
}

// GetAll obtiene todos los eventos de la base de datos
func (r *EventRepository) GetAll() ([]entities.Event, error) {
    query := "SELECT id, type, created_at, unit FROM events"
    rows, err := r.db.Query(query)
    if err != nil {
        log.Printf("Error al obtener eventos: %s", err)
        return nil, err
    }
    defer rows.Close()

    var events []entities.Event
    for rows.Next() {
        var event entities.Event
        if err := rows.Scan(&event.ID, &event.Type, &event.CreatedAt, &event.Unit); err != nil {
            log.Printf("Error al escanear el evento: %s", err)
            continue
        }
        events = append(events, event)
    }

    return events, nil
}
