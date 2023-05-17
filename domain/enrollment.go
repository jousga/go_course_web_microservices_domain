package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Enrollement Con gorm al definir UserID y después User, o CourseID y después Course, él se encarga de hacer la relación automaticamente interpretando que el campo con ID es la foreignKey de esa entidad

type Enrollement struct {
	ID        string     `json:"id" gorm:"type:char(36);not null; primary_key;unique_index"`
	UserID    string     `json:"user_id,omitempty" gorm:"type:char(36)"` //Al usar el nombre de la entidad seguido de ID el ya sabe que será la foreignkey para esa entidad
	User      *User      `json:"user,omitempty"`
	CourseID  string     `json:"course_id" gorm:"type:char(36);not null"`
	Course    *Course    `json:"course,omitempty"`
	Status    string     `json:"status" gorm:"type:char(2)"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
}

// BeforeCreate Esta función actúa como Hook de GORM para añadir un ID al user antes de guardarlos en la BD
func (u *Enrollement) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return
}
