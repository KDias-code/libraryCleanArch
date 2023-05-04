package repository

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"

	todo "github.com/KDias-code/todoapp"
	"github.com/jmoiron/sqlx"
)

type TodoAuthorPostgres struct {
	db *sqlx.DB
}

func NewTodoAuthorPostgres(db *sqlx.DB) *TodoAuthorPostgres {
	return &TodoAuthorPostgres{db: db}
}

func (r *TodoAuthorPostgres) Create(userId int, author todo.TodoAuthor) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createAuthorQuery := fmt.Sprintf("INSERT INTO %s (fullname, pseudonym, spec) VALUES ($1, $2, $3) RETURNING id", todoAuthorsTable)
	row := tx.QueryRow(createAuthorQuery, author.Fullname, author.Pseudonym, author.Spec)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersAuthorQuery := fmt.Sprintf("INSERT INTO %s (user_id, author_id) VALUES ($1, $2)", usersAuthorsTable)
	_, err = tx.Exec(createUsersAuthorQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *TodoAuthorPostgres) GetAll(userId int) ([]todo.TodoAuthor, error) {
	var authors []todo.TodoAuthor

	query := fmt.Sprintf("SELECT ta.id, ta.fullname, ta.pseudonym, ta.spec FROM %s ta INNER JOIN %s ua on ta.id = ua.author_id WHERE ua.user_id = $1",
		todoAuthorsTable, usersAuthorsTable)
	err := r.db.Select(&authors, query, userId)

	return authors, err
}
func (r *TodoAuthorPostgres) GetById(userId, authorId int) (todo.TodoAuthor, error) {
	var author todo.TodoAuthor

	query := fmt.Sprintf(`SELECT ta.id, ta.fullname, ta.pseudonym, ta.spec FROM %s ta INNER JOIN %s ua on ta.id = ua.author_id WHERE ua.user_id = $1 AND ua.author_id = $2`,
		todoAuthorsTable, usersAuthorsTable)

	err := r.db.Get(&author, query, userId, authorId)

	return author, err
}

func (r *TodoAuthorPostgres) Delete(userId, authorId int) error {
	query := fmt.Sprintf("DELETE FROM %s ta USING %s ua WHERE ta.id = ua.author_id AND ua.user_id=$1 AND ua.author_id=$2",
		todoAuthorsTable, usersAuthorsTable)

	_, err := r.db.Exec(query, userId, authorId)

	return err
}

func (r *TodoAuthorPostgres) Update(userId, authorId int, input todo.UpdateAuthorInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	ardId := 1

	if input.Fullname != nil {
		setValues = append(setValues, fmt.Sprintf("fullname=$%d", ardId))
		args = append(args, *input.Fullname)
		ardId++
	}

	if input.Pseudonym != nil {
		setValues = append(setValues, fmt.Sprintf("pseudonym=$%d", ardId))
		args = append(args, *input.Pseudonym)
		ardId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s ta SET %s FROM %s ua WHERE ta.id = ua.author_id AND ua.author_id=$%d AND ua.user_id=$%d",
		todoAuthorsTable, setQuery, usersAuthorsTable, ardId, ardId+1)
	args = append(args, authorId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
