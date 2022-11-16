package login

import (
	"errors"
	"fmt"
)

var UserData map[string]string

func init() {
	UserData = map[string]string{
		"test": "123456",
	}
}

func Login() {
	var name, password string
	fmt.Printf("What is your name? ")
	fmt.Scanf("%s", &name)
	fmt.Printf("What is the password? ")
	fmt.Scanf("%s", &password)

	if userPass, _ := UserData[name]; userPass == password {
		fmt.Println("Login Success!")
	} else {
		fmt.Println("Login Error!")
	}
}

func CheckUserIsExist(username string) bool {
	_, isExist := UserData[username]
	return isExist
}

func CheckPassword(p1 string, p2 string) error {
	if p1 == p2 {
		return nil
	} else {
		return errors.New("password is not correct")
	}
}

func Auth(username string, password string) error {
	if isExist := CheckUserIsExist(username); isExist {
		return CheckPassword(UserData[username], password)
	} else {
		return errors.New("user is not exist")
	}
}
