package dao

import (
	"database/sql"
)

// DocenteDAO maneja las operaciones CRUD para Docente
type DocenteDAO struct {
	db *sql.DB
}

// NewDocenteDAO crea una nueva instancia de DocenteDAO
func NewDocenteDAO(db *sql.DB) *DocenteDAO {
	return &DocenteDAO{db: db}
}
