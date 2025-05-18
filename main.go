package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Сколько символов должен содержать пароль?")
	var length int
	fmt.Scan(&length)

	if length <= 0 {
		fmt.Println("Длина должна быть положительной!")
		return
	}

	if length < 3 {
		fmt.Println("Пароль не может содержать менее трех символов")
		return
	}

	specials := "!@#$%^&*"
	password := ""
	hasDigit := false
	usedChars := make(map[rune]bool)

	for len(password) < length {
		choice := rand.Intn(3)

		var char rune
		if choice == 0 {
			char = rune('0' + rand.Intn(10))
			hasDigit = true
		} else if choice == 1 {
			char = rune('a' + rand.Intn(26))
		} else {
			char = rune(specials[rand.Intn(len(specials))])
		}
		if !usedChars[char] {
			password += string(char)
			usedChars[char] = true
		}
	}
	if !hasDigit {
		if len(password) > 0 {
			password = string('0'+rune(rand.Intn(10))) + password[1:]
		}
	}

	fmt.Println("Ваш пароль", password)
}
