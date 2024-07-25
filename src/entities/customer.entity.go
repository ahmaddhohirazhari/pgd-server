package entities

import "time"

type Customer struct {
	ID        string    `gorm:"default:uuid_generate_v4()" json:"id"`
	Phone     string    `json:"phone"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	BirthDate time.Time `json:"birth_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type CustomerDto struct {
	Phone     string    `json:"phone"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	BirthDate time.Time `json:"birth_date"`
}
