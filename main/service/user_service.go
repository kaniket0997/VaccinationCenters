package service

type UserService interface {
	UserOnBoard(uniqId, name string, sex string, state string, district string, age int64) error
	UserDetailsPrint(userId string)
}
