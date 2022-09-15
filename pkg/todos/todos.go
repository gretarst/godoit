package todos

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const databaseFile string = "todos.db"

const createDatabase string = `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER NOT NULL PRIMARY KEY,
		time DATETIME NOT NULL,
		title TEXT
	)
`

type Todos struct {
	db *sql.DB
}

type Todo struct {
	Time  time.Time
	Title string
	ID    int
}

func NewTodos() (*Todos, error) {
	db, err := sql.Open("sqlite3", databaseFile)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(createDatabase); err != nil {
		return nil, err
	}

	return &Todos{
		db: db,
	}, nil
}

func (t *Todos) Insert(todo Todo) (int, error) {
	res, err := t.db.Exec("INSERT INTO todos VALUES(NULL,?,?)", todo.Time, todo.Title)
	if err != nil {
		return 0, err
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}
	return int(id), nil
}

func (t *Todos) List() ([]Todo, error) {
	res, err := t.db.Query("SELECT * from todos")
	if err != nil {
		return []Todo{}, err
	}

	todos := []Todo{}
	for res.Next() {
		i := Todo{}
		if err = res.Scan(&i.ID, &i.Time, &i.Title); err != nil {
			return nil, err
		}
		todos = append(todos, i)
	}

	return todos, nil
}

func (t *Todos) Delete(id int) (int, error) {
	res, err := t.db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return 0, err
	}

	var i int64
	if i, err = res.LastInsertId(); err != nil {
		return 0, err
	}
	return int(i), nil
}
