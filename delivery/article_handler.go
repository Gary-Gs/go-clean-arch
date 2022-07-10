package delivery

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	"github.com/Gary-Gs/go-clean-arch/domain"
)

// ArticleHandler  represent the httphandler for article
type ArticleHandler struct {
	AUsecase domain.ArticleUsecase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewArticleHandler(e *echo.Echo, us domain.ArticleUsecase) {
	handler := &ArticleHandler{
		AUsecase: us,
	}
	e.GET("/articles", handler.FetchArticle)
	e.POST("/articles", handler.Store)
	e.GET("/articles/:id", handler.GetByID)
	e.DELETE("/articles/:id", handler.Delete)
}

// FetchArticle will fetch the article based on given params
func (a *ArticleHandler) FetchArticle(c echo.Context) error {
	ctx := c.Request().Context()
	res, err := a.AUsecase.Fetch(ctx)
	if err != nil {
		return c.JSON(getStatusCode(err), HttpResponse{
			Code:    getStatusCode(err),
			Message: err.Error(),
		})
	}
	return c.JSON(getStatusCode(err), HttpResponse{
		Code:    getStatusCode(err),
		Message: OK,
		Data:    res,
	})
}

// GetByID will get article by given id
func (a *ArticleHandler) GetByID(c echo.Context) error {
	ids := c.Param("id")
	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil {
		return c.JSON(getStatusCode(err), HttpResponse{
			Code:    getStatusCode(err),
			Message: err.Error(),
		})
	}

	res, err := a.AUsecase.GetByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(getStatusCode(err), HttpResponse{
			Code:    getStatusCode(err),
			Message: err.Error(),
		})
	}
	return c.JSON(getStatusCode(err), HttpResponse{
		Code:    getStatusCode(err),
		Message: OK,
		Data:    res,
	})
}

// Store will store the article by given request body
func (a *ArticleHandler) Store(c echo.Context) (err error) {
	var article domain.Article
	err = c.Bind(&article)
	if err != nil {
		return c.JSON(getStatusCode(err), HttpResponse{
			Code:    getStatusCode(err),
			Message: err.Error(),
		})
	}

	if err = c.Validate(article); err != nil {
		return c.JSON(getStatusCode(err), HttpResponse{
			Code:    getStatusCode(err),
			Message: err.Error(),
		})
	}

	err = a.AUsecase.Store(c.Request().Context(), &article)
	if err != nil {
		return c.JSON(getStatusCode(err), HttpResponse{
			Code:    getStatusCode(err),
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, HttpResponse{
		Code:    http.StatusCreated,
		Message: OK,
		Data:    article,
	})
}

// Delete will delete article by given param
func (a *ArticleHandler) Delete(c echo.Context) error {
	ids := c.Param("id")
	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil {
		return c.JSON(getStatusCode(err), HttpResponse{
			Code:    getStatusCode(err),
			Message: err.Error(),
		})
	}

	err = a.AUsecase.Delete(c.Request().Context(), id)
	if err != nil {
		return c.JSON(getStatusCode(err), HttpResponse{
			Code:    getStatusCode(err),
			Message: err.Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)
}
