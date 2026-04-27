package user

import (
	"database/sql"
	"errors"
)

type UserRepository interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(email, passwordHash string) error
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) GetUserByEmail(email string) (*User, error) {
	user := User{}
	row := r.db.QueryRow("SELECT id, email, password_hash FROM users WHERE email = $1", email)
	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) CreateUser(email, passwordHash string) error {
	_, err := r.db.Exec("INSERT INTO users (email, password_hash) VALUES ($1, $2)", email, passwordHash)
	if err != nil {
		return err
	}
	return nil
}
