package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// storage
type PostgresStore struct {
	db *sql.DB
}

// storage interface
type Storager interface {
	GetTodos() ([]*TodoList, error)
	GetTodoByID(int) (*TodoList, error)
	CreateTodo(*TodoList) error
	DeleteTodo(int) error
}

type DBConfig struct {
	host     string
	user     string
	dbname   string
	password string
}

// opens connection to db when run
func NewPostgresStore() (*PostgresStore, error) {
	dbconfig, err := LoadEnv()
	if err != nil {
		return nil, err
	}

	conn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", dbconfig.host, dbconfig.user, dbconfig.dbname, dbconfig.password)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err

	}

	return &PostgresStore{db: db}, nil
}

// creates todo table on start if none exists
func (s *PostgresStore) Init() error {
	return s.initTodoTable()
}

func (s *PostgresStore) initTodoTable() error {
	query := `create table if not exists todo (
		id serial primary key,
		item text,
		created_at timestamp,
		completed boolean
	)`

	_, err := s.db.Exec(query)

	return err
}

func (s *PostgresStore) GetTodos() ([]*TodoList, error) {
	rows, err := s.db.Query("select * from todo")
	if err != nil {
		return nil, err
	}

	todos := []*TodoList{}
	for rows.Next() {
		account, err := scanIntoTodo(rows)
		if err != nil {
			return nil, err
		}
		todos = append(todos, account)
	}

	return todos, nil
}

func (s *PostgresStore) GetTodoByID(id int) (*TodoList, error) {
	query, err := s.db.Query("select * from todo where id = $1", id)
	if err != nil {
		return nil, err
	}

	for query.Next() {
		return scanIntoTodo(query)
	}

	return nil, fmt.Errorf("id %d not found", id)
}

func (s *PostgresStore) CreateTodo(todo *TodoList) error {
	query := `insert into todo 
	(item, created_at, completed)
	values ($1, $2, $3)`

	_, err := s.db.Query(
		query,
		todo.Item,
		todo.CreatedAt,
		todo.Completed)
	if err != nil {
		return nil
	}

	return nil
}

func (s *PostgresStore) DeleteTodo(id int) error {
	_, err := s.db.Query("delete from todo where id = $1", id)
	return err
}

// turns scanned sql items into TodoList struct
func scanIntoTodo(rows *sql.Rows) (*TodoList, error) {
	todo := new(TodoList)
	err := rows.Scan(
		&todo.ID,
		&todo.Item,
		&todo.CreatedAt,
		&todo.Completed)

	return todo, err
}
