package store

import (
	"context"
	"database/sql"
)

type UsersStore struct {
	db *sql.DB
}

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"` /// not returning password in json response
	CreatedAt string `json:"created_at"`
}

func (s *UsersStore) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (username, password, email)
		VALUES ($1, $2, $3) RETURNING id, created_at
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Password,
		user.Email,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	) // .Scan(dest any...): scan columns in the matched row and assign to values pointed at dest

	if err != nil {
		return err
	}

	return nil
}
