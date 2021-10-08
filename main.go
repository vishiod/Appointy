package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"instamongo/utils"
)

type Trainer struct {
	Name string
	Age  int
	City string
}


func main() {
	fmt.Println("Vishal")

	router := gin.Default()
	router.GET("/user", getUsers)
	router.GET("/users/:id", getUserByID)
	router.GET("/post", getPosts)
	router.GET("/posts/:id", getPostByID)
	router.GET("/posts/users/:id", getPostsOfAParticularUser)

	router.POST("/users", postUser)
	router.POST("/posts", postAnInstaPost)

	println(utils.IsValidEmail("yogeshsharma@locus.sh"))
	println(utils.IsValidEmail("yogesh-sharma@locus.sh"))
	checkSecrecy()
	router.Run("localhost:8080")
}

func checkSecrecy() {

	preBuiltString := "His money is twice tainted: 'taint yours and 'taint mine."
	preBuilt := utils.ComputeHmac256(preBuiltString)

	halfAString := "His money is twice tainted:"
	adJoinedString1 := halfAString + " 'taint yours and 'taint mine."
	adJoinedString2 := halfAString + " 'taint yours and 'taint mine. "

	adJoinedToBuild := utils.ComputeHmac256(adJoinedString1)
	println("Are we equal?:", utils.IsHMACEqual(preBuilt, adJoinedToBuild),
		"<------>" + "then use my code to save in DB:", utils.HMACToString(preBuilt))

	println("Is with space equal?:", utils.IsHMACEqual(preBuilt, utils.ComputeHmac256(adJoinedString2)))

}
