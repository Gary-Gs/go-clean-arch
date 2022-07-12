package domain

import (
	"context"
	"time"
)

// Author ...
type Author struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"-"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"-"`
}

func (Author) TableName() string {
	return "author"
}

// AuthorRepository represent the author's repository contract
type AuthorRepository interface {
	GetByID(ctx context.Context, id int64) (Author, error)
}
