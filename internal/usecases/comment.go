package usecases

import (
	"context"
	"cybersport-comments-go/internal/domain/models"
	"cybersport-comments-go/internal/domain/repositories"
)

type CommentUseCaseHandler interface {
	AddComment(ctx context.Context, comment *models.Comment) error
	GetComments(ctx context.Context) ([]*models.Comment, error)
	DeleteComment(ctx context.Context, id uint64) error
}

type CommentUseCase struct {
	commentRepository repositories.CommentRepository
}

func NewCommentUseCase(commentRepository repositories.CommentRepository) *CommentUseCase {
	return &CommentUseCase{
		commentRepository: commentRepository,
	}
}

//Implement UseCase interface
func (commentUseCase CommentUseCase) AddComment(ctx context.Context, comment *models.Comment) error {
	return commentUseCase.commentRepository.AddComment(ctx, comment)
}

func (commentUseCase CommentUseCase) GetComments(ctx context.Context) ([]*models.Comment, error) {
	return commentUseCase.commentRepository.GetComments(ctx)
}

func (commentUseCase CommentUseCase) DeleteComment(ctx context.Context, id uint64) error {
	return commentUseCase.commentRepository.DeleteComment(ctx, id)
}
