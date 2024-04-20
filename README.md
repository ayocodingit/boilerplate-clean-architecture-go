# Boilerplate Clean Architecture for Go

## Tech Stack

- Golang v1.22.2
- SqlDriver
- Docker
- Gin (HTTP framework)

## Design Pattern: Clean Architecture

The project follows the principles of Clean Architecture, emphasizing separation of concerns into distinct layers:

- **Entities**: Representing the core business entities.
- **Use Cases**: Defining application-specific business rules.
- **Interface Adapters**: Implementing details for external frameworks and tools.
- **Frameworks & Drivers**: Implementing details for external frameworks and tools (Express, databases, etc.).

## Folder Structure Modules

```bash
modules/
└── name module/
    ├── delivery/
    │   ├── http/
    │   │   └── handler.go
    │   ├── grpc/
    │   │   └── handler.go
    │   └── graphQL/
    │       └── handler.go
    ├── entity/
    │   └── interface.go // for the schema validation
    ├── repository/
    │   ├── mongo/
    │   │   └── repository.go
    │   ├── mySQL/
    │   │   └── repository.go
    │   └── postgreSQL/
    │       └── repository.go
    ├── usecase/
    │   └── usecase.go
    └── name module.go // for init the module to load in the main
```