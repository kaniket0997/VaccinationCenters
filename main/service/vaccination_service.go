package service

import "VaccinationCenter/dto"

type VaccinationCentreService interface {
	AddVaccinationCenter(district dto.District, stateName, districtName, centerId string)
	AddCapacity(centerId string, day, maxCapacity int64)
	ListAllVaccinationCenters(districtName string)
	BookVaccinationSlot(userId, centerId string, day int64) error
	CancelVaccinationSlot(userId, bookingId, centerId string) error
}
