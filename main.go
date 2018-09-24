package main

import (
	"database/sql"
	"fmt"

	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database := initDatabase()
	migrateDatabase(database)

	e := echo.New()

	e.GET("/tasks", getTasks)
	e.PUT("/tasks", updateTask)
	e.DELETE("/tasks/:id", deleteTask)

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

func getTasks(context echo.Context) error {
	return context.JSON(200, "GET TASKS")
}

func updateTask(context echo.Context) error {
	return context.JSON(200, "UPDATE TASKS")
}

func deleteTask(context echo.Context) error {
	return context.JSON(200, "DELETE TASKS"+context.Param("id"))
}
