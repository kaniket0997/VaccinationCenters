package controllers

import (
	"VaccinationCenter/dto"
	"VaccinationCenter/service"
	"fmt"
)

type UserController struct {
	idGlobal int64
	Country  *dto.Country
	userSvc  service.UserService
}

func InitUserController(idGlobal int64, country *dto.Country, userSvc service.UserService) IUserOperationHandler {

	return &UserController{
		Country:  country,
		idGlobal: idGlobal,
		userSvc:  userSvc,
	}
}

type IUserOperationHandler interface {
	OnBoardUser(uniqId, name string, sex string, state string, district string, age int64) error
	PrintUserDetails(userId string)
}

func (uc *UserController) OnBoardUser(uniqId, name string, sex string, state string, district string, age int64) error {
	if age < 18 {
		return fmt.Errorf("user is not eligible for registration, age=%d", age)
	}
	err := uc.userSvc.UserOnBoard(uniqId, name, sex, state, district, age)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserController) PrintUserDetails(userId string) {

	uc.userSvc.UserDetailsPrint(userId)
}
