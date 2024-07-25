package entities

import "time"

type AccessToken struct {
	ID          string    `gorm:"default:uuid_generate_v4()" json:"id"`
	UserId      string    `json:"user_id"`
	User        User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	AccessToken string    `json:"access_token"`
	ExpiredAt   time.Time `json:"expired_at"`
	Expired     bool      `json:"expired"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `gorm:"<-:update" json:"updated_at"`
}
