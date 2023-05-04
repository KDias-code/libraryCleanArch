package repository

import (
	"fmt"
	_ "github.com/KDias-code/todoapp"
	todo "github.com/KDias-code/todoapp"
	"github.com/jmoiron/sqlx"
	"strings"
)

type TodoBookPostgres struct {
	db *sqlx.DB
}

func NewTodoBookPostgres(db *sqlx.DB) *TodoBookPostgres {
	return &TodoBookPostgres{db: db}
}

func (r *TodoBookPostgres) Create(authorId int, book todo.TodoBook) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var bookId int
	createBookQuery := fmt.Sprintf("INSERT INTO %s (title, genre, isbn) values ($1, $2, $3) RETURNING id", todoBooksTable)

	row := tx.QueryRow(createBookQuery, book.Title, book.Genre, book.Isbn)
	err = row.Scan(&bookId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createAuthorBooksQuery := fmt.Sprintf("INSERT INTO %s (author_id, book_id) values ($1, $2)", authorsBooksTable)
	_, err = tx.Exec(createAuthorBooksQuery, authorId, bookId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return bookId, tx.Commit()
}
func (r *TodoBookPostgres) GetAll(userId, authorId int) ([]todo.TodoBook, error) {
	var books []todo.TodoBook
	query := fmt.Sprintf(`SELECT tb.id, tb.title, tb.genre, tb.isbn FROM %s tb INNER JOIN %s ab on ab.book_id = tb.id
									INNER JOIN %s ua on ua.author_id = ab.author_id WHERE ab.author_id = $1 AND ua.user_id = $2`,
		todoBooksTable, authorsBooksTable, usersAuthorsTable)
	if err := r.db.Select(&books, query, authorId, userId); err != nil {
		return nil, err
	}

	return books, nil
}

func (r *TodoBookPostgres) GetById(userId, bookId int) (todo.TodoBook, error) {
	var book todo.TodoBook
	query := fmt.Sprintf(`SELECT tb.id, tb.title, tb.genre, tb.isbn FROM %s tb INNER JOIN %s ab on ab.book_id = tb.id
									INNER JOIN %s ua on ua.author_id = lb.author_id WHERE tb.id = $1 AND ua.user_id = $2`,
		todoBooksTable, authorsBooksTable, usersAuthorsTable)
	if err := r.db.Get(&book, query, bookId, userId); err != nil {
		return book, err
	}

	return book, nil
}

func (r *TodoBookPostgres) Delete(userId, bookId int) error {
	query := fmt.Sprintf(`DELETE FROM %s tb USING %s ab, %s ua
								WHERE tb.id = ab.book_id AND ab.author_id = ua.author_id AND ua.user_id = $1 AND tb.id = $2`,
		todoBooksTable, authorsBooksTable, usersAuthorsTable)

	_, err := r.db.Exec(query, userId, bookId)
	return err
}
func (r *TodoBookPostgres) Update(userId, bookId int, input todo.UpdateBookInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	ardId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", ardId))
		args = append(args, *input.Title)
		ardId++
	}

	if input.Genre != nil {
		setValues = append(setValues, fmt.Sprintf("genre=$%d", ardId))
		args = append(args, *input.Genre)
		ardId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf(`UPDATE %s tb SET %s FROM %s ab, %s ua 
                    			WHERE tb.id = lb.book_id AND lb.author_id = ua.author_id AND ua.user_id = $%d AND tb.id = $%d`,
		todoBooksTable, setQuery, authorsBooksTable, usersAuthorsTable, ardId, ardId+1)
	args = append(args, userId, bookId)

	_, err := r.db.Exec(query, args...)
	return err
}
