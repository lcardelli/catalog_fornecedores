package database

import (
    "database/sql"
    "log"
    "os"

    _ "github.com/go-sql-driver/mysql"
)

func Init() *sql.DB {
    connStr := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ")/" + os.Getenv("DB_NAME")
    db, err := sql.Open("mysql", connStr)
    if err != nil {
        log.Fatal(err)
    }

    // Testa a conexão ao banco
    if err := db.Ping(); err != nil {
        log.Fatal("Erro ao conectar ao banco de dados:", err)
    }

    return db
}
