package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"instamongo/dbutils"
	"log"
	"net/http"
	"time"
)

//func getPosts(c *gin.Context) {
//	c.IndentedJSON(http.StatusOK, paginatedPosts)
//}
//
//func getPostByID(c *gin.Context) {
//	id := c.Param("id")
//
//	// Loop over the list of albums, looking for
//	// an album whose ID value matches the parameter.
//	for _, paginatedPost := range paginatedPosts {
//		for _, aPost := range paginatedPost.Posts {
//			if aPost.PostID == id {
//				c.IndentedJSON(http.StatusOK, aPost)
//				return
//			}
//		}
//	}
//	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Post not found"})
//}
//
//func postAnInstaPost(c *gin.Context) {
//	var newInstaPost Post
//
//	if err := c.BindJSON(&newInstaPost); err != nil {
//		return
//	}
//
//	posts = append(posts, newInstaPost)
//	c.IndentedJSON(http.StatusCreated, newInstaPost)
//}

func getPostsMongo(c *gin.Context){
	var  mongoPosts []*Post

	appDB := dbutils.GetDBStore().DB

	// Get a handle for your collection
	collection := appDB.Collection("instaPosts")

	findOptions := options.Find()

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through the cursor
	for cur.Next(context.TODO()) {
		var elem Post
		err := cur.Decode(&elem)


		if err != nil {
			log.Fatal(err)
		}

		mongoPosts = append(mongoPosts, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	fmt.Println("Found multiple documents (array): ", mongoPosts)

	cur.Close(context.TODO())

	c.IndentedJSON(http.StatusOK, &mongoPosts)

	if err != nil {
		panic(err)
	}
}

func getPostByIDMongo(c *gin.Context)  {
	id := c.Param("id")

	appDB := dbutils.GetDBStore().DB
	// Get a handle for your collection
	collection := appDB.Collection("instaPosts")


	filterCursor, err := collection.Find(c, bson.M{"PostID": id})
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

func saveAnInstaPostByMongo(c *gin.Context){
	var newPost Post

	if err := c.BindJSON(&newPost); err != nil {
		return
	}
	newPost.TimeStamp = time.Now().String()
	appDB := dbutils.GetDBStore().DB
	// Get a handle for your collection
	collection := appDB.Collection("instaPosts")

	filterCursor, err := collection.Find(c, bson.M{"postid": newPost.PostID})

	if filterCursor.RemainingBatchLength() != 0{
		c.IndentedJSON(http.StatusBadRequest, "PostID Already Exists")
		return
	}

	insertResult, err := collection.InsertOne(context.TODO(), newPost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	c.IndentedJSON(http.StatusCreated, newPost)
}

