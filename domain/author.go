package domain

//go:generate mockgen -destination=../resources/mock/generated/mock_$GOFILE -source=$GOFILE -package=mocks

import (
	"context"
	"time"
)

// Author ...
type Author struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"-" swaggerignore:"true"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"-" swaggerignore:"true"`
}

func (Author) TableName() string {
	return "author"
}

// AuthorRepository ...
type AuthorRepository interface {
	GetByID(ctx context.Context, id int64) (Author, error)
}
