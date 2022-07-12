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
	e.POST("/api/v1/articles", handler.CreateOrUpdate)
	e.GET("/api/v1/articles", handler.FetchArticle)
	e.GET("/api/v1/articles/:id", handler.GetByID)
	e.DELETE("/api/v1/articles/:id", handler.Delete)
}

// CreateOrUpdate godoc
// @Summary      Create or update articles
// @Description  Create or update articles
// @Tags         articles
// @Accept       json
// @Produce      json
// @Param  article  body      domain.Article  true  "article object"
// @Success      200  {object}   HttpResponse
// @Failure      400  {object}  HttpResponse
// @Failure      404  {object}  HttpResponse
// @Failure      500  {object}  HttpResponse
// @Router       /api/v1/articles [post]
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

// FetchArticle godoc
// @Summary      Get articles
// @Description  Get articles
// @Tags         articles
// @Accept       json
// @Produce      json
// @Success      200  {object}   HttpResponse
// @Failure      400  {object}  HttpResponse
// @Failure      404  {object}  HttpResponse
// @Failure      500  {object}  HttpResponse
// @Router       /api/v1/articles [get]
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

// GetByID godoc
// @Summary      Get articles by id
// @Description  Get articles by id
// @Tags         articles
// @Accept       json
// @Produce      json
// @Param id   path      int  true  "article ID"
// @Success      200  {object}   HttpResponse
// @Failure      400  {object}  HttpResponse
// @Failure      404  {object}  HttpResponse
// @Failure      500  {object}  HttpResponse
// @Router       /api/v1/articles/{id} [get]
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

// Delete godoc
// @Summary      Delete articles by id
// @Description  Delete articles by id
// @Tags         articles
// @Accept       json
// @Produce      json
// @Param id   path      int  true  "article ID"
// @Success      200  {object}   HttpResponse
// @Failure      400  {object}  HttpResponse
// @Failure      404  {object}  HttpResponse
// @Failure      500  {object}  HttpResponse
// @Router       /api/v1/articles/{id} [delete]
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
