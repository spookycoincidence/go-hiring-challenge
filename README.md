## Go Hiring Challenge — Product Catalog

This project implements a REST API in Go to manage a product catalog, including CRUD operations, pagination, filtering, and categorization.  
It was developed as part of the **Go Hiring Challenge**, following idiomatic Go best practices, clean architecture principles, and comprehensive testing.

## Main Features

- **Product Management:** Full CRUD with variants and prices.  
- **Categories:** Associate products with categories (Clothing, Shoes, Accessories).  
- **Pagination & Filtering:** Support for `offset`, `limit`, `category`, and `maxPrice` query parameters.  
- **Price Inheritance:** Variants without a specific price inherit the parent product’s price.  
- **Standardized Responses:** Consistent JSON response format across the API.  
- **Testing:** Coverage for handlers, repositories, and response validation.

## Project Structure

```bash
├── cmd/
│   ├── server/main.go         # Server entry point
│   └── seed/main.go           # Script to populate the database
├── app/
│   ├── api/response.go        # JSON response formatting
│   ├── catalog/handler.go     # Catalog endpoints
│   ├── categories/handler.go  # Category endpoints
│   └── database/pg.go         # PostgreSQL connection
├── models/                    # Models and repositories (Product, Category, Variant)
├── sql/                       # Migrations and initial data
├── Makefile                   # Build and execution commands
└── docker-compose.yml         # Docker services configuration
```


## Tech Stack

- **Language:** Go 1.24  
- **Database:** PostgreSQL 16  
- **ORM:** GORM  
- **Testing:** testify/assert  
- **Containers:** Docker + Docker Compose  

## Installation & Execution

```bash
# Install dependencies
make tidy

# Start containers
make docker-up

# Seed the database
make seed

# Run the API
make run
```

## The API will be available at:
http://localhost:8484


## Main Endpoints:

```bash

Catalog

GET /catalog — Retrieve a list of products
Optional query params: offset, limit, category, maxPrice

GET /catalog/:code — Retrieve product details (including variants and category)

Categories

GET /categories — Retrieve category list
POST /categories — Create a new category

```

## Implemented Changes

Database:

- Added categories table with code and name fields.

- Added products.category_id to associate products with categories.

- Updated migrations and seed data to include base categories and initial records.

Models & Repositories:

- New Category model and corresponding repository.

- Refactored ProductsRepository to support: Pagination (offset, limit)

- Filtering by category and max price

- Implemented GetProductByCode() for detailed retrieval.

API & Handlers:

- New endpoints: /catalog, /catalog/:code, /categories.

- Standardized responses using OKResponse and ErrorResponse.

- Included category and variant data in product responses.

Testing:

- Unit tests for catalog and category handlers.

- Tests for JSON response formatting (api/response.go).

- Coverage for pagination, filtering, and error handling scenarios.

Best Practices Applied:

- Idiomatic Go: Minimal use of interfaces, focus on clarity and simplicity.

- Dependency Injection: Handlers receive repositories as dependencies.

- Separation of Concerns: Decoupled logic, persistence, and API layers.

- Error Handling: Consistent and meaningful JSON error responses.

Environment Variables (.env):

```bash

HTTP_PORT=8484
POSTGRES_USER=postgres
POSTGRES_PASSWORD=password
POSTGRES_DB=challenge
POSTGRES_PORT=5432
POSTGRES_SQL_DIR=./sql
```

Potential Improvements:

- Authentication and Authorization

- Sorting and Search

- Bulk Operations

- Product Reviews or Stock Management System

Project developed for technical evaluation purposes.

## Author: Huilen Vilches, fullstack developer 



#Spanish version

## Go Hiring Challenge — Catálogo de Productos

Este proyecto implementa una API REST en Go para gestionar un catálogo de productos, incluyendo operaciones CRUD, paginación, filtrado y categorización.
Fue desarrollado como parte del Go Hiring Challenge, siguiendo las mejores prácticas de código idiomático en Go, arquitectura limpia y testing.

## Funcionalidades principales:

- Gestión de productos: CRUD completo con variantes y precios.

- Categorías: asociación de productos a categorías (Clothing, Shoes, Accessories).

- Paginación y filtrado: parámetros offset, limit, category y maxPrice.

- Herencia de precios: las variantes sin precio usan el del producto padre.

- Respuestas estandarizadas: formato JSON consistente en toda la API.

- Testing: cobertura de tests para handlers, repositorios y respuestas.

## Estructura del proyecto

```bash
├── cmd/
│   ├── server/main.go         # Punto de entrada del servidor
│   └── seed/main.go           # Script para poblar la base de datos
├── app/
│   ├── api/response.go        # Formato de respuestas JSON
│   ├── catalog/handler.go     # Endpoints del catálogo
│   ├── categories/handler.go  # Endpoints de categorías
│   └── database/pg.go         # Conexión a PostgreSQL
├── models/                    # Modelos y repositorios (Product, Category, Variant)
├── sql/                       # Migraciones y datos iniciales
├── Makefile                   # Comandos de build y ejecución
└── docker-compose.yml         # Configuración de servicios Docker
```

## Stack Tecnológico

- **Languaje:** Go 1.24  
- **DataBase:** PostgreSQL 16  
- **ORM:** GORM  
- **Testing:** testify/assert  
- **Containers:** Docker + Docker Compose  

## Instalación y ejecución

```bash
# Instalar dependencias
make tidy

# Levantar contenedores
make docker-up

# Base de datos
make seed

# Ejecutar la API
make run
```


# La API estará disponible en:
http://localhost:8484

# Endpoints principales:

```bash
-Catálogo

GET /catalog — Listado de productos

Parámetros opcionales:
offset, limit, category, maxPrice

GET /catalog/:code — Detalle de producto (con variantes y categoría)

# Categorías

GET /categories — Listado de categorías

POST /categories — Crear una nueva categoría
```


# Cambios e implementaciones realizadas
Base de datos:

- Agregada tabla categories con código y nombre.

- Agregada products.category_id para asociar productos.

- Migraciones y seeds actualizados para incluir categorías y datos base.

Modelos y repositorios:

- Nuevo modelo Category y repositorio correspondiente.

Refactor de ProductsRepository para soportar:

- Paginación (offset, limit)

- Filtrado por categoría y precio máximo

- Implementado GetProductByCode() para obtener detalle individual.

API y handlers:

- Nuevos endpoints /catalog, /catalog/:code, /categories.

- Estandarización de respuestas con OKResponse y ErrorResponse.

- Inclusión de categoría y variantes en respuestas de productos.

Testing:

- Tests unitarios para catalog y categories handlers.

- Tests de formato de respuesta (api/response.go).

- Cobertura de casos de paginación, filtrado y manejo de errores.

Buenas prácticas aplicadas:

- Código idiomático en Go: uso mínimo de interfaces, claridad y simplicidad.

- Inyección de dependencias en handlers.

- Separación de responsabilidades: lógica, persistencia y API desacopladas.

- Errores bien manejados y respuestas consistentes en JSON.

Variables de entorno (.env)

```bash
HTTP_PORT=8484
POSTGRES_USER=postgres
POSTGRES_PASSWORD=password
POSTGRES_DB=challenge
POSTGRES_PORT=5432
POSTGRES_SQL_DIR=./sql
```

Oportunidades de mejoras posibles:

- Autenticación y autorización

- Ordenamiento y búsqueda

- Operaciones masivas (bulk updates)

- Sistema de reviews o stock


Proyecto desarrollado con fines de evaluación técnica.

## Autor: Huilen Vilches, fullstack developer 
## Portfolio: https://www.huilen.dev/




## Go Hiring Challenge — Product Catalog

This project implements a REST API in Go to manage a product catalog, including CRUD operations, pagination, filtering, and categorization.  
It was developed as part of the **Go Hiring Challenge**, following idiomatic Go best practices, clean architecture principles, and comprehensive testing.

## Main Features

- **Product Management:** Full CRUD with variants and prices.  
- **Categories:** Associate products with categories (Clothing, Shoes, Accessories).  
- **Pagination & Filtering:** Support for `offset`, `limit`, `category`, and `maxPrice` query parameters.  
- **Price Inheritance:** Variants without a specific price inherit the parent product’s price.  
- **Standardized Responses:** Consistent JSON response format across the API.  
- **Testing:** Coverage for handlers, repositories, and response validation.

## Project Structure

```bash
├── cmd/
│   ├── server/main.go         # Server entry point
│   └── seed/main.go           # Script to populate the database
├── app/
│   ├── api/response.go        # JSON response formatting
│   ├── catalog/handler.go     # Catalog endpoints
│   ├── categories/handler.go  # Category endpoints
│   └── database/pg.go         # PostgreSQL connection
├── models/                    # Models and repositories (Product, Category, Variant)
├── sql/                       # Migrations and initial data
├── Makefile                   # Build and execution commands
└── docker-compose.yml         # Docker services configuration
```


## Tech Stack

- **Language:** Go 1.24  
- **Database:** PostgreSQL 16  
- **ORM:** GORM  
- **Testing:** testify/assert  
- **Containers:** Docker + Docker Compose  

## Installation & Execution

```bash
# Install dependencies
make tidy

# Start containers
make docker-up

# Seed the database
make seed

# Run the API
make run
```

## The API will be available at:
http://localhost:8484


## Main Endpoints:

```bash

Catalog

GET /catalog — Retrieve a list of products
Optional query params: offset, limit, category, maxPrice

GET /catalog/:code — Retrieve product details (including variants and category)

Categories

GET /categories — Retrieve category list
POST /categories — Create a new category

```

## Implemented Changes

Database:

- Added categories table with code and name fields.

- Added products.category_id to associate products with categories.

- Updated migrations and seed data to include base categories and initial records.

Models & Repositories:

- New Category model and corresponding repository.

- Refactored ProductsRepository to support: Pagination (offset, limit)

- Filtering by category and max price

- Implemented GetProductByCode() for detailed retrieval.

API & Handlers:

- New endpoints: /catalog, /catalog/:code, /categories.

- Standardized responses using OKResponse and ErrorResponse.

- Included category and variant data in product responses.

Testing:

- Unit tests for catalog and category handlers.

- Tests for JSON response formatting (api/response.go).

- Coverage for pagination, filtering, and error handling scenarios.

Best Practices Applied:

- Idiomatic Go: Minimal use of interfaces, focus on clarity and simplicity.

- Dependency Injection: Handlers receive repositories as dependencies.

- Separation of Concerns: Decoupled logic, persistence, and API layers.

- Error Handling: Consistent and meaningful JSON error responses.

Environment Variables (.env):

```bash

HTTP_PORT=8484
POSTGRES_USER=postgres
POSTGRES_PASSWORD=password
POSTGRES_DB=challenge
POSTGRES_PORT=5432
POSTGRES_SQL_DIR=./sql
```

Potential Improvements:

- Authentication and Authorization

- Sorting and Search

- Bulk Operations

- Product Reviews or Stock Management System

Project developed for technical evaluation purposes.

## Author: Huilen Vilches, fullstack developer 
