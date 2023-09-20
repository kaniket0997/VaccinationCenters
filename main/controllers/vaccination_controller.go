package controllers

import (
	"VaccinationCenter/dto"
	"VaccinationCenter/service"
	"fmt"
)

type VaccinationController struct {
	bookingId      int64
	Country        *dto.Country
	vaccinationSvc service.VaccinationCentreService
}

func InitVaccinationController(bookingId int64, country *dto.Country,
	vaccinationSvc service.VaccinationCentreService) IVaccinationCentreOperationHandler {

	return &VaccinationController{
		bookingId:      bookingId,
		Country:        country,
		vaccinationSvc: vaccinationSvc,
	}
}

type IVaccinationCentreOperationHandler interface {
	AddVaccinationCenter(stateName, districtName, centerId string)
	AddCapacity(centerId string, day, maxCapacity int64)
	ListAllVaccinationCenters(districtName string)
	BookVaccinationSlot(userId, centerId string, day int64) error
	CancelVaccinationSlot(userId, bookingId, centerId string) error
}

func (vc *VaccinationController) AddVaccinationCenter(stateName, districtName, centerId string) {

	state, ok := vc.Country.States[stateName]
	if !ok {
		fmt.Printf("state not found, state=%s", stateName)
		return
	}
	district, ok := state.Districts[districtName]
	if !ok {
		fmt.Printf("district not found, district=%s", districtName)
		return
	}
	vc.vaccinationSvc.AddVaccinationCenter(district, stateName, districtName, centerId)
}

func (vc *VaccinationController) AddCapacity(centerId string, day, maxCapacity int64) {

	vc.vaccinationSvc.AddCapacity(centerId, day, maxCapacity)
}

func (vc *VaccinationController) ListAllVaccinationCenters(districtName string) {

	vc.vaccinationSvc.ListAllVaccinationCenters(districtName)
}

func (vc *VaccinationController) BookVaccinationSlot(userId, centerId string, day int64) error {

	if userId == "" {
		return fmt.Errorf("userId is empty")
	}
	if centerId == "" {
		return fmt.Errorf("centerId is empty")
	}
	_, ok := vc.Country.Users[userId]
	if !ok {
		return fmt.Errorf("user not found, userId=%s", userId)
	}

	err := vc.vaccinationSvc.BookVaccinationSlot(userId, centerId, day)
	if err != nil {
		return err
	}
	return nil

}

func (vc *VaccinationController) CancelVaccinationSlot(userId, bookingId, centerId string) error {
	if userId == "" {
		return fmt.Errorf("userId is empty")
	}
	if bookingId == "" {
		return fmt.Errorf("bookingId is empty")
	}
	if centerId == "" {
		return fmt.Errorf("centerId is empty")
	}
	_, ok := vc.Country.Users[userId]
	if !ok {
		return fmt.Errorf("user not found, userId=%s", userId)
	}
	err := vc.vaccinationSvc.CancelVaccinationSlot(userId, bookingId, centerId)
	if err != nil {
		return err
	}
	return nil
}
