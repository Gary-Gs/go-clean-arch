package mysql

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/Gary-Gs/go-clean-arch/domain"
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
	return res, m.db.WithContext(ctx).Find(&res).Error
}

// GetByID will get the article by primary key
func (m *mysqlArticleRepository) GetByID(ctx context.Context, id int64) (res domain.Article, err error) {
	return res, m.db.WithContext(ctx).Where("id = ?", id).First(&res).Error
}

// Upsert will update or create the article
func (m *mysqlArticleRepository) Upsert(ctx context.Context, o *domain.Article) (err error) {
	return m.db.WithContext(ctx).Model(o).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(o).Error
}

// Delete will delete the article by primary key
func (m *mysqlArticleRepository) Delete(ctx context.Context, id int64) (err error) {
	return m.db.WithContext(ctx).Delete(&domain.Article{}, id).Error
}
