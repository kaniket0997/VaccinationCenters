package impl

import (
	"VaccinationCenter/dto"
	"fmt"
)

type VaccinationServiceImpl struct {
	Country   *dto.Country
	bookingId int64 // Add a bookingId field
}

func (vc *VaccinationServiceImpl) AddVaccinationCenter(district dto.District, stateName, districtName, centerId string) {

	vaccinationCenter := dto.VaccinationCenter{
		//State:        stateName,
		//DistrictName: districtName,
		//CenterId:     centerId,
		DayDetails: make([]dto.DayDetail, 7), // Initialize with a slice of 7 DayDetail structs
	}
	if district.VaccinationCenters == nil {
		district.VaccinationCenters = make(map[string]dto.VaccinationCenter)
	}
	district.VaccinationCenters[centerId] = vaccinationCenter
	vc.Country.States[stateName].Districts[districtName] = district
}

func (vc *VaccinationServiceImpl) AddCapacity(centerId string, day, maxCapacity int64) {

	found := false
	for _, state := range vc.Country.States {
		for _, district := range state.Districts {
			if vaccinationCenter, ok := district.VaccinationCenters[centerId]; ok {
				found = true
				dayDetail := dto.DayDetail{
					MaxCapacity:    maxCapacity,
					AvailableSlots: maxCapacity,
				}
				vaccinationCenter.DayDetails[day-1] = dayDetail
			}
		}
	}
	if !found {
		fmt.Printf("vaccination center not found, centerId=%s", centerId)
		return
	}
}

func (vc *VaccinationServiceImpl) ListAllVaccinationCenters(districtName string) {

	found := false
	for stateName, state := range vc.Country.States {
		district, ok := state.Districts[districtName]
		if !ok {
			continue // Continue to the next district if the current one doesn't match the specified districtName
		}
		found = true
		for vaccinationCenterId, vaccinationCenter := range district.VaccinationCenters {
			fmt.Printf("vaccination center: %s %s %s\n", stateName, districtName,
				vaccinationCenterId)
			for _, dayDetail := range vaccinationCenter.DayDetails {
				for idx, bookingId := range dayDetail.Bookings {
					fmt.Printf("booking by user: %s for booking id %s \n", bookingId, idx)
				}
			}
		}
	}
	if !found {
		fmt.Printf("district not found, district=%s", districtName)
		return
	}
}

func InitVaccinationService(country *dto.Country) *VaccinationServiceImpl {
	return &VaccinationServiceImpl{
		Country:   country,
		bookingId: 0, // Initialize bookingId
	}
}

func (vc *VaccinationServiceImpl) BookVaccinationSlot(userId, centerId string, day int64) error {

	for _, state := range vc.Country.States {
		for _, district := range state.Districts {
			if vaccinationCenter, ok := district.VaccinationCenters[centerId]; ok {
				//dayNo := utils.ConvertDayToNumber(day)
				//if dayNo == -1 {
				//	return fmt.Errorf("invalid day, day=%s", day)
				//}
				if vaccinationCenter.DayDetails[day-1].AvailableSlots == 0 {
					return fmt.Errorf("no available slots for the specified day, day=%s", day)
				}
				if vaccinationCenter.DayDetails[day-1].Bookings == nil {
					vaccinationCenter.DayDetails[day-1].Bookings = make(map[string]string)
				}
				vc.bookingId++
				bookingIdString := fmt.Sprintf("%v", vc.bookingId)
				vaccinationCenter.DayDetails[day-1].Bookings[bookingIdString] = userId
				vaccinationCenter.DayDetails[day-1].AvailableSlots--
				return nil
			}
		}
	}
	return fmt.Errorf("vaccination center not found, centerId=%s", centerId)
}

func (vc *VaccinationServiceImpl) CancelVaccinationSlot(userId, bookingId, centerId string) error {

	found := false
	for idx1, state := range vc.Country.States {
		for idx2, district := range state.Districts {
			center, ok := district.VaccinationCenters[centerId]
			if !ok {
				continue
			}
			for idx3, dayDetail := range center.DayDetails {
				uId, ok := dayDetail.Bookings[bookingId]
				if !ok {
					continue
				}
				if uId != userId {
					return fmt.Errorf("booking not found for user, bookingId=%s, userId=%s", bookingId, userId)
				}
				found = true
				delete(dayDetail.Bookings, bookingId)
				vc.Country.States[idx1].Districts[idx2].VaccinationCenters[centerId].DayDetails[idx3].AvailableSlots++
				return nil
			}
		}
	}
	if !found {
		return fmt.Errorf("booking not found, bookingId=%s, centerId=%s, userId=%s", bookingId, centerId, userId)
	}
	return nil
}
