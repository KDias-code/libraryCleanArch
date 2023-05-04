package todo

import "errors"

type TodoAuthor struct {
	Id        int    `json:"id" db:"id"`
	Fullname  string `json:"fullname" db:"fullname" binding:"required"`
	Pseudonym string `json:"pseudonym" db:"pseudonym"`
	Spec      string `json:"spec"`
}

type UserAuthor struct {
	Id       int
	UserId   int
	AuthorId int
}

type TodoBook struct {
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title" binding:"required"`
	Genre string `json:"genre" db:"genre"`
	Isbn  int    `json:"isbn" db:"isbn"`
}

type AuthorsBooks struct {
	Id       int
	AuthorId int
	BookId   int
}

type UpdateAuthorInput struct {
	Fullname  *string `json:"title"`
	Pseudonym *string `json:"description"`
}

func (i UpdateAuthorInput) Validate() error {
	if i.Fullname == nil && i.Pseudonym == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type UpdateBookInput struct {
	Title *string `json:"title"`
	Genre *string `json:"genre"`
}

func (i UpdateBookInput) Validate() error {
	if i.Title == nil && i.Genre == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type UpdateUserInput struct {
	Name *string `json:"name"`
}

func (i UpdateUserInput) Validate() error {
	if i.Name == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
