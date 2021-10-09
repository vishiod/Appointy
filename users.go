package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"instamongo/utils"
	"log"
	"net/http"
)

type User struct {
	ID       string `json:"instaHandle"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PaginatedUsers struct {
	//PageNumber int `json:"pageNumber"`
	//TotalPages int `json:"totalPages"`
	Users	 []User `json:"Users"`
}

var users = []User {
	{ID: "1", Name: "Ash",   Email: "vs@gmail.com", Password: "111"},
	{ID: "2", Name: "Misty", Email: "ys@gmail.com", Password: "111"},
	{ID: "3", Name: "Brock", Email: "ss@gmail.com", Password: "111"},
}

var paginatedUsers = [] PaginatedUsers {
	{
	//PageNumber: 1, TotalPages: 1,
		Users: users,
	},
}

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

	appDB := getDBStore().db
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

	appDB := getDBStore().db
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

	id := c.Param("id")

	appDB := getDBStore().db
	// Get a handle for your collection
	collection := appDB.Collection("instaPosts")

	filterCursor, err := collection.Find(c, bson.M{"UserID": id})
	if err != nil {
		log.Fatal(err)
		return
	}

	var postsFiltered []bson.M
	if err = filterCursor.All(c, &postsFiltered); err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(postsFiltered)
	fmt.Println("Found document (array): ", postsFiltered)

	c.IndentedJSON(http.StatusOK, postsFiltered)
}

func postUserByMongo(c *gin.Context){
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	//users = append(users, newUser)
	appDB := getDBStore().db
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

	var tempPassword = newUser.Password

	newUser.Password = checkSecrecy(tempPassword)

	if  utils.IsValidEmail(newUser.Email) {
		insertResult, err := collection.InsertOne(context.TODO(), newUser)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
		c.IndentedJSON(http.StatusCreated, newUser)
	}else{
		c.IndentedJSON(http.StatusBadRequest, "Invalid Email")
	}

}
