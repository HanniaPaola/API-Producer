// src/infrastructure/controllers/event_consumer.go

package controllers

import (
    "log"
    "github.com/streadway/amqp"
    "proyecto/src/domain/entities"
)

type EventConsumer struct {
    rabbitMQ *adapters.RabbitMQ
}

// NewEventConsumer crea un nuevo EventConsumer
func NewEventConsumer(rabbitMQ *adapters.RabbitMQ) *EventConsumer {
    return &EventConsumer{
        rabbitMQ: rabbitMQ,
    }
}

// StartListening inicia la escucha de eventos desde RabbitMQ
func (ec *EventConsumer) StartListening() {
    messages, err := ec.rabbitMQ.Channel.Consume(
        ec.rabbitMQ.Queue,
        "",    // consumer
        false, // auto-ack
        false, // exclusive
        false, // no-local
        false, // no-wait
        nil,   // arguments
    )
    if err != nil {
        log.Fatalf("Error al consumir mensajes: %s", err)
    }

    for msg := range messages {
        var event entities.Event
        if err := json.Unmarshal(msg.Body, &event); err != nil {
            log.Printf("Error al deserializar el mensaje: %s", err)
            continue
        }

        // Procesa el evento aquí (por ejemplo, guardarlo en la base de datos)
        log.Printf("Evento recibido: %+v", event)

        // Confirma el mensaje como procesado
        msg.Ack(false)
    }
}
