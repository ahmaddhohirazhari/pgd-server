package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"pgd-server.com/helpers"
	"pgd-server.com/src/entities"
)

type UserController struct{}

func (uc *UserController) Create(c *gin.Context) {
	response := helpers.GeneralResponse{}
	userDto := entities.UserDto{}

	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, helpers.GeneralResponse{
			Status:  false,
			Message: "Gagal",
			Error:   "Data input tidak valid",
		})
		return
	}

	result, err := userService.CrateUser(&userDto, c)
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Error = "Gagal memproses data"
		c.JSON(http.StatusConflict, response)
		return
	} else {
		response.Status = true
		response.Message = "Kami telah mengirim email dengan kode verifikasi ke " + userDto.Email
		response.Data = result
		c.JSON(http.StatusOK, response)
		return
	}
}

func (uc *UserController) FindAll(c *gin.Context) {
	response := new(helpers.GeneralResponse)
	results, _ := userService.FindAllUser(c)

	response.Status = true
	response.Message = "Berhasil"
	response.Data = results

	c.JSON(http.StatusOK, response)
}

func (uc *UserController) FindOne(c *gin.Context) {
	response := new(helpers.GeneralResponse)
	result, err := userService.FindOneUser(c.Param("id"), c)

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

func (uc *UserController) Update(c *gin.Context) {
	response := helpers.GeneralResponse{}
	userDto := new(entities.UserDto)

	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, helpers.GeneralResponse{
			Status:  false,
			Message: "Gagal",
			Error:   "Data input tidak valid",
		})
		return
	}

	result, err := userService.UpdateUser(c.Param("id"), userDto, c)
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

func (uc *UserController) Delete(c *gin.Context) {
	response := new(helpers.GeneralResponse)
	if err := userService.DeleteUser(c.Param("id"), c); err != nil {
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
