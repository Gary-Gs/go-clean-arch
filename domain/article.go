package domain

import (
	"context"
	"time"
)

// Article ...
type Article struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	AuthorID  int64     `json:"author_id" validate:"required"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (Article) TableName() string {
	return "article"
}

type ArticleResponse struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  Author `json:"author"`
}

// ArticleUsecase represent the article's usecases
type ArticleUsecase interface {
	Fetch(ctx context.Context) ([]*ArticleResponse, error)
	GetByID(ctx context.Context, id int64) (*ArticleResponse, error)
	Update(ctx context.Context, ar *Article) error
	Store(context.Context, *Article) error
	Delete(ctx context.Context, id int64) error
}

// ArticleRepository represent the article's repository contract
type ArticleRepository interface {
	Fetch(ctx context.Context) (res []Article, err error)
	GetByID(ctx context.Context, id int64) (Article, error)
	Update(ctx context.Context, ar *Article) error
	Store(ctx context.Context, a *Article) error
	Delete(ctx context.Context, id int64) error
}
