package entities

// type Users struct {
// 	gorm.Model
// 	Name   string
// 	Email  string
// 	Status UserStatus
// }

type UserStatus uint

const UserRegistered UserStatus = 101
const UserConfirmed UserStatus = 102
const UserBlocked UserStatus = 103
