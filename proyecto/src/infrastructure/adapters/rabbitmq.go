package adapters

import (
    "log"

    "github.com/streadway/amqp"
)

const (
    rabbitMQURL = "amqp://hannia:hannia@3.80.229.200:5672/"
    QueueName    = "consumidor" // Cambia a QueueName para exportarlo
)

// RabbitMQConnection encapsula la conexión y el canal a RabbitMQ.
type RabbitMQConnection struct {
    Conn    *amqp.Connection
    Channel *amqp.Channel
}

// NewRabbitMQConnection establece una nueva conexión a RabbitMQ y declara la cola.
func NewRabbitMQConnection() (*RabbitMQConnection, error) {
    conn, err := amqp.Dial(rabbitMQURL)
    if err != nil {
        return nil, err
    }

    channel, err := conn.Channel()
    if err != nil {
        return nil, err
    }

    // Asegúrate de que la cola exista
    _, err = channel.QueueDeclare(
        QueueName, // Usa QueueName aquí
        true,      // durable
        false,     // delete when unused
        false,     // exclusive
        false,     // no-wait
        nil,       // arguments
    )
    if err != nil {
        return nil, err
    }

    return &RabbitMQConnection{
        Conn:    conn,
        Channel: channel,
    }, nil
}

// Close cierra la conexión y el canal de RabbitMQ.
func (r *RabbitMQConnection) Close() {
    if err := r.Channel.Close(); err != nil {
        log.Printf("Error al cerrar el canal: %s", err)
    }
    if err := r.Conn.Close(); err != nil {
        log.Printf("Error al cerrar la conexión: %s", err)
    }
}
