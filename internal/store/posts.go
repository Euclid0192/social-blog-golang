package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type PostsStore struct {
	db *sql.DB
}

// / Data model
// Remeber to add json marshalling encoder
type Post struct {
	ID        int64    `json:"id"`
	Content   string   `json:"content"`
	Title     string   `json:"title"`
	UserID    int64    `json:"user_id"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

// Create a post
func (s *PostsStore) Create(ctx context.Context, post *Post) error {
	query := `
		INSERT INTO posts (content, title, user_id, tags)
		VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at	
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Content,
		post.Title,
		post.UserID,
		pq.Array(post.Tags),
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	) // .Scan(dest any...): scan columns in the matched row and assign to values pointed at dest

	if err != nil {
		return err
	}

	return nil
}
