package controllers

import (
	"github.com/gin-gonic/gin"
	model "mekari/models"
	service "mekari/services"
	"net/http"
	"strconv"
	"time"
)

func Create(context *gin.Context) {
	var payload model.PayloadEmployee
	if err := context.ShouldBindJSON(&payload); err != nil {
		context.JSON(http.StatusBadRequest, model.ResponseData{
			Errors: "invalid json",
		})
		return
	}

	// validate hire date
	hireDate, err := time.Parse(model.DateLayoutISO, payload.HireDate)
	if err != nil {
		context.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			model.ResponseData{
				Errors: "invalid format date. Use DD/MM/YYYY Format",
			})
		return
	}

	responseService := service.CreateEmployee(model.Employee{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		HireDate:  hireDate,
	})
	if responseService.Error != nil {
		if responseService.Error.Error() == "record not found" {
			context.AbortWithStatusJSON(
				http.StatusNotFound,
				model.ResponseData{
					Errors: responseService.Errors,
				})
			return
		}
		context.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			model.ResponseData{
				Errors:  responseService.Errors,
				Message: responseService.Message,
			})
		return
	}

	context.JSON(http.StatusOK, responseService)
}

func GetEmployees(context *gin.Context) {
	responseService := service.GetEmployees()
	if responseService.Error != nil {
		context.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			model.ResponseData{
				Errors: responseService.Errors,
			})
		return
	}
	context.JSON(http.StatusOK, responseService.Data)
}

func FindEmployee(context *gin.Context) {
	id := context.Param("id")
	idRequest, err := strconv.Atoi(id)
	if err != nil {
		context.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			model.ResponseData{
				Errors: "invalid id employee",
			})
		return
	}

	responseService := service.FindEmployee(idRequest)
	if responseService.Error != nil {
		context.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			model.ResponseData{
				Errors: responseService.Errors,
			})
		return
	}
	context.JSON(http.StatusOK, responseService.Data)
}

func UpdateEmployee(context *gin.Context) {
	var payload model.PayloadEmployee
	if err := context.ShouldBindJSON(&payload); err != nil {
		context.JSON(http.StatusBadRequest, model.ResponseData{
			Errors: "invalid json",
		})
		return
	}

	id := context.Param("id")
	idRequest, err := strconv.Atoi(id)
	if err != nil {
		context.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			model.ResponseData{
				Errors: "invalid id employee",
			})
		return
	}

	// validate hire date
	hireDate, err := time.Parse(model.DateLayoutISO, payload.HireDate)
	if err != nil {
		context.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			model.ResponseData{
				Errors: "invalid format date. Use DD/MM/YYYY Format",
			})
		return
	}

	responseService := service.UpdateEmployee(model.Employee{
		ID:        idRequest,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		HireDate:  hireDate,
	})
	if responseService.Error != nil {
		if responseService.Error.Error() == "record not found" {
			context.AbortWithStatusJSON(
				http.StatusNotFound,
				model.ResponseData{
					Errors: responseService.Errors,
				})
			return
		}
		context.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			model.ResponseData{
				Errors:  responseService.Errors,
				Message: responseService.Message,
			})
		return
	}

	context.JSON(http.StatusOK, responseService)
}

func DeleteEmployee(context *gin.Context) {
	id := context.Param("id")
	idRequest, err := strconv.Atoi(id)
	if err != nil {
		context.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			model.ResponseData{
				Errors: "invalid id employee",
			})
		return
	}

	responseService := service.DeleteEmployee(idRequest)
	if responseService.Error != nil {
		if responseService.Error.Error() == "record not found" {
			context.AbortWithStatusJSON(
				http.StatusNotFound,
				model.ResponseData{
					Errors: responseService.Errors,
				})
			return
		}
		context.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			model.ResponseData{
				Errors:  responseService.Errors,
				Message: responseService.Message,
			})
		return
	}

	context.JSON(http.StatusOK, responseService)
}
