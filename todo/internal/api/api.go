package api

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"todo/pkg/config"
	"todo/pkg/data"
	"todo/pkg/services"
)

type App struct {
	server  *gin.Engine
	userSvc services.IUserService
	todoSvc services.ITodoService
	cfg     *config.Settings
}

func New(cfg *config.Settings, client *mongo.Client) *App {
	server := gin.Default()

	userProvider := data.NewUserProvider(cfg, client)

	userSvc := services.NewUserService(cfg, userProvider)

	todoProvider := data.NewTodoProvider(cfg, client)

	todoSvc := services.NewTodoService(cfg, todoProvider)

	return &App{
		server:  server,
		userSvc: userSvc,
		todoSvc: todoSvc,
		cfg:     cfg,
	}
}

func (a App) ConfigureRoutes() {
	a.server.GET("/v1/public/healthy", a.HealthCheck)
	a.server.POST("/v1/public/register", a.Register)
	a.server.POST("/v1/public/login", a.Login)
	a.server.POST("/v1/public/todos", a.CreateTodo)
	a.server.GET("/v1/public/todos", a.GetAllTodos)
	a.server.POST("/v1/public/delete_todos", a.DeleteTodos)

	protected := a.server.Group("v1/api")

	middleware := Middleware{config: a.cfg}

	protected.Use(middleware.Auth())

	protected.GET("/secret", func(c *gin.Context) {
		userId := c.GetString("user")
		c.String(200, userId)
	})
}

func (a App) Start() {
	a.ConfigureRoutes()
	a.server.Run("localhost:8080")
}
