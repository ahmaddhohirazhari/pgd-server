package entities

import (
	"encoding/base64"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string    `gorm:"default:uuid_generate_v4()" json:"id"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	BirthDate time.Time `json:"birth_date"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type UserDto struct {
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	BirthDate time.Time `json:"birth_date"`
	Password  string    `json:"password"`
}

type UserResult struct {
	ID        string    `gorm:"default:uuid_generate_v4()" json:"id"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	BirthDate time.Time `json:"birth_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (user *User) HashPassword() {
	unhashedPass := user.Password
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(unhashedPass), 8)
	user.Password = string(hashedPass)
}

func Encode(s string) string {
	data := base64.StdEncoding.EncodeToString([]byte(s))
	return string(data)
}

func Decode(s string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
