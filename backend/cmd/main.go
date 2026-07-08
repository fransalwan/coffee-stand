package main

import (
	"log"
	"os"

	"backend/internal/database"
	"backend/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	_ = godotenv.Load()

	// Connect DB & Auto-migrate
	database.ConnectDatabase()

	// Setup Gin Router
	r := gin.Default()

	// CORS Middleware (WAJIB biar frontend Vue lu bisa akses API ini nanti)
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Public Routes (Customer View)
	api := r.Group("/api")
	{
		api.GET("/menu", handlers.GetAllMenus)
	}

	// Admin/Owner Routes (Nanti kita kunci pakai JWT Middleware di sini)
	admin := api.Group("/admin")
	{
		admin.GET("/menu", handlers.GetAllMenusAdmin)
		admin.POST("/menu", handlers.CreateMenu)
		admin.PUT("/menu/:id", handlers.UpdateMenu)
		admin.DELETE("/menu/:id", handlers.DeleteMenu)
	}

	// Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server starting on http://localhost:%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
