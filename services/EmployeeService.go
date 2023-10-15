package services

import (
	model "mekari/models"
	repo "mekari/repositories"
	"time"
)

func CreateEmployee(request model.Employee) model.ResponseData {
	var employee repo.Employee

	//validate if email already exist
	dataEmployee, errFind := employee.FindByEntity(model.Employee{Email: request.Email})
	if dataEmployee != nil && errFind == nil {
		return model.ResponseData{
			Error:  errFind,
			Errors: "email already exist",
		}
	}

	// save employee
	err := employee.Create(request)
	if err != nil {
		return model.ResponseData{
			Error:  err,
			Errors: err.Error(),
		}
	}

	return model.ResponseData{
		Message: "Success Create Employee",
	}
}

func GetEmployees() model.ResponseData {
	var employee repo.Employee
	result, err := employee.Get()
	if err != nil {
		return model.ResponseData{
			Error:  err,
			Errors: err.Error(),
		}
	}

	return model.ResponseData{
		Data:    result,
		Message: "Success Get Employees",
	}
}

func FindEmployee(id int) model.ResponseData {
	var employee repo.Employee
	result, err := employee.Find(id)
	if err != nil {
		return model.ResponseData{
			Error:  err,
			Errors: err.Error(),
		}
	}

	return model.ResponseData{
		Data:    result,
		Message: "Success Get Employee",
	}
}

func UpdateEmployee(payload model.Employee) model.ResponseData {
	var employee repo.Employee

	// validate id employee
	dataEmployee, err := employee.Find(payload.ID)
	if err != nil {
		return model.ResponseData{
			Error:  err,
			Errors: err.Error(),
		}
	}

	if payload.FirstName == "" {
		payload.FirstName = dataEmployee.FirstName
	}
	if payload.LastName == "" {
		payload.LastName = dataEmployee.LastName
	}
	if payload.Email == "" {
		payload.Email = dataEmployee.Email
	}
	timeZero := time.Time{}
	if payload.HireDate == timeZero {
		payload.HireDate = dataEmployee.HireDate
	}

	// update employee
	result, err := employee.Update(payload)
	if err != nil {
		return model.ResponseData{
			Error:  err,
			Errors: err.Error(),
		}
	}

	return model.ResponseData{
		Data:    result,
		Message: "Success Updated Employee",
	}
}

func DeleteEmployee(id int) model.ResponseData {
	var employee repo.Employee

	// validate id employee
	_, err := employee.Find(id)
	if err != nil {
		return model.ResponseData{
			Error:  err,
			Errors: err.Error(),
		}
	}

	// delete employee
	err = employee.Delete(id)
	if err != nil {
		return model.ResponseData{
			Error:  err,
			Errors: err.Error(),
		}
	}

	return model.ResponseData{
		Message: "Success Deleted Employee",
	}
}
