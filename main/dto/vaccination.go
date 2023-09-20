package dto

type Country struct {
	Name   string `json:"name"`
	States map[string]State
	Users  map[string]User
}

type State struct {
	Districts map[string]District `json:"districts"`
}

type District struct {
	VaccinationCenters map[string]VaccinationCenter `json:"vaccination_centers"` // Center Id-> key
}

type VaccinationCenter struct {
	DayDetails []DayDetail `json:"day_details"` // 7 days
}

type DayDetail struct {
	MaxCapacity    int64             `json:"max_capacity"`
	AvailableSlots int64             `json:"available_slots"`
	Bookings       map[string]string `json:"bookings"` // booking id to user id map
}

type User struct {
	Id       string `json:"idGlobal"`
	Name     string `json:"name"`
	Sex      string `json:"sex"`
	Age      int64  `json:"age"`
	State    string `json:"state"`
	District string `json:"district"`
}
