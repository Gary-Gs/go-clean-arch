package mysql

import (
	"context"
	_ "database/sql"
	"gorm.io/gorm"

	"github.com/Gary-Gs/go-clean-arch/domain"
	_ "github.com/sirupsen/logrus"
)

type mysqlArticleRepository struct {
	db *gorm.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewMysqlArticleRepository(conn *gorm.DB) domain.ArticleRepository {
	return &mysqlArticleRepository{db: conn}
}

// Fetch will get all articles
func (m *mysqlArticleRepository) Fetch(ctx context.Context) (res []domain.Article, err error) {
	err = m.db.Find(&res).Error
	return res, err
}

// GetByID will get the article by primary key
func (m *mysqlArticleRepository) GetByID(ctx context.Context, id int64) (res domain.Article, err error) {
	m.db.Where("id = ?", id).First(&res)
	return
}

// GetByTitle will get the article by title
func (m *mysqlArticleRepository) GetByTitle(ctx context.Context, title string) (res domain.Article, err error) {
	m.db.Where("title = ?", title).First(&res)
	return
}

// Store will create a new article
func (m *mysqlArticleRepository) Store(ctx context.Context, a *domain.Article) (err error) {
	m.db.Create(a)
	return
}

// Delete will delete the article by primary key
func (m *mysqlArticleRepository) Delete(ctx context.Context, id int64) (err error) {
	m.db.Delete(&domain.Article{}, id)
	return
}

// Update will update the article by struct
func (m *mysqlArticleRepository) Update(ctx context.Context, o *domain.Article) (err error) {
	m.db.Model(&o).Updates(&o)
	return
}
