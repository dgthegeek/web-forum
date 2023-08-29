// models/post.go

package models

type Post struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Author    string   `json:"author"`
	LikedBy   []string `json:"likedBy"`
	Categories []string   `json:"categories"`
	Likes	   int 	 `json:"likes"`
	Dislikes	int 	 `json:"dislikes"`
	
	// Add more fields as needed
}
