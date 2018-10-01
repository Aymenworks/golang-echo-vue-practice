package models

import (
  "database/sql"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
  Id int      `json:"name"`
  Name string `json:"id"`
}

type TaskCollection struct {
  Tasks []Task `json:"tasks"`
}
