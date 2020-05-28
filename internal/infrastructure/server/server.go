package server

import (
	"context"
	"cybersport-comments-go/internal/infrastructure/config"
	"cybersport-comments-go/internal/infrastructure/dbconnection"
	"cybersport-comments-go/internal/interfaces/controllers"
	sqlrepostirories "cybersport-comments-go/internal/interfaces/repositories/sql"
	"cybersport-comments-go/internal/usecases"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	httpServer     *http.Server
	commentUseCase *usecases.CommentUseCase
}

func NewApp() *App {

	db, err := dbconnection.CreateDBConnection(config.Conf.DatabaseUrl)
	if err != nil {
		log.Println("Incorrect DB Url:" + config.Conf.DatabaseUrl)
		log.Fatalln(err)
	}

	commentRepository := sqlrepostirories.NewCommentRepository(db)
	commentUseCase := usecases.NewCommentUseCase(commentRepository)

	return &App{
		commentUseCase: commentUseCase,
	}
}

func (app *App) Run(port string) error {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(gin.Logger())
	api := router.Group("/v1/api")

	commentController := controllers.NewCommentController(app.commentUseCase)
	commentController.RegisterHTTPEndpoints(api)

	log.Println("PORT::::" + port)
	app.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		if err := app.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return app.httpServer.Shutdown(ctx)
}
