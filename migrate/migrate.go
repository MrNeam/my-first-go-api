package main

import (
	"github.com/MrNeam/my-first-go-api/initializers"
	"github.com/MrNeam/my-first-go-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}