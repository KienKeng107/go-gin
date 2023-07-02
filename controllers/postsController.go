package controllers

import (
	"go-gin/initializers"
	"go-gin/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get Data of resquest body
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	// create a post
	post := models.Post{Title: body.Title, Boby: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// get post
	var posts []models.Post
	initializers.DB.Find(&posts)

	// response
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get Id of url
	id := c.Param("id")
	// get post
	var post models.Post
	initializers.DB.First(&post, id)

	// response
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get the id of the URL
	id := c.Param("id")

	// Get the data of req body
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	// Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Boby: body.Body})

	// response
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsDelete(c *gin.Context) {
	// Get the ID of the URL
	id := c.Param("id")

	// Delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// response
	c.Status(200)

}