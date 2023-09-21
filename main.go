package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var (
	lowerCharSet = "abcdefghijklmnopqrstuvwxyz"
	upperCharSet = "ABCDEFGHIGKLMNOPQRSTUVWXYZ"
	specialCharSet = "!#$%&*"
	numberSet = "0123456789"
	allCharSet = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func main() {
	for {
		var options int
		fmt.Println("Ingrese una opción:")
		fmt.Println("1. Generar password")
		fmt.Println("2 Salir")
		fmt.Scanln(&options)
		if options == 1 {
			// rand.Seed(time.Now().Unix())
			minSpecialChar := 1
			minNum := 1
			minUpperCase := 1
			fmt.Println("Length")
			var lenPwd int
			fmt.Scanln(&lenPwd)
			fmt.Println("Username")
			var username string
			fmt.Scanln(&username)
			var system string
			fmt.Println("System")
			fmt.Scanln(&system)
			pwd := generatePassword(lenPwd, minSpecialChar, minNum, minUpperCase)
			fmt.Println(pwd)
			fmt.Println(username)
			d1 := []byte("Hello\ngo\n")
			err := os.WriteFile("./dat1.txt", d1, 0644)
			check(err)
			// f, err := os.Create("./dat2")
			// check(err)
			// defer f.Close()
		}

		if options == 2 {
			break
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
	var password strings.Builder

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remaingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remaingLength; i ++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j],inRune[i]
	})

	return string(inRune)
}