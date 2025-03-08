package core

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "strconv"

    _ "github.com/go-sql-driver/mysql" // Importa el driver de MySQL
    "github.com/joho/godotenv"
)

func NewMySQLConnection() (*sql.DB, error) {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error cargando el archivo .env: %v", err)
    }

    host := os.Getenv("DB_HOST")
    portStr := os.Getenv("DB_PORT")
    username := os.Getenv("DB_USERNAME")
    password := os.Getenv("DB_PASSWORD")
    database := os.Getenv("DB_DATABASE")

    // Convertir puerto a int
    port, err := strconv.Atoi(portStr)
    if err != nil {
        return nil, fmt.Errorf("Error convirtiendo el puerto a entero: %v", err)
    }

    // Crear la dirección del servidor
    addr := fmt.Sprintf("%s:%d", host, port)

    // Conectar a la base de datos
    db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, addr, database))
    if err != nil {
        return nil, fmt.Errorf("Error conectando a MySQL: %v", err)
    }

    // Verificar la conexión
    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("Error haciendo ping a la base de datos: %v", err)
    }

    fmt.Println("Conexión exitosa a MySQL")
    return db, nil
}