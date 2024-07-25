package services

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	"pgd-server.com/config"
	"pgd-server.com/src/entities"
)

type UserService struct{}

func (us *UserService) CrateUser(dto *entities.UserDto, c *gin.Context) (*entities.User, error) {
	var user = new(entities.User)
	user.Email = dto.Email
	user.Name = dto.Name
	user.Phone = dto.Phone
	user.Address = dto.Address
	user.BirthDate = dto.BirthDate

	user.Password = dto.Password
	user.CreatedAt = time.Now()

	user.HashPassword()

	if result := config.DB.Create(&user); result.Error != nil {
		return nil, result.Error

	}

	return user, nil

}

func (us *UserService) FindAllUser(c *gin.Context) (pg paginate.Page, err error) {

	var result []entities.UserResult
	model := config.DB.
		Model(&entities.User{})

	model.
		Order("created_at desc").
		Find(&result)

	users := config.PG.
		With(model).
		Request(c.Request).
		Response(&[]entities.UserResult{})
	return users, err
}

func (us *UserService) FindOneByEmailOrPhone(username string) (*entities.User, error) {
	var user entities.User
	result := config.DB.
		Where("phone = ? OR email = ?", username, username).
		First(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (us *UserService) FindOneUser(id string, c *gin.Context) (*entities.User, error) {

	var user entities.User
	model := config.DB.
		Model(&entities.User{}).
		Where("users.id = ?", id).
		First(&user)

	if model.Error != nil {
		return nil, model.Error
	}

	return &user, nil
}

func (us *UserService) UpdateUser(id string, dto *entities.UserDto, c *gin.Context) (*entities.User, error) {

	user := new(entities.User)
	user.Email = dto.Email
	user.Phone = dto.Phone
	user.Name = dto.Name
	user.Address = dto.Address
	user.BirthDate = dto.BirthDate

	result := config.DB.
		Model(&entities.User{}).
		Where("id = ?", id).
		Updates(user)
	updateModel, err := us.FindOneUser(id, c)
	if err != nil {
		return updateModel, err
	}
	return updateModel, result.Error
}

func (us *UserService) DeleteUser(id string, c *gin.Context) *error {
	model, err := us.FindOneUser(id, c)
	if err != nil {
		return &err
	}
	result := config.DB.Where("id = ?", model.ID).Delete(&entities.User{})

	if result.Error != nil {
		return &result.Error
	}

	return nil
}
