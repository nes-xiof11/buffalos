package configurations

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// GetPostgresDSN constructs the PostgreSQL DSN string from environment variables.
func GetPostgresDSN() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")
	schema := os.Getenv("DB_SCHEMA")

	if sslmode == "" {
		sslmode = "disable"
	}

	// Example DSN: "host=localhost port=5432 user=postgres password=pass dbname=db sslmode=disable search_path=public"
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s search_path=%s",
		host, port, user, password, dbname, sslmode, schema,
	)
}

//// GetPostgresURL constructs the PostgreSQL connection URL (without password or schema)
//func GetPostgresURL() string {
//	host := os.Getenv("DB_HOST")
//	port := os.Getenv("DB_PORT")
//	user := os.Getenv("DB_USER")
//	//password := os.Getenv("DB_PASSWORD")
//	dbname := os.Getenv("DB_NAME")
//	//schema := os.Getenv("DB_SCHEMA")
//
//	// Example URL: "postgres://user@localhost:5432/dbname?sslmode=disable"
//	return fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=disable", user, host, port, dbname)
//
//}

// SetupDB opens a connection to the PostgreSQL database and pings it to verify connectivity.
func SetupDB() *sql.DB {
	dsn := GetPostgresDSN()
	db, err := sql.Open(os.Getenv("DRIVER_NAME"), dsn)
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging DB: %v", err)
	}
	log.Printf("Connection with %s established!", os.Getenv("DBMS_NAME"))

	return db
}
