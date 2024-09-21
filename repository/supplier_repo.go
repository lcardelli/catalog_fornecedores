package repository

import (
	"database/sql"
	"errors"
	"github.com/lcardelli/catalog_fornecedores.git/models"
)

// GetAllSuppliers - Obtém todos os fornecedores com paginação e filtros
func GetAllSuppliers(db *sql.DB, filter string, page int, pageSize int) ([]models.Supplier, error) {
	offset := (page - 1) * pageSize
	query := "SELECT id, name, email, phone, address, category FROM suppliers WHERE name LIKE ? LIMIT ? OFFSET ?"
	rows, err := db.Query(query, "%"+filter+"%", pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var suppliers []models.Supplier
	for rows.Next() {
		var supplier models.Supplier
		if err := rows.Scan(&supplier.ID, &supplier.Name, &supplier.Email, &supplier.Phone, &supplier.Address, &supplier.Category); err != nil {
			return nil, err
		}
		suppliers = append(suppliers, supplier)
	}

	return suppliers, nil
}

// GetSupplierByID - Obtém um fornecedor pelo ID
func GetSupplierByID(db *sql.DB, id string) (*models.Supplier, error) {
	query := "SELECT id, name, email, phone, address, category FROM suppliers WHERE id = ?"
	var supplier models.Supplier
	err := db.QueryRow(query, id).Scan(&supplier.ID, &supplier.Name, &supplier.Email, &supplier.Phone, &supplier.Address, &supplier.Category)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("fornecedor não encontrado")
		}
		return nil, err
	}
	return &supplier, nil
}

// CreateSupplier - Cria um novo fornecedor no banco de dados
func CreateSupplier(db *sql.DB, supplier *models.Supplier) error {
	query := "INSERT INTO suppliers (name, email, phone, address, category) VALUES (?, ?, ?, ?, ?)"
	result, err := db.Exec(query, supplier.Name, supplier.Email, supplier.Phone, supplier.Address, supplier.Category)
	if err != nil {
		return err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	supplier.ID = int(lastInsertID)
	return nil
}

// UpdateSupplier - Atualiza um fornecedor existente
func UpdateSupplier(db *sql.DB, id string, supplier *models.Supplier) error {
	query := "UPDATE suppliers SET name = ?, email = ?, phone = ?, address = ?, category = ? WHERE id = ?"
	result, err := db.Exec(query, supplier.Name, supplier.Email, supplier.Phone, supplier.Address, supplier.Category, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("nenhum fornecedor foi atualizado")
	}
	return nil
}

// DeleteSupplier - Deleta um fornecedor
func DeleteSupplier(db *sql.DB, id string) error {
	query := "DELETE FROM suppliers WHERE id = ?"
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("nenhum fornecedor foi deletado")
	}
	return nil
}
