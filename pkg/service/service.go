package service

import (
	todo "github.com/KDias-code/todoapp"
	"github.com/KDias-code/todoapp/pkg/repository"
)

type Autharization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	GetAll(userId int) ([]todo.User, error)
	Delete(userId int) error
	Update(userId int, id int, input todo.UpdateUserInput) error
}

type TodoAuthor interface {
	Create(userId int, author todo.TodoAuthor) (int, error)
	GetAll(userId int) ([]todo.TodoAuthor, error)
	GetById(userId, authorId int) (todo.TodoAuthor, error)
	Delete(userId, authorId int) error
	Update(userId, authorId int, input todo.UpdateAuthorInput) error
}

type TodoBook interface {
	Create(userId, authorId int, book todo.TodoBook) (int, error)
	GetAll(userId, authorId int) ([]todo.TodoBook, error)
	GetById(userId, bookId int) (todo.TodoBook, error)
	Delete(userId, bookId int) error
	Update(userId, bookId int, input todo.UpdateBookInput) error
}

type Service struct {
	Autharization
	TodoAuthor
	TodoBook
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		NewAuthService(repos.Autharization),
		NewTodoAuthorService(repos.TodoAuthor),
		NewTodoBookService(repos.TodoBook, repos.TodoAuthor),
	}
}
