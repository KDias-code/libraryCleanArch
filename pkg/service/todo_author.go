package service

import (
	todo "github.com/KDias-code/todoapp"
	"github.com/KDias-code/todoapp/pkg/repository"
)

type TodoAuthorService struct {
	repo repository.TodoAuthor
}

func NewTodoAuthorService(repo repository.TodoAuthor) *TodoAuthorService {
	return &TodoAuthorService{repo: repo}
}

func (s *TodoAuthorService) Create(userId int, author todo.TodoAuthor) (int, error) {
	return s.repo.Create(userId, author)
}

func (s *TodoAuthorService) GetAll(userId int) ([]todo.TodoAuthor, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoAuthorService) GetById(userId, authorId int) (todo.TodoAuthor, error) {
	return s.repo.GetById(userId, authorId)
}

func (s *TodoAuthorService) Delete(userId, authorId int) error {
	return s.repo.Delete(userId)
}

func (s *TodoAuthorService) Update(userId, authorId int, input todo.UpdateAuthorInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, authorId, input)
}
