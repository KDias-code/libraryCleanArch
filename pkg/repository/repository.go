package repository

import (
	todo "github.com/KDias-code/todoapp"
	"github.com/jmoiron/sqlx"
)

type Autharization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
	GetAll(userId int) ([]todo.User, error)
	//DeleteUser(userId int) error
	Update(userId int, input todo.UpdateUserInput) error
	Delete(userId int) error
}

type TodoAuthor interface {
	Create(userId int, author todo.TodoAuthor) (int, error)
	GetAll(userId int) ([]todo.TodoAuthor, error)
	GetById(userId, authorId int) (todo.TodoAuthor, error)
	Delete(userId int, authorId int) error
	Update(userId, authorId int, input todo.UpdateAuthorInput) error
}

type TodoBook interface {
	Create(authorId int, book todo.TodoBook) (int, error)
	GetAll(userId, authors int) ([]todo.TodoBook, error)
	GetById(userId, bookId int) (todo.TodoBook, error)
	Delete(userId, bookId int) error
	Update(userId, bookId int, input todo.UpdateBookInput) error
}

type Repository struct {
	Autharization
	TodoAuthor
	TodoBook
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Autharization: NewAuthPostgres(db),
		TodoAuthor:    NewTodoAuthorPostgres(db),
		TodoBook:      NewTodoBookPostgres(db),
	}
}
