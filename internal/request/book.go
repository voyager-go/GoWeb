package request

import "github.com/voyager-go/GoWeb/pkg/formatTime"

type BookListReq struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Tags     string `json:"tags"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}

type BookCreateReq struct {
	Title             string          `json:"title"`
	Description       string          `json:"description"`
	Tags              string          `json:"tags"`
	Cover             string          `json:"cover"`
	AuthorName        string          `json:"author_name"`
	AuthorDescription string          `json:"author_description"`
	Status            int8            `json:"status"`
	CreatedAt         formatTime.Time `json:"created_at"`
	UpdatedAt         formatTime.Time `json:"updated_at"`
}
