package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/lcardelli/catalog_fornecedores.git/models"
	"github.com/lcardelli/catalog_fornecedores.git/repository"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// GetSuppliers - Manipula requisições para listar fornecedores com paginação e filtros
func GetSuppliers(c *gin.Context) {
	filter := c.DefaultQuery("filter", "")
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")

	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	db := c.MustGet("db").(*sql.DB)
	suppliers, err := repository.GetAllSuppliers(db, filter, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, suppliers)
}

// GetSupplier - Manipula requisições para obter um fornecedor pelo ID
func GetSupplier(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("db").(*sql.DB)

	supplier, err := repository.GetSupplierByID(db, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, supplier)
}

// CreateSupplier - Manipula requisições para criar um fornecedor
func CreateSupplier(c *gin.Context) {
	var supplier models.Supplier

	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validações
	if err := validate.Struct(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*sql.DB)
	if err := repository.CreateSupplier(db, &supplier); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, supplier)
}

// UpdateSupplier - Manipula requisições para atualizar um fornecedor
func UpdateSupplier(c *gin.Context) {
	id := c.Param("id")
	var supplier models.Supplier

	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*sql.DB)
	if err := repository.UpdateSupplier(db, id, &supplier); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, supplier)
}

// DeleteSupplier - Manipula requisições para deletar um fornecedor
func DeleteSupplier(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("db").(*sql.DB)

	if err := repository.DeleteSupplier(db, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Fornecedor deletado com sucesso"})
}
