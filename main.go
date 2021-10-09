package main

import (
	"github.com/gin-gonic/gin"
)


func main() {

	router := gin.Default()

	// User related APIs
	router.GET("/user", getUsersMongo)
	router.GET("/users/:id", getUserByIDMongo)

	router.POST("/users", saveUserByMongo)

	// Posts related APIs
	router.GET("/post", getPostsMongo)
	router.GET("/posts/:id", getPostByIDMongo)
	router.GET("/posts/users/:id", getPostsOfAParticularUserByMongo)

	router.POST("/posts", saveAnInstaPostByMongo)

	err := router.Run("localhost:8080")

	if err != nil {
		println("Could not start the server")
		return 
	}
}
