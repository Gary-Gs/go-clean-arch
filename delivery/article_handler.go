package delivery

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	"github.com/Gary-Gs/go-clean-arch/domain"
)

// ArticleHandler ...
type ArticleHandler struct {
	AUsecase domain.ArticleUsecase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewArticleHandler(e *echo.Echo, us domain.ArticleUsecase) {
	handler := &ArticleHandler{
		AUsecase: us,
	}
	e.POST("/articles", handler.CreateOrUpdate)
	e.GET("/articles", handler.FetchArticle)
	e.GET("/articles/:id", handler.GetByID)
	e.DELETE("/articles/:id", handler.Delete)
}

func (a *ArticleHandler) CreateOrUpdate(c echo.Context) (err error) {
	var article domain.Article
	if err = c.Bind(&article); err != nil {
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

	if err = a.AUsecase.Upsert(c.Request().Context(), &article); err != nil {
		return c.JSON(getStatusCode(err), HttpResponse{
			Code:    getStatusCode(err),
			Message: err.Error(),
		})
	}
	return c.JSON(getStatusCode(err), HttpResponse{
		Code:    getStatusCode(err),
		Message: OK,
		Data:    article,
	})
}

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

func (a *ArticleHandler) GetByID(c echo.Context) error {
	ids := c.Param("id")
	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil {
		return c.JSON(getStatusCode(err), HttpResponse{
			Code:    getStatusCode(err),
			Message: BadRequest,
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

func (a *ArticleHandler) Delete(c echo.Context) (err error) {
	ids := c.Param("id")
	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil {
		return c.JSON(getStatusCode(err), HttpResponse{
			Code:    getStatusCode(err),
			Message: BadRequest,
		})
	}

	if err = a.AUsecase.Delete(c.Request().Context(), id); err != nil {
		return c.JSON(getStatusCode(err), HttpResponse{
			Code:    getStatusCode(err),
			Message: err.Error(),
		})
	}
	return c.NoContent(http.StatusNoContent)
}
