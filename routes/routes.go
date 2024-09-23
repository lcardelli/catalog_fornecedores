package routes

import (
	"github.com/lcardelli/catalog_fornecedores.git/controllers"
	"github.com/lcardelli/catalog_fornecedores.git/database"
	"github.com/lcardelli/catalog_fornecedores.git/middleware"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	// Inicializar a conexão com o banco de dados
	db := database.Init()

	// Middleware para passar a conexão do banco para o contexto
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Definir rotas para os fornecedores com proteção de autenticação
	api := router.Group("/api")
	{
		api.GET("/suppliers", middleware.AuthMiddleware(), controllers.GetSuppliers)       // Listar fornecedores
		api.POST("/suppliers", middleware.AuthMiddleware(), controllers.CreateSupplier)     // Criar novo fornecedor
		api.PUT("/suppliers/:id", middleware.AuthMiddleware(), controllers.UpdateSupplier)  // Atualizar fornecedor
		api.DELETE("/suppliers/:id", middleware.AuthMiddleware(), controllers.DeleteSupplier) // Deletar fornecedor
	}


	auth := router.Group("/auth")
	{
		auth.GET("/google/login", controllers.HandleGoogleLogin)       // Iniciar login com Google
		auth.GET("/google/callback", controllers.HandleGoogleCallback) // Callback do Google
	}
}



