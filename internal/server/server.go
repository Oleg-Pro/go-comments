package server

import (
	"context"
	"cybersport-comments-go/internal/comment"
	commenthttp "cybersport-comments-go/internal/comment/delivery/http"
	commentsql "cybersport-comments-go/internal/comment/repository/sql"
	commentusecase "cybersport-comments-go/internal/comment/usecase"
	"cybersport-comments-go/internal/config"
	"cybersport-comments-go/internal/dbconnection"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	httpServer     *http.Server
	commentUseCase comment.UseCase
}

func NewApp() *App {

	db, err := dbconnection.CreateDBConnection(config.Conf.DatabaseUrl)
	if err != nil {
		log.Println("Incorrect DB Url:" + config.Conf.DatabaseUrl)
		log.Fatalln(err)
	}

	commentRepository := commentsql.NewCommentRepository(db)
	commentUseCase := commentusecase.NewCommentUseCase(commentRepository)

	return &App{
		commentUseCase: commentUseCase,
	}
}

func (a *App) Run(port string) error {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.New() //gin.Default()

	router.Use(gin.Logger())
	api := router.Group("/v1/api")
	commenthttp.RegisterHTTPEndpoints(api, a.commentUseCase)

	log.Println("PORT::::" + port)
	a.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}
