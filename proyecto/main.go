package main

import (
<<<<<<< HEAD
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
=======
    "database/sql"
    "log"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    _ "github.com/go-sql-driver/mysql"
    "proyecto/src/infrastructure/routes"
    "proyecto/src/application/services"
    "proyecto/src/application/useCase"
    "proyecto/src/application/repositories"
    "proyecto/src/infrastructure/adapters" // Asegúrate de importar el paquete de RabbitMQ
    "proyecto/src/domain/entities"
)

func main() {
    // Cargar variables de entorno
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error cargando el archivo .env")
    }

    // Conectar a la base de datos
    db, err := sql.Open("mysql", "root:hannia@tcp(localhost:3306)/sensor")
    if err != nil {
        log.Fatalf("Error conectando a la base de datos: %v", err)
    }
    defer db.Close()

    // Crear una instancia de Gin
    r := gin.Default()

    // Inicializar el repositorio y servicio de eventos
    eventRepo := repositories.NewEventRepository(db)
    eventService := services.NewEventService(eventRepo)
    eventUseCase := useCase.NewEventUseCase(eventService)

    // Inicializar RabbitMQ
    rabbitMQ, err := adapters.NewRabbitMQ("amqp://hannia:hannia@3.80.229.200:5672/") // Cambia la URL según tu configuración
    if err != nil {
        log.Fatalf("Error al crear RabbitMQ: %s", err)
    }
    defer rabbitMQ.Close()

    // Ejemplo de envío de un evento a RabbitMQ
    event := entities.Event{
        Type:      "sensor_data",
        CreatedAt: time.Now(),
        Unit:      "temperature",
    }

    if err := rabbitMQ.SendEvent(event); err != nil {
        log.Fatalf("Error al enviar el evento: %s", err)
    }

    // Registrar rutas
    routes.RegisterRoutes(r, eventUseCase)

    // Iniciar el servidor en el puerto 3000
    if err := r.Run(":3000"); err != nil {
        log.Fatalf("❌ Error al iniciar el servidor: %v", err)
    }
}
>>>>>>> 38898a8ba7fbd5bbf09ea0fbf4f9d96afd59e257
