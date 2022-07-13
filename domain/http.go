package domain

const (
	OK         = "OK"
	BadRequest = "bad request"
)

// ResponseError ...
type ResponseError struct {
	Message string `json:"message"`
}

// HttpResponse ...
// @Description http response
type HttpResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Pagination ...
// @Description pagination
type Pagination struct {
	Page      int    `json:"page,omitempty" query:"page" validate:"gte=1,lte=1000"`
	Size      int    `json:"size,omitempty" query:"size" validate:"gte=1,lte=1000"`
	Sort      string `json:"sort,omitempty" query:"sort"`
	TotalPage int64  `json:"total_pages,omitempty"`
}

func NewPagination() Pagination {
	return Pagination{
		Page: 1,
		Size: 1000,
	}
}
