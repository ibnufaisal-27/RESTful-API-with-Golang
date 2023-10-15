package controllers

import (
	"github.com/gin-gonic/gin"
	model "mekari/models"
	service "mekari/services"
	"net/http"
	"strconv"
	"time"
)

// CreateEmployee godoc
// @Summary Create New Employee based on payload
// @Description Create New Employee
// @ID create-employee
// @Accept json
// @Produce json
// @Param employee body models.Employee true "Employee Data"
// @Success 200 {object} model.ResponseData{data=models.Employee,code=int,message=string}
// @Failure 422 {object} model.ResponseData{code=int,message=string}
// @Router /employees [post]
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

	context.JSON(http.StatusCreated, responseService)
}

// GetEmployees godoc
// @Summary Get All Employee
// @Description Retrive Data All Employee
// @ID get-employees
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseData{data=models.Employee,code=int,message=string}
// @Failure 422 {object} model.ResponseData{code=int,message=string}
// @Router /employees [get]
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

// FindEmployee godoc
// @Summary Find Employee detail
// @Description Find Data Employee by ID
// @ID find-employee
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} model.ResponseData{data=models.Employee,code=int,message=string}
// @Failure 400 {object} model.ResponseData{code=int,message=string}
// @Failure 422 {object} model.ResponseData{code=int,message=string}
// @Router /employees/:id [get]
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
				Errors: responseService.Errors,
			})
		return
	}
	context.JSON(http.StatusOK, responseService.Data)
}

// UpdateEmployee godoc
// @Summary Update Employee data
// @Description Update Data Employee by ID
// @ID update-employee
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} model.ResponseData{data=models.Employee,code=int,message=string}
// @Failure 400 {object} model.ResponseData{code=int,message=string}
// @Failure 422 {object} model.ResponseData{code=int,message=string}
// @Router /employees [put]
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

// DeleteEmployee godoc
// @Summary Delete Employee data
// @Description Delete Data Employee by ID
// @ID delete-employee
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} model.ResponseData{data=models.Employee,code=int,message=string}
// @Failure 400 {object} model.ResponseData{code=int,message=string}
// @Failure 422 {object} model.ResponseData{code=int,message=string}
// @Router /employees [delete]
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
