package comment

import (
	"context"
	"cybersport-comments-go/internal/models"
)

type Repository interface {
	AddComment(ctx context.Context, comment *models.Comment) error
	GetComments(ctx context.Context) ([]*models.Comment, error)
	DeleteComment(ctx context.Context, id uint64) error
}
