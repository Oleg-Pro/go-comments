package http

import (
	"cybersport-comments-go/internal/comment"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc comment.UseCase) {
	h := NewHandler(uc)

	bookmarks := router.Group("/comments")
	{
		bookmarks.POST("", h.Create)
		bookmarks.GET("", h.GetComments)
		bookmarks.DELETE("", h.Delete)
	}
}
