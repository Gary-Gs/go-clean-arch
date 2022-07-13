package domain

//go:generate mockgen -destination=../resources/mock/generated/mock_$GOFILE -source=$GOFILE -package=mocks

import (
	"context"
	"time"
)

// Article ...
// @Description article
type Article struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title" validate:"required"`
	Content   string     `json:"content" validate:"required"`
	AuthorID  int64      `json:"author_id" validate:"required"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"-" swaggerignore:"true"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"-" swaggerignore:"true"`
}

func (Article) TableName() string {
	return "article"
}

// ArticlesResponse ...
// @Description articles response
type ArticlesResponse struct {
	ArticleWithAuthors []*ArticleWithAuthor `json:"articles"`
	Pagination         Pagination           `json:"pagination"`
}

type ArticleWithAuthor struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  Author `json:"author"`
}

// ArticleUsecase ...
type ArticleUsecase interface {
	Fetch(ctx context.Context, pagination Pagination) (*ArticlesResponse, error)
	GetByID(ctx context.Context, id int64) (*ArticleWithAuthor, error)
	Upsert(ctx context.Context, article *Article) error
	Delete(ctx context.Context, id int64) error
}

// ArticleRepository ...
type ArticleRepository interface {
	Fetch(ctx context.Context, pagination Pagination) (res []Article, err error)
	GetByID(ctx context.Context, id int64) (Article, error)
	Upsert(ctx context.Context, article *Article) error
	Delete(ctx context.Context, id int64) error
	CountAll(ctx context.Context) (int64, error)
}
