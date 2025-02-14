package mysql

import (
	"context"
	"gorm.io/gorm"

	"github.com/Gary-Gs/go-clean-arch/domain"
)

type mysqlAuthorRepo struct {
	db *gorm.DB
}

// NewMysqlAuthorRepository ...
func NewMysqlAuthorRepository(db *gorm.DB) domain.AuthorRepository {
	return &mysqlAuthorRepo{db: db}
}

// GetByID ...
func (m *mysqlAuthorRepo) GetByID(ctx context.Context, id int64) (res domain.Author, err error) {
	return res, m.db.WithContext(ctx).Where("id = ?", id).First(&res).Error
}
