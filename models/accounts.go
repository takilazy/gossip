package models

import (
	u "github.com/takilazy/gossip/utils"
)

type Account struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

//func (account *Account) Validate() (map[string] interface{}, bool)  {
//
//}
//Create
func (account *Account) Create() (map[string] interface{})  {
	account.InsertUser()

	response := u.Message(true, "Account has been created")
	response["account"] = account
	return response
}
//Login
//GetUser
func (account *Account) Get() (map[string] interface{})  {
	account.GetUser()

	response := u.Message(true, "Anus created")
	response["account"] = account
	return response
}