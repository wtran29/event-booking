package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func getDSN() string {
	godotenv.Load(".env")
	host := os.Getenv("host")
	port := os.Getenv("port")
	user := os.Getenv("user")
	password := os.Getenv("password")
	// dbname := os.Getenv("dbname")
	return fmt.Sprintf("host=%v port=%v user=%v password=%v sslmode=disable timezone=UTC connect_timeout=5", host, port, user, password)
}
func InitDB() {
	dsn := getDSN()
	fmt.Println(dsn)
	var err error
	DB, err = sql.Open("pgx", dsn)
	fmt.Println(err)
	if err != nil {
		panic("could not connect to database")
	}

	err = DB.Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err)
		panic("could not ping database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	// dropEventsTable := `DROP TABLE IF EXISTS events`
	// _, err := DB.Exec(dropEventsTable)
	// if err != nil {
	// 	panic("could not drop events table")
	// }
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("could not create users table")
	}
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime TIMESTAMPTZ NOT NULL,
		user_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("could not create events table")
	}

}
