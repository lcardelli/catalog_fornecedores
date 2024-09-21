package config

import (
    "log"
    "os"
)

var DBHost string
var DBUser string
var DBPassword string
var DBName string

func Load() {
    DBHost = os.Getenv("DB_HOST")
    DBUser = os.Getenv("DB_USER")
    DBPassword = os.Getenv("DB_PASSWORD")
    DBName = os.Getenv("DB_NAME")

    if DBHost == "" || DBUser == "" || DBPassword == "" || DBName == "" {
        log.Fatal("Configurações de banco de dados não foram definidas corretamente")
    }
}