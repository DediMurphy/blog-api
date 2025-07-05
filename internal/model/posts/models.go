package posts

import (
	"time"
)

// untuk client / frontend
type (
	CreatePostRequest struct {
		PostTitle string `json:"postTitle"`
		PostContent string `json:"postContent"`
		PostHashtags []string `json:"postHashtags"`
	}
)

// untuk model database

type (
	PostModel struct {
		ID int64 `db:"id"`
		UserID int64 `db:"user_id"`
		PostTitle string `db:"post_title"`
		PostContent string `db:"post_content"`
		PostHashtags string `db:"post_hashtags"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt  time.Time `db:"updated_at"`
		CreatedBy string `db:"created_by"`
		UpdatedBy  string `db:"update_by"`
	}
)

type (
	GetAllPostResponse struct {
		Data []Post `json:"data"`
		Pagination Pagination `json:"pagination"`
	}

	Post struct {
		ID 	int64 `json:"id"`
		UserID int64 `json:"UserID"`
		Username string 	`json:"username"`
		PostTitle string `json:"postTitle"`
		PostContent string `json:"postContent"`
		PostHashtags []string `json:"postHashtags"`
		IsLiked bool `json:"isLiked"`
	}

	Pagination struct {
		Limit int	`json:"limit"`
		Offset int `json:"offset"`
	}

	GetPostResponse struct {
		PostDetail	Post `json:"postDetail"`
		LikeCount	int	`json:"likeCount"`
		Comments	[]Comment `json:"comments"`
	}

	Comment struct {
		ID	int64 `json:"id"`
		UserID int64 `json:"user_id"`
		Username string `json:"username"`
		CommentContent string `json:"commentContent"`
	}
)
