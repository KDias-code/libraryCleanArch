package service

import (
	todo "github.com/KDias-code/todoapp"
	"github.com/KDias-code/todoapp/pkg/repository"
)

type TodoBookService struct {
	repo       repository.TodoBook
	authorRepo repository.TodoAuthor
}

func NewTodoBookService(repo repository.TodoBook, authorRepo repository.TodoAuthor) *TodoBookService {
	return &TodoBookService{repo: repo, authorRepo: authorRepo}
}

func (s *TodoBookService) Create(userId, authorId int, book todo.TodoBook) (int, error) {
	_, err := s.authorRepo.GetById(userId, authorId)
	if err != nil {
		return 0, err
	}

	return s.repo.Create(authorId, book)
}

func (s *TodoBookService) GetAll(userId, authorId int) ([]todo.TodoBook, error) {
	return s.repo.GetAll(userId, authorId)
}

func (s *TodoBookService) GetById(userId, bookId int) (todo.TodoBook, error) {
	return s.repo.GetById(userId, bookId)
}

func (s *TodoBookService) Delete(userId, bookId int) error {
	return s.repo.Delete(userId, bookId)
}

func (s *TodoBookService) Update(userId, bookId int, input todo.UpdateBookInput) error {
	return s.repo.Update(userId, bookId, input)
}
