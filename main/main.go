package main

import (
	"VaccinationCenter/controllers"
	"VaccinationCenter/dto"
	"VaccinationCenter/service/impl"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var idGlobal int64 = 0
var bookingIdGlobal int64 = 0

func CreateCountry() *dto.Country {
	return &dto.Country{
		Name: "India",
		States: map[string]dto.State{
			"Karnataka": {
				//Name: "Karnataka",
				Districts: map[string]dto.District{
					"Bangalore": {
						//Name:               "Bangalore",
						VaccinationCenters: make(map[string]dto.VaccinationCenter),
					},
					"Mysore": {
						//Name:               "Mysore",
						VaccinationCenters: make(map[string]dto.VaccinationCenter),
					},
				},
			},
			"Kerala": {
				Districts: map[string]dto.District{},
			},
			"TamilNadu": {
				Districts: map[string]dto.District{},
			},
		},
	}
}

func main() {
	fmt.Println("Welcome to the Vaccination Appointment Booking System CLI")
	country := CreateCountry()
	userSvc := impl.InitUserService(idGlobal, country)
	vaccSvc := impl.InitVaccinationService(country)
	userCtrl := controllers.InitUserController(idGlobal, country, userSvc)
	vaccinationCtrl := controllers.InitVaccinationController(bookingIdGlobal, country, vaccSvc)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		commandParts := strings.Fields(input)
		if len(commandParts) == 0 {
			continue
		}

		command := commandParts[0]
		args := commandParts[1:]

		commandOrchestrator(userCtrl, vaccinationCtrl, command, args)
	}
}

func commandOrchestrator(uc controllers.IUserOperationHandler, vc controllers.IVaccinationCentreOperationHandler,
	command string, args []string) {

	switch command {
	case "ADD_USER":
		if len(args) < 6 {
			fmt.Println("Usage: ADD_USER <unique_identification> <name> <gender> <age> <current_state> <current_district>")
			break
		}
		uniqId := args[0]
		name := args[1]
		gender := args[2]
		age, _ := strconv.ParseInt(args[3], 10, 64)
		state := args[4]
		district := args[5]
		err := uc.OnBoardUser(uniqId, name, gender, state, district, age)
		if err != nil {
			fmt.Println(err)
		}
	case "PRINT_USER_DETAILS":
		if len(args) < 1 {
			fmt.Println("Usage: PRINT_USER_DETAILS <unique_identification>")
			break
		}
		userId := args[0]
		uc.PrintUserDetails(userId)
	case "ADD_VACCINATION_CENTER":
		state := args[0]
		district := args[1]
		centerId := args[2]
		vc.AddVaccinationCenter(state, district, centerId)
	case "ADD_CAPACITY":
		centerId := args[0]
		day, _ := strconv.ParseInt(args[1], 10, 64)
		maxCapacity, _ := strconv.ParseInt(args[2], 10, 64)
		vc.AddCapacity(centerId, day, maxCapacity)
	case "LIST_VACCINATION_CENTERS":
		district := args[0]
		vc.ListAllVaccinationCenters(district)
	case "BOOK_VACCINATION":
		centerId := args[0]
		day, _ := strconv.ParseInt(args[1], 10, 64)
		userId := args[2]
		err := vc.BookVaccinationSlot(userId, centerId, day)
		if err != nil {
			fmt.Println(err)
		}
	case "CANCEL_BOOKING":
		centerId := args[0]
		bookingId := args[1]
		userId := args[2]
		err := vc.CancelVaccinationSlot(userId, bookingId, centerId)
		if err != nil {
			fmt.Println(err)
		}
	case "exit":
		fmt.Println("Exiting Vaccination Appointment Booking System CLI.")
		return
	case "help":
		fmt.Println("Available commands:")
		fmt.Println("ADD_USER | ADD_VACCINATION_CENTER | ADD_CAPACITY ...")
		fmt.Println("LIST_VACCINATION_CENTERS | LIST_ALL_BOOKINGS")
		fmt.Println("BOOK_VACCINATION | CANCEL_BOOKING")
		fmt.Println("exit")
	default:
		fmt.Println("Invalid command. Type 'help' for available commands.")
	}
}
