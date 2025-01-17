package commonentities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type BaseEntity struct {
	gorm.Model
	ID        string `gorm:"type:uuid;primary_key;"`
	CreatedAt int
	UpdatedAt int
	DeletedAt soft_delete.DeletedAt
}

func (b *BaseEntity) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.NewString()
	}
	if b.CreatedAt == 0 {
		b.CreatedAt = int(time.Now().Unix())
	}
	return nil
}

func (b *BaseEntity) BeforeUpdate(tx *gorm.DB) (err error) {
	b.UpdatedAt = int(time.Now().Unix())
	return nil
}
