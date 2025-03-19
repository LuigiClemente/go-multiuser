package handlers

import (
    "database/sql"
    "net/http"

    "github.com/gin-gonic/gin"
    "go-multiuser/models"
)

// RegisterRoutes registers user-related routes
func RegisterRoutes(router *gin.Engine, db *sql.DB) {
    router.GET("/users", func(c *gin.Context) { GetUsers(c, db) })
    router.GET("/users/:id", func(c *gin.Context) { GetUser(c, db) })
}

// GetUsers retrieves all users from the database
func GetUsers(c *gin.Context, db *sql.DB) {
    rows, err := db.Query("SELECT id, email, name FROM users")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
        return
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        if err := rows.Scan(&user.ID, &user.Email, &user.Name); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan user"})
            return
        }
        users = append(users, user)
    }

    c.JSON(http.StatusOK, users)
}

// GetUser retrieves a single user by ID
func GetUser(c *gin.Context, db *sql.DB) {
    id := c.Param("id")
    var user models.User
    if err := db.QueryRow("SELECT id, email, name FROM users WHERE id = $1", id).Scan(&user.ID, &user.Email, &user.Name); err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
        }
        return
    }

    c.JSON(http.StatusOK, user)
}