package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Post struct {
	UserID   string `json:"userId"`
	PostID	 string `json:"postId"`
	Caption	 string `json:"caption"`
	URL 	 string `json:"url"`
	TimeStamp time.Time `json:"TimePosted"`
}

type PaginatedPosts struct {
	PageNumber int `json:"pageNumber"`
	TotalPages int `json:"totalPages"`
	Posts	[]Post `json:"Posts"`
}

var posts = []Post{
	{ PostID: "1", UserID: "1", Caption: "Hello Brother", URL: "Thor", TimeStamp: time.Now()},
	{ PostID: "2", UserID: "1", Caption: "Hello Mother", URL: "Brother", TimeStamp: time.Now()},
	{ PostID: "3", UserID: "2", Caption: "Hello Sister", URL: "Hello", TimeStamp: time.Now()},
	{ PostID: "4", UserID: "3", Caption: "Hello Myself", URL: "Bruh", TimeStamp: time.Now()},
}

var paginatedPosts = [] PaginatedPosts {
	{PageNumber: 1, TotalPages: 1, Posts: posts},
}

func getPosts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, paginatedPosts)
}

func getPostByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, paginatedPost := range paginatedPosts {
		for _, aPost := range paginatedPost.Posts {
			if aPost.PostID == id {
				c.IndentedJSON(http.StatusOK, aPost)
				return
			}
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Post not found"})
}
