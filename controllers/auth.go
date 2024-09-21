package controllers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
	"github.com/lcardelli/catalog_fornecedores.git/database"
	"github.com/lcardelli/catalog_fornecedores.git/models"
)

var jwtSecret = []byte("sua-chave-secreta") // Altere para uma chave secreta segura
var audience = "SEU_CLIENT_ID.apps.googleusercontent.com" // Substitua pelo seu Client ID

// GoogleLogin - Manipula requisições de login com Google
func GoogleLogin(c *gin.Context) {
	idToken := c.PostForm("id_token") // O token deve ser enviado no corpo da requisição

	// Cria um validador
	ctx := c.Request.Context()
	validator, err := idtoken.NewValidator(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar validador"})
		return
	}

	// Verifica o token do Google
	payload, err := validator.Validate(ctx, idToken, audience)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	}

	// Extraí informações do payload
	googleID := payload.Subject
	email := payload.Claims["email"].(string)
	username := payload.Claims["name"].(string)

	// Conectar ao banco de dados
	db := database.Init()

	// Tente encontrar o usuário no banco de dados
	var user models.User
	err = db.QueryRow("SELECT id FROM users WHERE google_id = ?", googleID).Scan(&user.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			// Se o usuário não existe, cria um novo
			_, err = db.Exec("INSERT INTO users (username, email, google_id) VALUES (?, ?, ?)", username, email, googleID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário"})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao acessar o banco de dados"})
			return
		}
	}

	// Se o usuário foi encontrado ou criado, gera o token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"google_id": googleID,
		"exp":       time.Now().Add(time.Hour * 72).Unix(), // Token expira em 72 horas
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
