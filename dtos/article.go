package dtos

import "time"

type CreateArticlesRequest struct {
	AdministratorID uint   `json:"administrator_id" form:"administrator_id" example:"1"`
	Thumbnail       string `json:"thumbnail" form:"thumbnail" example:"gambar1.jpg"`
	Title           string `json:"title" form:"title" example:"judulArticle"`
	Abstract        string `json:"abstract" form:"abstract" example:"abstract/pengantar"`
	Description     string `json:"description" form:"description" example:"isi artikel"`
	Image           string `json:"image" form:"image" example:"link image"`
	Label           string `json:"label" form:"label" example:"kebugaran"`
}

type UpdateArticlesRequest struct {
	AdministratorID uint   `json:"administrator_id" form:"administrator_id" example:"1"`
	Thumbnail       string `json:"thumbnail" form:"thumbnail" example:"gambar1.jpg"`
	Title           string `json:"title" form:"title" example:"judulArticle" validate:"required"`
	Abstract        string `json:"abstract" form:"abstract" example:"abstract/pengantar" validate:"required"`
	Description     string `json:"description" form:"description" example:"isi artikel" validate:"required"`
	Image           string `json:"image" form:"image" example:"gambar2.jpg"`
	Label           string `json:"label" form:"label" example:"kebugaran" validate:"required"`
}

type CreateArticlesResponse struct {
	Thumbnail   string `json:"thumbnail" form:"thumbnail" example:"gambar1.jpg"`
	Title       string `json:"title" form:"title" example:"judulArticle"`
	Abstract    string `json:"abstract" form:"abstract" example:"abstract/pengantar"`
	Description string `json:"description" form:"description" example:"isi artikel"`
	Image       string `json:"image" form:"image" example:"gambar2.jpg"`
	Label       string `json:"label" form:"label" example:"kebugaran"`
	Slug        string `json:"slug" form:"slug" example:"judularticle"`
}

type UpdateArticleResponse struct {
	Thumbnail   string `json:"thumbnail" form:"thumbnail" example:"gambar1.jpg"`
	Title       string `json:"title" form:"title" example:"judulArticle"`
	Abstract    string `json:"abstract" form:"abstract" example:"abstract/pengantar"`
	Description string `json:"description" form:"description" example:"isi artikel"`
	Image       string `json:"image" form:"image" example:"gambar2.jpg"`
	Label       string `json:"label" form:"label" example:"kebugaran"`
}

type ArticleDetailResponse struct {
	ArticleID   uint      `json:"article_id" example:"1"`
	Thumbnail   string    `json:"thumbnail" form:"thumbnail" example:"gambar1.jpg"`
	Title       string    `json:"title" form:"title" example:"judulArticle"`
	Abstract    string    `json:"abstract" form:"abstract" example:"abstract/pengantar"`
	Image       string    `json:"image" form:"image" example:"gambar2.jpg"`
	Description string    `json:"description" form:"description" example:"isi artikel"`
	Label       string    `json:"label" form:"label" example:"kebugaran"`
	Slug        string    `json:"slug" form:"slug" example:"judularticle"`
	CreatedAt   time.Time `json:"created_at" example:"2023-05-17T15:07:16.504+07:00"`
	UpdatedAt   time.Time `json:"updated_at" example:"2023-05-17T15:07:16.504+07:00"`
}
