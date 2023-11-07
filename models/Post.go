package model

type Post struct {
	ID            string
	Image         any
	Title         string
	Content       string
	UserID        string
	LikesCount    int
	DisLikesCount int
	CommentsCounts int
	CreatedAt     string
	Username      string
	FirstName     string
	LastName      string
	Avatar        string
	Comments      []Comment
	Categories    []Category
	likeStatus    string
}
