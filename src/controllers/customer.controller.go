package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"pgd-server.com/helpers"
	"pgd-server.com/src/entities"
)

type CustomerController struct{}

func (cc *CustomerController) Create(c *gin.Context) {
	response := helpers.GeneralResponse{}
	customerDto := entities.CustomerDto{}

	if err := c.ShouldBindJSON(&customerDto); err != nil {
		c.JSON(http.StatusBadRequest, helpers.GeneralResponse{
			Status:  false,
			Message: "Gagal",
			Error:   "Data input tidak valid",
		})
		return
	}

	result, err := customerService.CreateCustomer(&customerDto, c)
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Error = "Gagal memproses data"
		c.JSON(http.StatusConflict, response)
		return
	} else {
		response.Status = true
		response.Message = "Berhasil"
		response.Data = result
		c.JSON(http.StatusOK, response)
		return

	}
}

func (cc *CustomerController) FindAll(c *gin.Context) {
	response := new(helpers.GeneralResponse)
	results, _ := customerService.FindAllCustomer(c)

	response.Status = true
	response.Message = "Berhasil"
	response.Data = results

	c.JSON(http.StatusOK, response)

}

func (cc *CustomerController) FindOne(c *gin.Context) {
	response := new(helpers.GeneralResponse)
	result, err := customerService.FindOneCustomer(c.Param("id"), c)

	if err != nil {
		response.Status = false
		response.Message = "Item not found"
		c.JSON(http.StatusNotFound, response)
		return
	}

	response.Status = true
	response.Message = "Berhasil"
	response.Data = result
	c.JSON(http.StatusOK, response)
}

func (cc *CustomerController) Update(c *gin.Context) {
	response := helpers.GeneralResponse{}
	customerDto := new(entities.CustomerDto)

	if err := c.ShouldBindJSON(&customerDto); err != nil {
		c.JSON(http.StatusBadRequest, helpers.GeneralResponse{
			Status:  false,
			Message: "Gagal",
			Error:   "Data input tidak valid",
		})
		return
	}

	result, err := customerService.UpdateCustomer(c.Param("id"), customerDto, c)
	if err != nil {
		response.Status = false
		response.Message = "Gagal"
		response.Error = "Gagal memproses data"
		c.JSON(http.StatusConflict, response)
		return
	} else {
		response.Status = true
		response.Message = "Berhasil"
		response.Data = result
		c.JSON(http.StatusOK, response)
		return
	}

}

func (cc *CustomerController) Delete(c *gin.Context) {
	response := new(helpers.GeneralResponse)
	if err := customerService.DeleteCustomer(c.Param("id"), c); err != nil {
		response.Status = false
		response.Message = "Gagal"
		response.Error = "Gagal memproses data"
		c.JSON(http.StatusConflict, response)
		return
	}

	response.Status = true
	response.Message = "Berhasil"
	response.Data = nil
	c.JSON(http.StatusOK, response)

}
