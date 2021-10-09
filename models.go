package main

type Post struct {
	UserID  string `json:"userId"`
	PostID  string `json:"postId"`
	Caption string `json:"caption"`
	URL     string `json:"url"`
	//TimeStamp time.Time `json:"TimePosted"`
	TimeStamp string `json:"TimePosted"`
}

type User struct {
	ID       string `json:"instaHandle"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
