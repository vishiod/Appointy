package main

import (
	"fmt"
	formatterUtils "instamongo/utils"

	"github.com/gin-gonic/gin"
)

type Trainer struct {
	Name string
	Age  int
	City string
}


func main() {
	fmt.Println("Vishal")

	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
	router.GET("/posts", getPosts)
	router.GET("/posts/:id", getPostByID)
	router.GET("/posts/users/:id", getPostsOfAParticularUser)

	formatterUtils.IsValidEmail("yogeshsharma@locus.sh")
	router.Run("localhost:8080")
}
