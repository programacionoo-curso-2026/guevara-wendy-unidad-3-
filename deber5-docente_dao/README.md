# Deber 5 - Docente DAO

## Descripción

Este proyecto implementa el patrón DAO utilizando Go y SQLite para la gestión de docentes.

## Funcionalidades

- Conexión a la base de datos SQLite.
- Creación de la tabla `docentes`.
- Inserción de un docente.
- Búsqueda de un docente por ID.
- Búsqueda de un docente por correo electrónico.

## Ejecución

Para ejecutar el programa:

```bash
go run .
```

## Evidencia de ejecución
PS C:\Wendy\gh repo clone programacionoo-curso-2026\guevara-wendy-unidad-3-\deber5-docente_dao> go run .
2026/07/28 22:05:38 ¡Conectado a SQLite con éxito!
2026/07/28 22:05:38 Base de datos inicializada correctamente
2026/07/28 22:05:38 Tabla docentes creada/verificada exitosamente
2026/07/28 22:05:38 Error al insertar: error al insertar docente: constraint failed: UNIQUE constraint failed: docentes.email (2067)
2026/07/28 22:05:38 Docente encontrado: &{ID:D001 Nombre:Ana García Email:ana.garcia@email.com Departamento:Informática Cargo:Profesora AniosAntiguedad:5}
2026/07/28 22:05:38 Docente encontrado por email: &{ID:D002 Nombre:Carlos Ruiz Email:carlos.ruiz@email.com Departamento:Matemáticas Cargo:Profesor AniosAntiguedad:3}
PS C:\Wendy\gh repo clone programacionoo-curso-2026\guevara-wendy-unidad-3-\deber5-docente_dao> 

