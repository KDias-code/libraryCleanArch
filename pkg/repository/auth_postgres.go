package repository

import (
	"errors"
	"fmt"
	todo "github.com/KDias-code/todoapp"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (todo.User, error) {
	var user todo.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", userTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}

func (r *AuthPostgres) GetAll(userId int) ([]todo.User, error) {
	var members []todo.User

	query := fmt.Sprintf(`
		SELECT
			u.id,
			u.name,
			tb.title,
			tb.genre,
			tb.isbn
		FROM
			%s ua
			INNER JOIN %s ub ON ua.user_id = ub.user_id
			INNER JOIN %s tb ON tb.id = ub.book_id
			INNER JOIN %s ta ON ta.id = tb.author_id
			INNER JOIN %s u ON u.id = ua.user_id
		WHERE
			ua.user_id = $1
	`, usersAuthorsTable, usersBooksTable, todoBooksTable, todoAuthorsTable, userTable)

	err := r.db.Select(&members, query, userId)

	return members, err
}

func (r *AuthPostgres) Delete(userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", userTable)
	_, err := r.db.Exec(query, userId)
	return err
}

func (r *AuthPostgres) Update(userId int, input todo.UpdateUserInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if len(setValues) == 0 {
		return errors.New("no fields to update")
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d",
		userTable, setQuery, argId)
	args = append(args, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %v", args)

	_, err := r.db.Exec(query, args...)
	return err
}
