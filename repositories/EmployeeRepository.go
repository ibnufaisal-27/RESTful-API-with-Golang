package repositories

import (
	model "mekari/models"
)

type Employee model.Employee

func (e Employee) Create(request model.Employee) error {
	err := db.Create(&request).Error
	if err != nil {
		return err
	}
	return nil
}

func (e Employee) Get() ([]model.Employee, error) {
	employees := make([]model.Employee, 0)

	err := db.Model(Employee{}).Find(&employees).Error

	if err != nil {
		return nil, err
	}
	return employees, nil
}

func (e Employee) Find(id int) (*Employee, error) {
	err := db.Model(Employee{}).Where("id = ?", id).Take(&e).Error

	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (e Employee) FindByEntity(request model.Employee) (*Employee, error) {
	err := db.Model(Employee{}).Where(&request).Take(&e).Error

	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (e Employee) Update(request model.Employee) (*model.Employee, error) {
	err := db.Model(Employee{}).Where("id = ?", request.ID).Updates(&model.Employee{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		HireDate:  request.HireDate,
	}).Error

	if err != nil {
		return nil, err
	}

	return &request, nil
}

func (e Employee) Delete(id int) error {
	err := db.Delete(Employee{}, id).Error

	if err != nil {
		return err
	}
	return nil
}
