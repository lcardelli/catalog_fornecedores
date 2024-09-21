package main

import (
    "log"
    "github.com/lcardelli/catalog_fornecedores.git/config"
    "github.com/lcardelli/catalog_fornecedores.git/routes"

    "github.com/gin-gonic/gin"
)

func main() {
    // Carregar as configurações
    config.Load()

    // Criar uma nova instância do Gin
    router := gin.Default()

    // Configurar as rotas
    routes.Setup(router)

    // Iniciar o servidor
    log.Println("Iniciando o servidor na porta 8080...")
    router.Run(":8080")  // Inicia o servidor na porta 8080
}
