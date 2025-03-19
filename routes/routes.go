package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"go-multiuser/handlers"
	"go-multiuser/middleware"
)

// RegisterRoutes registers all API routes
func RegisterRoutes(router *gin.Engine, db *sql.DB) {
	api := router.Group("/api")

	// User routes
	userRoutes := api.Group("/users")
	userRoutes.Use(middleware.AuthMiddleware())
	userRoutes.GET("/", func(c *gin.Context) {
		handlers.GetUsers(c, db)
	})
	userRoutes.GET("/:id", func(c *gin.Context) {
		handlers.GetUser(c, db)
	})
}
