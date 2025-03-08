package adapters

import (
<<<<<<< HEAD
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
=======
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
>>>>>>> 38898a8ba7fbd5bbf09ea0fbf4f9d96afd59e257
    if err != nil {
        return nil, err
    }

    channel, err := conn.Channel()
    if err != nil {
        return nil, err
    }

<<<<<<< HEAD
    // Asegúrate de que la cola exista
    _, err = channel.QueueDeclare(
        QueueName, // Usa QueueName aquí
        true,      // durable
        false,     // delete when unused
        false,     // exclusive
        false,     // no-wait
        nil,       // arguments
=======
    queue := "consumidor" // Cambia esto por el nombre de tu cola

    _, err = channel.QueueDeclare(
        queue,
        true,  // durable
        false, // delete when unused
        false, // exclusive
        false, // no-wait
        nil,   // arguments
>>>>>>> 38898a8ba7fbd5bbf09ea0fbf4f9d96afd59e257
    )
    if err != nil {
        return nil, err
    }

<<<<<<< HEAD
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
=======
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
>>>>>>> 38898a8ba7fbd5bbf09ea0fbf4f9d96afd59e257
        log.Printf("Error al cerrar la conexión: %s", err)
    }
}
