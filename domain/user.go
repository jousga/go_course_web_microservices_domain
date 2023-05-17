package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID string `json:"id" gorm:"type:char(36);not null; primary_key;unique_index"`
	//ID        uint       `json:"id"`
	FirstName string         `json:"first_name" gorm:"type:char(50);not null"`
	LastName  string         `json:"last_name" gorm:"type:char(50);not null"`
	Email     string         `json:"email" gorm:"type:char(50);not null"`
	Phone     string         `json:"phone" gorm:"type:char(20);not null"`
	CreatedAt *time.Time     `json:"-"` //GORM se encarga de setear este campo automáticamente
	UpdatedAt *time.Time     `json:"-"` //GORM se encarga de setear este campo automáticamente
	Deleted   gorm.DeletedAt `json:"-"` //GORM se encarga de setear este campo automáticamente, esto hace que el delete sea soft es decir, que el registro no se borrar de la base de dastos, solo se rellena esta fecha y GORM ya se encarga de no mostrar estos valores. Esto es útil para tener un histórico
}

// BeforeCreate Esta función actúa como Hook de GORM para añadir un ID al user antes de guardarlos en la BD
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return
}
