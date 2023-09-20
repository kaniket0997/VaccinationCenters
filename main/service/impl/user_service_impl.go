package impl

import (
	"VaccinationCenter/dto"
	"errors"
	"fmt"
)

type UserServiceImpl struct {
	Country  *dto.Country
	idGlobal int64
}

func InitUserService(idGlobal int64, country *dto.Country) *UserServiceImpl {

	return &UserServiceImpl{
		Country:  country,
		idGlobal: idGlobal,
	}
}

func (u *UserServiceImpl) UserOnBoard(uniqId string, name string, sex string, state string, district string, age int64) error {

	if _, ok := u.Country.Users[uniqId]; ok {
		fmt.Printf("user already present with this uniq id, uniqId=%s", uniqId)
		return errors.New("user already present with this uniq id")
	}
	user := dto.User{
		Id:       uniqId,
		Name:     name,
		Sex:      sex,
		Age:      age,
		State:    state,
		District: district,
	}
	if u.Country.Users == nil {
		u.Country.Users = make(map[string]dto.User)
	}
	u.Country.Users[user.Id] = user // Add the user to the Users map
	return nil
}

func (u *UserServiceImpl) UserDetailsPrint(userId string) {

	user, ok := u.Country.Users[userId]
	if !ok {
		fmt.Printf("user not found, userId=%s", userId)
		return
	}
	fmt.Printf("User Details: ID=%s Name=%s Sex=%s Age=%d State=%s District=%s\n", user.Id, user.Name, user.Sex,
		user.Age, user.State, user.District)
}
