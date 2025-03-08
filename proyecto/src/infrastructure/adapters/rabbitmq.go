package adapters

import (
    "encoding/json"
    "log"

    "github.com/streadway/amqp"
    "proyecto/src/domain/entities"
)

// RabbitMQ encapsula la conexión y el canal a RabbitMQ.
type RabbitMQ struct {
    connection *amqp.Connection
    channel    *amqp.Channel
    queue      string
}

// NewRabbitMQ crea una nueva instancia de RabbitMQ
func NewRabbitMQ(url string) (*RabbitMQ, error) {
    conn, err := amqp.Dial(url)
    if err != nil {
        return nil, err
    }

    channel, err := conn.Channel()
    if err != nil {
        return nil, err
    }

    queue := "consumidor" // Cambia esto por el nombre de tu cola

    _, err = channel.QueueDeclare(
        queue,
        true,  // durable
        false, // delete when unused
        false, // exclusive
        false, // no-wait
        nil,   // arguments
    )
    if err != nil {
        return nil, err
    }

    return &RabbitMQ{
        connection: conn,
        channel:    channel,
        queue:      queue,
    }, nil
}

// SendEvent envía un evento a RabbitMQ
func (r *RabbitMQ) SendEvent(event entities.Event) error {
    body, err := json.Marshal(event)
    if err != nil {
        return err
    }

    err = r.channel.Publish(
        "",         // exchange
        r.queue,    // routing key
        false,      // mandatory
        false,      // immediate
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
        },
    )
    if err != nil {
        log.Printf("Error al enviar el evento a RabbitMQ: %s", err)
        return err
    }

    log.Printf("Evento enviado a la cola: %s", event)
    return nil
}

// Close cierra la conexión a RabbitMQ
func (r *RabbitMQ) Close() {
    if err := r.channel.Close(); err != nil {
        log.Printf("Error al cerrar el canal: %s", err)
    }
    if err := r.connection.Close(); err != nil {
        log.Printf("Error al cerrar la conexión: %s", err)
    }
}
