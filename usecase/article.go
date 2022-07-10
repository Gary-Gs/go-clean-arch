package usecase

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	"github.com/Gary-Gs/go-clean-arch/domain"
)

type articleUsecase struct {
	articleRepo    domain.ArticleRepository
	authorRepo     domain.AuthorRepository
	contextTimeout time.Duration
}

// NewArticleUsecase will create new an articleUsecase object representation of domain.ArticleUsecase interface
func NewArticleUsecase(a domain.ArticleRepository, ar domain.AuthorRepository, timeout time.Duration) domain.ArticleUsecase {
	return &articleUsecase{
		articleRepo:    a,
		authorRepo:     ar,
		contextTimeout: timeout,
	}
}

/*
* In this function below, I'm using errgroup with the pipeline pattern
* Look how this works in this package explanation
* in godoc: https://godoc.org/golang.org/x/sync/errgroup#ex-Group--Pipeline
 */
func (a *articleUsecase) fillAuthorDetails(c context.Context, data []domain.Article) ([]*domain.ArticleResponse, error) {
	g, ctx := errgroup.WithContext(c)
	res := make([]*domain.ArticleResponse, 0)

	// Get the author's id
	mapAuthors := map[int64]domain.Author{}

	for _, article := range data {
		mapAuthors[article.AuthorID] = domain.Author{}
	}
	// Using goroutine to fetch the author's detail
	chanAuthor := make(chan domain.Author)
	for authorID := range mapAuthors {
		authorID := authorID
		g.Go(func() error {
			res, err := a.authorRepo.GetByID(ctx, authorID)
			if err != nil {
				return err
			}
			chanAuthor <- res
			return nil
		})
	}

	go func() {
		err := g.Wait()
		if err != nil {
			logrus.Error(err)
			return
		}
		close(chanAuthor)
	}()

	for author := range chanAuthor {
		if author != (domain.Author{}) {
			mapAuthors[author.ID] = author
		}
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	// merge the author's data
	for _, item := range data {
		if a, ok := mapAuthors[item.AuthorID]; ok {
			res = append(res, &domain.ArticleResponse{
				ID:      item.ID,
				Title:   item.Title,
				Content: item.Content,
				Author:  a,
			})
		}
	}
	return res, nil
}

func (a *articleUsecase) Fetch(c context.Context) ([]*domain.ArticleResponse, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	articles, err := a.articleRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	res, err := a.fillAuthorDetails(ctx, articles)
	return res, err
}

func (a *articleUsecase) GetByID(c context.Context, id int64) (*domain.ArticleResponse, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	article, err := a.articleRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if article.ID == 0 {
		return nil, domain.ErrNotFound
	}

	author, err := a.authorRepo.GetByID(ctx, article.AuthorID)
	if err != nil {
		return nil, err
	}
	if author.ID == 0 {
		return nil, domain.ErrNotFound
	}

	return &domain.ArticleResponse{
		ID:      article.ID,
		Title:   article.Title,
		Content: article.Content,
		Author:  author,
	}, nil
}

func (a *articleUsecase) Update(c context.Context, ar *domain.Article) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	ar.UpdatedAt = time.Now()
	return a.articleRepo.Update(ctx, ar)
}

func (a *articleUsecase) Store(c context.Context, m *domain.Article) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	err = a.articleRepo.Store(ctx, m)
	return
}

func (a *articleUsecase) Delete(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedArticle, err := a.articleRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existedArticle == (domain.Article{}) {
		return domain.ErrNotFound
	}
	return a.articleRepo.Delete(ctx, id)
}
