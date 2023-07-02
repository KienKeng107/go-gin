package main

import (
	"go-gin/controllers"
	"go-gin/initializers"
	"go-gin/middleware"

	"github.com/gin-gonic/gin"
)
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()

	r.POST("/posts", middleware.RequireAuth, controllers.PostsCreate)
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.PUT("/posts/:id", middleware.RequireAuth,controllers.PostsUpdate)
	r.GET("/posts", middleware.RequireAuth, controllers.PostsIndex)
	r.GET("/posts/:id", middleware.RequireAuth, controllers.PostsShow)
	r.DELETE("/posts/:id", middleware.RequireAuth,controllers.PostsDelete)

	r.Run("localhost:3000") // listen and serve on 0.0.0.0:3000
}