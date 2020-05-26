package sql

import (
	"context"
	"cybersport-comments-go/internal/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type CommentRepository struct {
	db *sqlx.DB
}

const table = "comments"

func NewCommentRepository(db *sqlx.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (r CommentRepository) AddComment(ctx context.Context, comment *models.Comment) error {
	query := fmt.Sprintf(`INSERT INTO %s (comment, parent_id) VALUES ($1, $2) RETURNING id`,
		table)
	err := r.db.QueryRow(query, comment.Comment /*parentId*/, comment.ParentId).Scan(&comment.Id)
	return err
}

func (r CommentRepository) GetComments(ctx context.Context) ([]*models.Comment, error) {
	var comments = []*models.Comment{}

	query := fmt.Sprintf(`SELECT id, comment, parent_id
                                 FROM %s
                                 ORDER BY id`,
		table,
	)

	err := r.db.Select(&comments, query)
	return comments, err
}

func (r CommentRepository) DeleteComment(ctx context.Context, id uint64) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`,
		table)

	_, err := r.db.Exec(query, id)
	return err
}
