package db

import (
	"database/sql"
	"fmt"
	//_ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error // объявляем err отдельно
	DB, err = sql.Open("sqlite", "./src/RestApi/api.db")

	if err != nil {
		panic("failed to connect database")
	}

	if err = DB.Ping(); err != nil {
		panic(fmt.Sprintf("Failed to ping database: %v", err))
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()
}

func createTable() {
	if DB == nil {
		panic("failed to connect database123")
	}
	createUserTable()
	createEventTable()
	createRegistrationsTable()
}

// Создание таблицы пользователей
func createUserTable() {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)`

	_, err := DB.Exec(createUserTable)
	if err != nil {
		panic("Failed to create user stable")
	}
}

// Создание таблицы событий
func createEventTable() {
	createEventTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT not null,
		description TEXT not null,
		location TEXT not null,
		dateTime DATETIME NOT NULL,
		user_id integer,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);`

	_, err := DB.Exec(createEventTable)
	if err != nil {
		panic("failed to create event table")
	}
}

// Создание таблицы регистрации событий
func createRegistrationsTable() {
	createRegistrationsTable := `
		CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,	
		foreign key (event_id) REFERENCES events(id),
		foreign key (user_id) REFERENCES users(id)
		)`

	_, err := DB.Exec(createRegistrationsTable)
	if err != nil {
		panic("failed to create registrations table")
	}
}
