package model

import "time"

// User is the golang structure for table user.
type User struct {
	Id       uint      `json:"id"       description:"User ID"`
	Passport string    `json:"passport" description:"User Passport"`
	Password string    `json:"password" description:"User Password"`
	Nickname string    `json:"nickname" description:"User Nickname"`
	CreateAt time.Time `json:"createAt" description:"Created Time"`
	UpdateAt time.Time `json:"updateAt" description:"Updated Time"`
}
