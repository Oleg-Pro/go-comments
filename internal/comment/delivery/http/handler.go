package http

import (
	"cybersport-comments-go/internal/comment"
	"cybersport-comments-go/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Comment struct {
	Id       uint64  `json:"id"`
	Comment  string  `json:"comment"`
	ParentId *uint64 `json:"parent_id"`
	//Type uint8
	//ObjectId uint
	//Status uint8
	//UserId uint
}

type Handler struct {
	useCase comment.UseCase
}

func NewHandler(useCase comment.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type createInput struct {
	Id       uint64  `json:"id"`
	Comment  string  `json:"comment" binding:"required"`
	ParentId *uint64 `json:"parent_id"`
}

func (h *Handler) Create(c *gin.Context) {
	inp := new(createInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment := &models.Comment{
		Id:       inp.Id,
		Comment:  inp.Comment,
		ParentId: inp.ParentId,
	}

	err := h.useCase.AddComment(c.Request.Context(), comment)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newComment := toComment(comment)
	c.JSON(http.StatusOK, gin.H{"result": newComment})
}

type getResponse struct {
	Comments []*Comment `json:"comments"`
}

func (h *Handler) GetComments(c *gin.Context) {
	comments, err := h.useCase.GetComments(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &getResponse{
		Comments: toComments(comments),
	})
}

type deleteInput struct {
	ID uint64 `json:"id" binding:"required"`
}

func (h *Handler) Delete(c *gin.Context) {
	inp := new(deleteInput)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err := h.useCase.DeleteComment(c.Request.Context(), inp.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "ok"})
}

func toComments(comments []*models.Comment) []*Comment {
	out := make([]*Comment, len(comments))

	for i, comment := range comments {
		out[i] = toComment(comment)
	}

	return out
}

func toComment(comment *models.Comment) *Comment {
	return &Comment{
		Id:       comment.Id,
		Comment:  comment.Comment,
		ParentId: comment.ParentId,
	}
}
