package main

import (
	"mike/goolang/go-crud/initializers"
	"mike/goolang/go-crud/models"
)

func init() {
	initializers.LoadEnv()
	initializers.LoadDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
