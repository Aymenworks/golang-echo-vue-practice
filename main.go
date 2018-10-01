package main

import (
	"database/sql"
	"fmt"

	"./handlers"

	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database := initDatabase()
	migrateDatabase(database)

	e := echo.New()

	e.File("/", "public/index.html")
	e.GET("/tasks", handlers.GetTasks(database))
	e.PUT("/tasks", handlers.PutTasks(database))
	e.DELETE("/tasks/:id", handlers.DeleteTasks(database))

	e.Start(":1234")
}

func initDatabase() *sql.DB {
	database, error := sql.Open("sqlite3", "storage.db")

	if error != nil {
		panic(error)
	}

	if database == nil {
		panic("error database creation nil")
	}

	fmt.Println("Database created")

	return database
}

func migrateDatabase(database *sql.DB) {
	tasksTableCreationQuery := `
      CREATE TABLE IF NOT EXISTS tasks(
          id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
          name VARCHAR NOT NULL
      );
  `

	_, error := database.Exec(tasksTableCreationQuery)

	if error != nil {
		panic("error creating the SQL table")
	}

	fmt.Println("Table tasks created")
}
