package main

import (
	"fmt"
	"math/rand"
)

func generatePassword(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+<>?"

	password := make([]byte, length)
	for i := range password {
		password[i] = chars[rand.Intn(len(chars))]
	}

	return string(password)
}

func RunPasswordGenerator() {
	var length int
	fmt.Print("How long should your password be? ")
	fmt.Scan(&length)

	if length <= 0 {
		fmt.Println("Come on, enter a valid number.")
		return
	}

	pass := generatePassword(length)
	fmt.Println("Here's your secure password:", pass)
}
