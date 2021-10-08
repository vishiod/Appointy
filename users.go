package main

import (
	"github.com/gin-gonic/gin"
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

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func getUserByID(c *gin.Context) {
	id := c.Param("id")

	for _, page := range paginatedUsers {
		for _, user := range page.Users {
			if user.ID == id{
				c.IndentedJSON(http.StatusOK, user)
			}
		}
	}
}

func getPostsOfAParticularUser(c *gin.Context) {

	id := c.Param("id")
	var accumulatedPostsOfAParticularUser []Post
	var relevantPostsPresent = false

	for _, aPaginatedPost := range paginatedPosts {
		for _, aPost := range aPaginatedPost.Posts {
			if aPost.UserID == id {
				accumulatedPostsOfAParticularUser = append(accumulatedPostsOfAParticularUser, aPost)
				relevantPostsPresent = true
			}
		}
	}
	if !relevantPostsPresent {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Post not found"})
	} else {
		c.IndentedJSON(http.StatusOK, accumulatedPostsOfAParticularUser)
	}
}

func postUser(c *gin.Context) {
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}
