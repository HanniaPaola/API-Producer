package main

import (
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