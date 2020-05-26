package usecase

import (
	"context"
	"cybersport-comments-go/internal/comment"
	"cybersport-comments-go/internal/models"
)

type CommentUseCase struct {
	commentRepository comment.Repository
}

func NewCommentUseCase(commentRepository comment.Repository) *CommentUseCase {
	return &CommentUseCase{
		commentRepository: commentRepository,
	}
}

func (commentUseCase CommentUseCase) AddComment(ctx context.Context, comment *models.Comment) error {
	return commentUseCase.commentRepository.AddComment(ctx, comment)
}

func (commentUseCase CommentUseCase) GetComments(ctx context.Context) ([]*models.Comment, error) {
	return commentUseCase.commentRepository.GetComments(ctx)
}

func (commentUseCase CommentUseCase) DeleteComment(ctx context.Context, id uint64) error {
	return commentUseCase.commentRepository.DeleteComment(ctx, id)
}
