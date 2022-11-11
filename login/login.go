package main

import (
	"fmt"
)

func main() {
	var name, password string
	db := map[string]string{
		"admin": "112233",
		"user":  "123456",
	}

	fmt.Printf("What is your name? ")
	fmt.Scanf("%s\n", &name)
	fmt.Printf("What is the password? ")
	fmt.Scanf("%s\n", &password)

	if dbPassword, _ := db[name]; dbPassword == password {
		fmt.Printf("Login Success!")
	} else {
		fmt.Printf("Login Error!")
	}

}
