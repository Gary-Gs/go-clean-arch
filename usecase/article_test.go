package usecase

import (
	"context"
	"github.com/Gary-Gs/go-clean-arch/domain"
	mocks "github.com/Gary-Gs/go-clean-arch/resources/mock/generated"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

// GoMock interface mocking usage example
func TestArticleUsecase_GetByID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockArticleRepo := mocks.NewMockArticleRepository(mockCtrl)
	mockArticleRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(domain.Article{
		ID: int64(1),
	}, nil)
	mockAuthorRepo := mocks.NewMockAuthorRepository(mockCtrl)
	mockAuthorRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(domain.Author{
		ID: int64(1),
	}, nil)

	u := articleUsecase{
		articleRepo: mockArticleRepo,
		authorRepo:  mockAuthorRepo,
	}
	res, err := u.GetByID(context.Background(), 1)
	if err != nil {
		t.Errorf("error: %v", err.Error())
	}

	assert.Equal(t, int64(1), res.ID)
	assert.Equal(t, int64(1), res.Author.ID)
}
