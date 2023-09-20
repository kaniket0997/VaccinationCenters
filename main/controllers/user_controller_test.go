package controllers

import (
	"testing"
)

// Define a mock UserService for testing purposes
type MockUserService struct{}

func (m *MockUserService) UserDetailsPrint(userId string) {
	//TODO implement me
	panic("implement me")
}

func (m *MockUserService) UserOnBoard(name string, sex string, state string, district string, age int64) error {
	// Implement the mock behavior here if needed
	// For testing, you can simply return nil to indicate success
	return nil
}

func TestOnBoardUser(t *testing.T) {
	// Create an instance of UserController with the mock UserService
	mockUserService := &MockUserService{}
	userController := UserController{
		userSvc: mockUserService,
	}

	// Test case 1: Valid user registration
	err := userController.OnBoardUser("John Doe", "Male", "State1", "District1", 25)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Test case 2: Invalid user registration (age < 18)
	err = userController.OnBoardUser("Alice Smith", "Female", "State2", "District2", 16)
	expectedErrorMsg := "user is not eligible for registration, age=16"
	if err == nil || err.Error() != expectedErrorMsg {
		t.Errorf("Expected error: %v, but got: %v", expectedErrorMsg, err)
	}
}
