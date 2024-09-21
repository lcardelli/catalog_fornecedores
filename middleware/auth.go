package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte("sua-chave-secreta") // Altere para uma chave secreta segura

// AuthMiddleware - Middleware para proteger rotas com autenticação JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Pega o token do cabeçalho Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			c.Abort()
			return
		}

		// Divide o cabeçalho em partes para obter o token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Verifica o token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verifica a assinatura do token
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		// Se o token for válido, continua para a próxima função
		c.Next()
	}
}
