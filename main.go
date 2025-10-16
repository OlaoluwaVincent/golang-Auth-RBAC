package main

import (
	"go/auth/controllers"
	middlewares "go/auth/middleware"
	"go/auth/repositories"
	"go/auth/services"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	db, err := gorm.Open("sqlite3", "app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.AutoMigrate(&repositories.User{})

	userRepo := repositories.NewGormUserRepository(db)
	authSvc := services.NewAuthService(userRepo, os.Getenv("JWT_SECRET"), time.Minute*15, time.Hour*24*7)
	authCtrl := controllers.NewAuthController(authSvc)

	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/register", authCtrl.Register)
		api.POST("/login", authCtrl.Login)
		api.POST("/refresh", authCtrl.Refresh)
		protected := api.Group("/protected")
		protected.Use(middlewares.AuthMiddleware(authSvc))
		{
			protected.GET("/me", authCtrl.Me)
			admin := protected.Group("/admin")
			admin.Use(middlewares.RBACMiddleware("admin"))
			{
				admin.GET("/dashboard", authCtrl.AdminDashboard)
			}
		}
	}

	r.Run(":8080")
}
