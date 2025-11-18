# `buffalos` Template
[Ler em Português](README.pt.md)

## About

The **`buffalos` Template** is a robust and opinionated base for building backend applications and APIs using **Go**. It adheres to Go ecosystem best practices, adopting a clean and modular architecture that facilitates code scalability, testability, and maintenance.

This template is ideal for developers looking to quickly start projects with a structure already prepared to handle **databases, routing, configurations**, and well-defined **domain logic**.

## Project Structure

The `buffalos` architecture is based on the concept of Separation of Concerns, organizing the code into logical layers:

```
├── .gitignore             # Files ignored by Git
├── go.mod / go.sum        # Go modules and dependency management
├── LICENSE                # License information (e.g., MIT)
├── migrations             # SQL scripts for database schema management
├── README.md              # This file
└── src
├── .env               # Environment variables
├── internal           # Code not importable externally (application logic)
│   ├── configurations # Application startup configurations (e.g., database, server)
│   ├── controllers    # Input layer (HTTP handlers or interface logic)
│   ├── domain         # Data structures and interfaces (Core business logic entities)
│   ├── misc           # Utility functions and helper code
│   ├── repositories   # Persistence layer (database access logic)
│   └── services       # Application services layer (coordinates domain and persistence)
└── main.go            # Application entry point
```

## Key Technologies

* **Language:** Go (Golang)
* **Dependency Management:** Go Modules (`go.mod`)
* **Database:** PostgreSQL
* **Migrations:** `golang-migrate/migrate`

## Getting Started

### 1\. Prerequisites

Ensure you have the following installed on your machine:

* **Go (version 1.24 or higher)**
* **[Optional: Docker/Docker Compose]**
* **[Optional: The main database (PostgreSQL)**

### 2\. Installation and Setup

If you are using the `buffalos-cli`:

```bash
buffalos install
```

Otherwise, clone and initialize manually:

```bash
# 1. Clone the repository
git clone https://github.com/nes-xiof11/buffalos.git
cd buffalos

# 2. Configure the environment
# Edit src/.env with your credentials and configuration settings

# 3. Download dependencies
go mod tidy

# 4. Run migrations
# migrate -database "postgres://..." -path migrations up

# 5. Start the application
go run src/
```

## Database Migrations

The `migrations` folder contains the ordered SQL files that define the database schema.

* `000_init.sql`
* `001_create_table_users.sql`
* `002_create_table_project.sql`

It is crucial to execute these migrations before starting the application for the first time.

## Contributions

Contributions are welcome\! If you find a bug or have a suggestion for improvement, please open an *issue* or submit a *Pull Request*.