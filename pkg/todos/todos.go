package todos

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const databaseFile string = "todos.db"

type Todos struct {
	db *gorm.DB
}

type Todo struct {
	gorm.Model
	Time  time.Time
	Title string
}

func NewTodos() (*Todos, error) {
	db, err := gorm.Open(sqlite.Open(databaseFile), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Todo{})

	if err != nil {
		return nil, err
	}

	return &Todos{
		db: db,
	}, nil
}

func (t *Todos) Insert(todo Todo) (int, error) {
	res := t.db.Create(&todo)
	if res.Error != nil {
		return 0, res.Error
	}

	return int(todo.ID), nil
}

func (t *Todos) List() ([]Todo, error) {
	allTodos := []Todo{}
	res := t.db.Find(&allTodos)
	if res.Error != nil {
		return []Todo{}, res.Error
	}

	return allTodos, nil
}

func (t *Todos) Delete(id int) (int, error) {
	res := t.db.Delete(&Todo{}, id)
	if res.Error != nil {
		return 0, res.Error
	}

	return int(id), nil
}
