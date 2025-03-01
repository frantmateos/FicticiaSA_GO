package main

import (
	repo "Golang/clients"
	controller "Golang/controller"
	"Golang/middleware"
	service "Golang/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetUserByName(c *gin.Context)
	UsuarioInsert(c *gin.Context)
	UpdateUser(c *gin.Context)
	Extrac(c *gin.Context)
	GetUserById(c *gin.Context)
	Login(c *gin.Context)
}

func main() {

	sqlconfig := repo.MySQLConfig{
		Name: "users",
		User: "root",
		Pass: "Tomas1927",
		Host: "localhost",
	}

	mainRepo := repo.NewSql(sqlconfig)
	Service := service.NewService(mainRepo)
	Controller := controller.NewController(Service)
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type, X-Auth-Token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})
	router.POST("/users", Controller.UsuarioInsert)
	router.POST("/users/login", Controller.Login)
	router.GET("/users/token", Controller.Extrac)

	router.GET("/users/all", middleware.AuthMiddleware(), Controller.GetAllUsers)
	router.GET("/users", middleware.AuthMiddleware(), Controller.GetUserByName)
	router.GET("/users/:id", middleware.AuthMiddleware(), Controller.GetUserById)
	router.PUT("/users", middleware.AuthMiddleware(), Controller.UpdateUser)

	router.Run(":8080")

}
