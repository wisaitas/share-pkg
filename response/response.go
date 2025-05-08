package response

import (
	"time"

	"github.com/google/uuid"
	"github.com/wisaitas/share-pkg/pkg/db/model"
)

type Success struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type Error struct {
	Message string `json:"message"`
}

type Base struct {
	ID        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy *uuid.UUID `json:"created_by"`
	UpdatedAt time.Time  `json:"updated_at"`
	UpdatedBy *uuid.UUID `json:"updated_by"`
}

func (r *Base) ModelToResponse(baseModel model.Base) Base {
	r.ID = baseModel.ID
	r.CreatedAt = baseModel.CreatedAt
	r.UpdatedAt = baseModel.UpdatedAt
	r.CreatedBy = baseModel.CreatedBy
	r.UpdatedBy = baseModel.UpdatedBy

	return *r
}
