package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Course struct {
	ID        string         `json:"id" gorm:"type:char(36);not null; primary_key;unique_index"`
	Name      string         `json:"name" gorm:"type:char(50);not null"`
	StartDate time.Time      `json:"start_date"`
	EndDate   time.Time      `json:"end_date"`
	CreatedAt *time.Time     `json:"-"`
	UpdatedAt *time.Time     `json:"-"`
	Deleted   gorm.DeletedAt `json:"-"`
}

// BeforeCreate Esta función actúa como Hook de GORM para añadir un ID al user antes de guardarlos en la BD
func (c *Course) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return
}
