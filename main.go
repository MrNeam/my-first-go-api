package main

import (
	"github.com/MrNeam/my-first-go-api/controllers"
	"github.com/MrNeam/my-first-go-api/initializers"
	"github.com/gin-gonic/gin"
)

func init(){
    initializers.LoadEnvVariables()
    initializers.ConnectToDB()
}

func main() {
    r := gin.Default()

    r.GET("/", controllers.PostsCreate)

    r.Run()
}