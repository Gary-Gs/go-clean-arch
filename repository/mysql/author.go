package mysql

import (
	"context"
	"gorm.io/gorm"

	"github.com/Gary-Gs/go-clean-arch/domain"
)

type mysqlAuthorRepo struct {
	db *gorm.DB
}

// NewMysqlAuthorRepository will create an implementation of author.Repository
func NewMysqlAuthorRepository(db *gorm.DB) domain.AuthorRepository {
	return &mysqlAuthorRepo{db: db}
}

// GetByID will get the author by primary key
func (m *mysqlAuthorRepo) GetByID(ctx context.Context, id int64) (res domain.Author, err error) {
	m.db.Where("id = ?", id).First(&res)
	return
}
