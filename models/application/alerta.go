package application

import (
	"time"

	"gorm.io/gorm"
)

type Alerta struct {
	ID          int            `json:"alerta_id"`
	Latitude    string         `json:"latitude"`
	Longitude   string         `json:"longitude"`
	UsuarioID   int            `json:"usuario_id,omitempty"`
	Usuario     *AppUser       `json:"usuario,omitempty"`
	Estado      string         `json:"estado"`
	Multimedias []Multimedia   `json:"multimedias"`
	CreatedAt   time.Time      `json:"created"`
	UpdatedAt   time.Time      `json:"updated,omitempty"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted,omitempty"`
}
