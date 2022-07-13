package usecase

import (
	"context"
	"fmt"
	"math"
	"sync"
	"time"

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

// fillAuthorDetails is a concurrent function to fill author details implemented with goroutines and sync.WaitGroup
func (a *articleUsecase) fillAuthorDetails(c context.Context, data []domain.Article) (res []*domain.ArticleWithAuthor, errs []error) {
	// get all ids
	mapAuthors := map[int64]domain.Author{}
	for _, article := range data {
		mapAuthors[article.AuthorID] = domain.Author{}
	}

	var wg sync.WaitGroup
	resultChannel := make(chan domain.Author, len(mapAuthors))
	errChannel := make(chan error, len(mapAuthors))

	// get data with goroutine
	for k := range mapAuthors {
		wg.Add(1)
		go func(ctx context.Context, id int64, resChan chan domain.Author, errChan chan error, wg *sync.WaitGroup) {
			author, err := a.authorRepo.GetByID(ctx, id)
			defer func() {
				resChan <- author
				errChan <- err
				wg.Done()
			}()
		}(c, k, resultChannel, errChannel, &wg)
	}

	// goroutine synchronization
	wg.Wait()
	close(resultChannel)
	close(errChannel)

	// process data returned by goroutine
	for v := range resultChannel {
		if v != (domain.Author{}) {
			mapAuthors[v.ID] = v
		}
	}
	for v := range errChannel {
		if v != nil {
			errs = append(errs, v)
		}
	}
	if len(errs) > 0 {
		return nil, errs
	}

	// merge data
	for _, article := range data {
		author := mapAuthors[article.AuthorID]
		res = append(res, &domain.ArticleWithAuthor{
			ID:      article.ID,
			Title:   article.Title,
			Content: article.Content,
			Author:  author,
		})
	}

	return res, nil
}

func (a *articleUsecase) Fetch(c context.Context, pagination domain.Pagination) (res *domain.ArticlesResponse, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	articles, err := a.articleRepo.Fetch(ctx, pagination)
	if err != nil {
		return nil, err
	}

	awa, errs := a.fillAuthorDetails(ctx, articles)
	if errs != nil {
		return nil, fmt.Errorf("failed to fill author details: %v", errs)
	}

	count, err := a.articleRepo.CountAll(ctx)
	if err != nil {
		return nil, err
	}
	pagination.TotalPage = int64(math.Ceil(float64(count) / float64(pagination.Size)))
	return &domain.ArticlesResponse{
		ArticleWithAuthors: awa,
		Pagination:         pagination,
	}, err
}

func (a *articleUsecase) GetByID(c context.Context, id int64) (*domain.ArticleWithAuthor, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	article, err := a.articleRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	author, err := a.authorRepo.GetByID(ctx, article.AuthorID)
	if err != nil {
		return nil, err
	}
	return &domain.ArticleWithAuthor{
		ID:      article.ID,
		Title:   article.Title,
		Content: article.Content,
		Author:  author,
	}, nil
}

func (a *articleUsecase) Upsert(c context.Context, m *domain.Article) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	return a.articleRepo.Upsert(ctx, m)
}

func (a *articleUsecase) Delete(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	return a.articleRepo.Delete(ctx, id)
}
