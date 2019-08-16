package models

import (
	u "github.com/takilazy/gossip/utils"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type Account struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (account *Account) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "email address is invalid"), false
	}

	if len(account.Password) < 10 {
		return u.Message(false, "password length cannot be < 10"), false
	}

	email,_ := account.GetUser()
	if email != "" {
		return u.Message(false, "email address is used"), false
	}

	return u.Message(false, "validation OK"), true

}

//Create
func (account *Account) Create() map[string]interface{} {

	msg, status := account.Validate()
	if !status {
		return msg
	}

	cryptedPassword, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.MinCost)
	if err != nil {
		return u.Message(false, "something wrong with password")
	}

	account.Password = string(cryptedPassword)
	account.InsertUser()

	response := u.Message(true, "Account has been created")
	response["account"] = account
	return response
}

//Login
func (account *Account) Login() map[string]interface{}{

	email,_ := account.GetUser()
	if email == "" {
		return u.Message(false, "email address is not found")
	}

	_, password := account.GetUser()
	err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "wrong password")
	}

	//need to made some tokens and other shit

	return u.Message(true, "ok, logged")


}
//GetUser
func (account *Account) Get() map[string]interface{} {
	user, _ := account.GetUser()

	if user == "" {
		return u.Message(false, "User not found")
	}

	return u.Message(true, "User found")
}
