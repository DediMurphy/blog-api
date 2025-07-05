package posts

import "time"

type (
	CreateCommentsRequest struct {
		CommentContent string `json:"commentContent"`
	}
)

type (
	CommentModel struct {
		ID             int64     `db:"id"`
		UserID         int64     `db:"user_id"`
		PostID         int64     `db:"post_id"`
		CommentContent string    `db:"comment_content"`
		CreatedAt      time.Time `db:"created_at"`
		UpdatedAt      time.Time `db:"up	dated_at"`
		CreatedBy      string    `db:"created_by"`
		UpdatedBy      string    `db:"update_by"`
	}
)
