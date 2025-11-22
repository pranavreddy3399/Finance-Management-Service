package db

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func InitMySQL() (*sqlx.DB, error) {
	dbUser := "root"
	dbPass := "password"
	dbHost := "127.0.0.1"
	dbPort := "3306"
	dbName := "fms"

	// Add parseTime=true for proper time.Time scanning
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPass, dbHost, dbPort, dbName,
	)

	// Use sqlx.Connect (opens + pings automatically)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("❌ Failed to connect to MySQL: %v", err)
		return nil, err
	}

	// Set recommended connection pool sizes
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(0)

	fmt.Println("✅ Connected to MySQL using sqlx:", dbName)
	return db, nil
}
