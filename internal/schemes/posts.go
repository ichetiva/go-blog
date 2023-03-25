package schemes

type ReqCreatePost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ReqDeletePost struct {
	PostID uint `json:"post_id"`
}
