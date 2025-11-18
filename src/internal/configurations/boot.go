package configurations

import (
	repositories "buffalos/src/internal/repositories"
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	DB         *sql.DB
	Router     *gin.Engine
	Repos      Repositories
	Context    context.Context
	CancelFunc context.CancelFunc
}

type Repositories struct {
	Users *repositories.UserPG
}

type Services struct {
}

func Boot() *App {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	ctx, cancel := context.WithCancel(context.Background())

	db := SetupDB()
	repos := setupRepos(db)
	services := setupServices(repos)
	router := setupServer(services)

	return &App{
		DB:         db,
		Router:     router,
		Repos:      repos,
		Context:    ctx,
		CancelFunc: cancel,
	}
}

func setupRepos(db *sql.DB) Repositories {
	return Repositories{
		Users: repositories.NewUserPG(db),
	}
}

func setupServices(repos Repositories) *Services {
	return &Services{}
}

func setupServer(service *Services) *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
