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
	defer db.Close()

	log.Println("Base de datos inicializada correctamente")

	// Crear el DAO
	docenteDAO := dao.NewDocenteDAO(db)

	// 1. Crear la tabla
	if err := docenteDAO.CreateTable(); err != nil {
		log.Fatalf("Error al crear tabla: %v", err)
	}


