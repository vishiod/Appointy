package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"instamongo/dbutils"
	"instamongo/utils"
	"log"
	"net/http"
)

//func getUsers(c *gin.Context) {
//	c.IndentedJSON(http.StatusOK, users)
//}
//
//func getUserByID(c *gin.Context) {
//	id := c.Param("id")
//
//	for _, page := range paginatedUsers {
//		for _, user := range page.Users {
//			if user.ID == id{
//				c.IndentedJSON(http.StatusOK, user)
//			}
//		}
//	}
//}
//
//func getPostsOfAParticularUser(c *gin.Context) {
//
//	id := c.Param("id")
//	var accumulatedPostsOfAParticularUser []Post
//	var relevantPostsPresent = false
//
//	for _, aPaginatedPost := range paginatedPosts {
//		for _, aPost := range aPaginatedPost.Posts {
//			if aPost.UserID == id {
//				accumulatedPostsOfAParticularUser = append(accumulatedPostsOfAParticularUser, aPost)
//				relevantPostsPresent = true
//			}
//		}
//	}
//	if !relevantPostsPresent {
//		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Post not found"})
//	} else {
//		c.IndentedJSON(http.StatusOK, accumulatedPostsOfAParticularUser)
//	}
//}
//
//func postUser(c *gin.Context) {
//	var newUser User
//
//	if err := c.BindJSON(&newUser); err != nil {
//		return
//	}
//
//	users = append(users, newUser)
//	c.IndentedJSON(http.StatusCreated, newUser)
//}

func getUsersMongo(c *gin.Context){
	var  mongoUsers []*User

	appDB := dbutils.GetDBStore().DB
	// Get a handle for your collection
	collection := appDB.Collection("users")

	findOptions := options.Find()

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through the cursor
	for cur.Next(context.TODO()) {
		var elem User
		err := cur.Decode(&elem)


		if err != nil {
			log.Fatal(err)
		}

		mongoUsers = append(mongoUsers, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	fmt.Println("Found multiple documents (array): ", mongoUsers)

	cur.Close(context.TODO())

	c.IndentedJSON(http.StatusOK, &mongoUsers)
}

func getUserByIDMongo(c *gin.Context)  {
	id := c.Param("id")

	appDB := dbutils.GetDBStore().DB
	// Get a handle for your collection
	collection := appDB.Collection("users")

	filterCursor, err := collection.Find(c, bson.M{"instaHandle": id})
	if err != nil {
		log.Fatal(err)
	}

	var usersFiltered []bson.M
	if err = filterCursor.All(c, &usersFiltered); err != nil {
		log.Fatal(err)
	}
	fmt.Println(usersFiltered)
	fmt.Println("Found document (array): ", usersFiltered)

	c.IndentedJSON(http.StatusOK, usersFiltered)

}

func getPostsOfAParticularUserByMongo(c *gin.Context) {

	var  mongoPosts []*Post
	id := c.Param("id")

	appDB := dbutils.GetDBStore().DB
	// Get a handle for your collection
	collection := appDB.Collection("instaPosts")

	filterCursor, err := collection.Find(c, bson.M{"UserID": id})
	if err != nil {
		log.Fatal(err)
		return
	}

	for filterCursor.Next(context.TODO()) {
		var elem Post
		err := filterCursor.Decode(&elem)


		if err != nil {
			log.Fatal(err)
		}

		mongoPosts = append(mongoPosts, &elem)
	}

	fmt.Println("Found document (array): ", mongoPosts)

	c.IndentedJSON(http.StatusOK, mongoPosts)
}

func saveUserByMongo(c *gin.Context){

	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	//users = append(users, newUser)
	appDB := dbutils.GetDBStore().DB
	// Get a handle for your collection
	collection := appDB.Collection("users")

	filterCursor, err := collection.Find(c, bson.M{"email": newUser.Email})

	if filterCursor.RemainingBatchLength() != 0{
		c.IndentedJSON(http.StatusBadRequest, "Insta Handle Already Exists")
		return
	}
	if filterCursor.RemainingBatchLength() != 0{
		c.IndentedJSON(http.StatusBadRequest, "Email Already Exists")
		return
	}

	filterCursor, err = collection.Find(c, bson.M{"instaHandle": newUser.ID})

	if err != nil {
		log.Fatal(err)
		return
	}

	newUser.Password = utils.HMACToString(utils.ComputeHmac256(newUser.Password))

	if  utils.IsValidEmail(newUser.Email) {
		insertResult, err := collection.InsertOne(context.TODO(), newUser)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
		c.IndentedJSON(http.StatusCreated, newUser)
	} else {
		c.IndentedJSON(http.StatusBadRequest, "Invalid Email")
	}

}
