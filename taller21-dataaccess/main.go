package main

import (
	"CompetenciasDocentes/dao"
	"CompetenciasDocentes/dataaccess"
	"CompetenciasDocentes/model"

	"log"
)

func main() {
	// Inicializar la base de datos
	db := dataaccess.InitDB()
	defer db.Close() // Importante: cerrar la conexión al final

	log.Println("Base de datos inicializada correctamente")

	// Crear el DAO
	docenteDAO := dao.NewDocenteDAO(db)

	// Crear la tabla
	if err := docenteDAO.CreateTable(); err != nil {
		log.Fatalf("Error al crear tabla: %v", err)
	}

	// 1. INSERTAR un docente
	docente1 := &model.Docente{
		ID:              "D001",
		Nombre:          "Ana García",
		Email:           "ana.garcia@email.com",
		Departamento:    "Informática",
		Cargo:           "Profesora",
		AniosAntiguedad: 5,
	}

	if err := docenteDAO.Insert(docente1); err != nil {
		log.Printf("Error al insertar: %v", err)
	}

	// 2. INSERTAR otro docente
	docente2 := &model.Docente{
		ID:              "D002",
		Nombre:          "Carlos Ruiz",
		Email:           "carlos.ruiz@email.com",
		Departamento:    "Matemáticas",
		Cargo:           "Profesor",
		AniosAntiguedad: 3,
	}

	if err := docenteDAO.Insert(docente2); err != nil {
		log.Printf("Error al insertar: %v", err)
	}

	// 3. BUSCAR por ID
	docente, err := docenteDAO.GetByID("D001")
	if err != nil {
		log.Printf("Error al buscar: %v", err)
	} else {
		log.Printf("Docente encontrado: %+v", docente)
	}

	// 4. BUSCAR por Email
	docentePorEmail, err := docenteDAO.GetByEmail("carlos.ruiz@email.com")
	if err != nil {
		log.Printf("Error al buscar por email: %v", err)
	} else {
		log.Printf("Docente encontrado por email: %+v", docentePorEmail)
	}
}
