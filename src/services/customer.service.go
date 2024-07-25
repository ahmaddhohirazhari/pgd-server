package services

import (
	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	"pgd-server.com/config"
	"pgd-server.com/src/entities"
)

type CustomerService struct{}

func (cs *CustomerService) CreateCustomer(dto *entities.CustomerDto, c *gin.Context) (*entities.Customer, error) {
	var customer = new(entities.Customer)

	customer.Name = dto.Name
	customer.Phone = dto.Phone
	customer.Address = dto.Address
	customer.BirthDate = dto.BirthDate

	if result := config.DB.Create(&customer); result.Error != nil {
		return nil, result.Error
	}

	return customer, nil

}

func (cs *CustomerService) FindAllCustomer(c *gin.Context) (page *paginate.Page, err error) {
	customers := []entities.Customer{}

	model := config.DB.Order("created_at desc").Find(&customers)

	if model.Error != nil {
		return nil, model.Error
	}

	result := config.PG.
		With(model).
		Request(c.Request).
		Response(&customers)
	return &result, err

}

func (cs *CustomerService) FindOneCustomer(id string, c *gin.Context) (*entities.Customer, error) {
	costumer := entities.Customer{}

	if result := config.DB.Where("id = ?", id).First(&costumer); result.Error != nil {
		return nil, result.Error
	}

	return &costumer, nil
}

func (cs *CustomerService) UpdateCustomer(id string, dto *entities.CustomerDto, c *gin.Context) (*entities.Customer, error) {

	customer := new(entities.Customer)
	customer.Phone = dto.Phone
	customer.Name = dto.Name
	customer.Address = dto.Address
	customer.BirthDate = dto.BirthDate

	findCustomer, err := cs.FindOneCustomer(id, c)
	if err != nil {
		return findCustomer, err
	}

	result := config.DB.
		Model(&customer).
		Where("id = ?", id).
		Updates(&customer)
	if result.Error != nil {
		return nil, result.Error
	}

	return customer, result.Error
}

func (cs *CustomerService) DeleteCustomer(id string, c *gin.Context) *error {
	model, err := cs.FindOneCustomer(id, c)
	if err != nil {
		return &err
	}
	result := config.DB.Where("id = ?", model.ID).Delete(&entities.Customer{})

	if result.Error != nil {
		return &result.Error
	}

	return nil
}
