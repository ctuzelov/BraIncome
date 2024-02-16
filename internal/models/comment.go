package models

type Comment struct {
	ID           int        `json:"id"`
	PostID       int        `json:"post_id"`
	UserID       int        `json:"user_id"`
	Content      string     `json:"content"`
	ParentID     int        `json:"parent_id"`
	Creator      string     `json:"creator"`
	Replies      []*Comment `json:"replies"`
	Likes        Reaction   `json:"likes"`
	Dislikes     Reaction   `json:"dislikes"`
	UserReaction string     `json:"user_reaction"`
}

type Reaction struct {
	Count int      `json:"count"`
	Users []string `json:"users"`
}
