package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"instamongo/utils"
)


func main() {
	fmt.Println("Vishal")

	router := gin.Default()
	router.GET("/user", getUsersMongo)
	router.GET("/users/:id", getUserByIDMongo)
	router.GET("/post", getPostsMongo)
	router.GET("/posts/:id", getPostByIDMongo)
	router.GET("/posts/users/:id", getPostsOfAParticularUserByMongo)

	router.POST("/users", postUserByMongo)
	router.POST("/posts", postAnInstaPostByMongo)

	//println(utils.IsValidEmail("yogeshsharma@locus.sh"))
	//println(utils.IsValidEmail("yogesh-sharma@locus.sh"))

	router.Run("localhost:8080")
}

func checkSecrecy(pass string) string {

	//preBuiltString := "His money is twice tainted: 'taint yours and 'taint mine."
	preBuilt := utils.ComputeHmac256(pass)

	halfAString := "His money is twice tainted:"
	adJoinedString1 := halfAString + " 'taint yours and 'taint mine."
	adJoinedString2 := halfAString + " 'taint yours and 'taint mine. "

	adJoinedToBuild := utils.ComputeHmac256(adJoinedString1)
	fmt.Println("Are we equal?:", utils.IsHMACEqual(preBuilt, adJoinedToBuild),
		"<------>" + "then use my code to save in DB:", utils.HMACToString(preBuilt))

	fmt.Println("Is with space equal?:", utils.IsHMACEqual(preBuilt, utils.ComputeHmac256(adJoinedString2)))

	return  utils.HMACToString(preBuilt)
}

