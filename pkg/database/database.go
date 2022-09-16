package database

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const databaseFile string = "todos.db"

type Database struct {
	db *gorm.DB
}

type Todo struct {
	gorm.Model
	Time  time.Time
	Title string
}

// Open a connection to database and returns it in a Database struct.
func NewConnection() (*Database, error) {
	db, err := gorm.Open(sqlite.Open(databaseFile), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Todo{})

	if err != nil {
		return nil, err
	}

	return &Database{
		db: db,
	}, nil
}

// Insert a new Todo into database.
func (t *Database) Insert(todo Todo) (int, error) {
	res := t.db.Create(&todo)
	if res.Error != nil {
		return 0, res.Error
	}

	return int(todo.ID), nil
}

// List all Todos in database.
func (t *Database) List() ([]Todo, error) {
	allTodos := []Todo{}
	res := t.db.Find(&allTodos)
	if res.Error != nil {
		return []Todo{}, res.Error
	}

	return allTodos, nil
}

// Delete a Todo from database based on an ID.
func (t *Database) Delete(id int) (int, error) {
	res := t.db.Delete(&Todo{}, id)
	if res.Error != nil {
		return 0, res.Error
	}

	return int(id), nil
}
