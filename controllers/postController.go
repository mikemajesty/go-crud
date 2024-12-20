package controllers

import (
	"mike/goolang/go-crud/initializers"
	"mike/goolang/go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(controller *gin.Context) {

	var body struct {
		Title string
		Body  string
	}

	controller.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		controller.JSON(400, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	controller.JSON(200, result.RowsAffected)
}

func PostIndex(controller *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	controller.JSON(200, posts)

}

func PostShow(controller *gin.Context) {

	id := controller.Param("id")

	if id == "" {
		controller.JSON(400, gin.H{
			"message": "id is required",
		})
		return
	}

	var post models.Post
	postFound := initializers.DB.First(&post, id)

	if postFound.Error != nil {
		controller.JSON(400, gin.H{
			"message": "Post not found",
		})
		return
	}

	controller.JSON(200, post)
}

func PostUpdate(controller *gin.Context) {

	id := controller.Param("id")

	if id == "" {
		controller.JSON(400, gin.H{
			"message": "id is required",
		})
		return
	}

	var post models.Post
	postFound := initializers.DB.First(&post, id)

	if postFound.Error != nil {
		controller.JSON(400, gin.H{
			"message": "Post not found",
		})
		return
	}

	var body struct {
		Title string
		Body  string
	}

	controller.Bind(&body)

	post.Title = body.Title
	post.Body = body.Body

	result := initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	if result.Error != nil {
		controller.JSON(400, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	controller.JSON(200, post)
}

func PostDelete(controller *gin.Context) {

	id := controller.Param("id")

	if id == "" {
		controller.JSON(400, gin.H{
			"message": "id is required",
		})
		return
	}

	var post models.Post
	postFound := initializers.DB.First(&post, id)

	if postFound.Error != nil {
		controller.JSON(400, gin.H{
			"message": "Post not found",
		})
		return
	}

	result := initializers.DB.Delete(&post, id)

	if result.Error != nil {
		controller.JSON(400, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	controller.JSON(200, gin.H{
		"message": "Post deleted",
	})
}
