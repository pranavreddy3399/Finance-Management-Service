package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// InitMySQL establishes a global MySQL connection
func InitMySQL() (*sql.DB, error) {
	dbUser := "root"
	dbPass := "password"
	dbHost := "127.0.0.1"
	dbPort := "3306"
	dbName := "fms"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	dbConn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("❌ Error opening DB: %v", err)
		return nil, err
	}

	if err = dbConn.Ping(); err != nil {
		log.Fatalf("❌ Cannot connect to MySQL: %v", err)
		return nil, err
	}

	fmt.Println("✅ Connected to MySQL database:", dbName)
	return dbConn, nil
}
