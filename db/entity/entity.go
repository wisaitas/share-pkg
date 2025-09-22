package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Entity struct {
	Id        uuid.UUID       `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Version   int             `gorm:"type:integer;not null;default:0"`
	CreatedAt time.Time       `gorm:"type:timestamp;not null;default:now()"`
	CreatedBy *uuid.UUID      `gorm:"type:uuid"`
	UpdatedAt time.Time       `gorm:"type:timestamp;not null;default:now()"`
	UpdatedBy *uuid.UUID      `gorm:"type:uuid"`
	DeletedAt *gorm.DeletedAt `gorm:"type:timestamp"`
	DeletedBy *uuid.UUID      `gorm:"type:uuid"`
}

func (r *Entity) BeforeUpdate(tx *gorm.DB) (err error) {
	r.Version++
	return
}
