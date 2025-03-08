package main

import (
    "log"
    "net/http"
    "proyecto/src/infrastructure/adapters"
    "proyecto/src/core"
)

const port = ":8081"

func main() {
    // Conectar a RabbitMQ
    rabbitMQConn, err := adapters.NewRabbitMQConnection()
    if err != nil {
        log.Fatalf("Error al conectar a RabbitMQ: %s", err)
    }
    defer rabbitMQConn.Close()

    log.Println("Conexión exitosa a RabbitMQ")

    // Conectar a MySQL
    db, err := core.NewMySQLConnection()
    if err != nil {
        log.Fatalf("Error al conectar a la base de datos: %s", err)
    }
    defer db.Close()

    log.Println("Conexión exitosa a MySQL")

    // Consumir mensajes de la cola
    msgs, err := rabbitMQConn.Channel.Consume(
        adapters.QueueName,
        "", // consumer
        false, // auto-ack
        false, // exclusive
        false, // no-local
        false, // no-wait
        nil,   // args
    )
    
    if err != nil {
        log.Fatalf("Error al consumir mensajes: %s", err)
    }

    log.Printf("Esperando mensajes en la cola: %s", adapters.QueueName)

    // Iniciar un goroutine para procesar los mensajes
    go func() {
        for msg := range msgs {
            log.Printf("Mensaje recibido: %s", msg.Body)

            // Aquí puedes agregar tu lógica de procesamiento, como guardar en la base de datos
            // Por ejemplo:
            // err := processMessage(msg.Body)
            // if err != nil {
            //     log.Printf("Error al procesar el mensaje: %s", err)
            //     // Aquí puedes decidir si reintentar o no
            // }

            // Acknowledge (confirmar) que el mensaje ha sido procesado
            msg.Ack(false)
        }
    }()

    // Definir un manejador para una ruta específica
    http.HandleFunc("/tu-ruta", func(w http.ResponseWriter, r *http.Request) {
        // Lógica para manejar la solicitud
        w.Write([]byte("¡Hola desde la ruta!"))
    })

    // Iniciar el servidor HTTP
    log.Printf("Iniciando servidor en el puerto %s", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatalf("Error al iniciar el servidor: %s", err)
    }
}
