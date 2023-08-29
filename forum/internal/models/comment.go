// models/comment.go

package models

type Comment struct {
	ID        int `json:"id"`
	PostID    string `json:"postID"`
	Content   string `json:"content"`
	Username  string `json:"username"`
	Likes	   int 	 `json:"likes"`
	Dislikes	int 	 `json:"dislikes"`
	// Add more fields as needed
}
