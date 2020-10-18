package repository

import (
	"fmt"

	"github.com/SolidShake/dividend-calendar-go/pkg/repository/models"
	"github.com/jmoiron/sqlx"
)

type AuthMysql struct {
	db *sqlx.DB
}

func NewAuthMysql(db *sqlx.DB) *AuthMysql {
	return &AuthMysql{db: db}
}

func (r *AuthMysql) CreateUser(user models.User) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (username, email, password_hash) VALUES (?, ?, ?)", userTable)
	result, err := r.db.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *AuthMysql) GetUser(email, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email = ? AND password_hash = ?", userTable)
	err := r.db.Get(&user, query, email, password)

	return user, err
}
