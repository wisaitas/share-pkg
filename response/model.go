package response

import (
	"time"

	"github.com/google/uuid"
)

type EntityResponse struct {
	Id        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy *uuid.UUID `json:"created_by"`
	UpdatedAt time.Time  `json:"updated_at"`
	UpdatedBy *uuid.UUID `json:"updated_by"`
}

type ApiResponse[T any] struct {
	Message        string      `json:"message"`
	HttpStatusCode int         `json:"http_status_code"`
	Code           *string     `json:"code,omitempty"`
	Data           T           `json:"data,omitempty"`
	Error          error       `json:"error,omitempty"`
	Timestamp      time.Time   `json:"timestamp"`
	Pagination     *Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
	Page     int  `json:"page"`
	PageSize int  `json:"page_size"`
	HasPrev  bool `json:"has_prev"`
	HasNext  bool `json:"has_next"`
	Total    int  `json:"total"`
}
