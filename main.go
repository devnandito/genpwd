package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
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
	var options int
	var lenPwd int
	var username string
	var system string
	var filename string
	var s1 string
	var s2 string
	var gree string
	minSpecialChar := 1
	minNum := 1
	minUpperCase := 1
	for {
		get_menu()
		fmt.Print("Option: ")
		fmt.Scanln(&options)
		if options == 1 {
			// rand.Seed(time.Now().Unix())
			var data []byte
			fmt.Print("Length: ")
			fmt.Scanln(&lenPwd)
			fmt.Print("Username: ")
			fmt.Scanln(&username)
			fmt.Print("System: ")
			fmt.Scanln(&system)
			fmt.Print("Filename: ")
			fmt.Scanln(&filename)
			fmt.Print("Gretings: ")
			fmt.Scanln(&s1,&s2)
			gree = s1 + " " +s2
			pwd := generatePassword(lenPwd, minSpecialChar, minNum, minUpperCase)
			fmt.Println("Password generate success")
			data = get_content_email(gree, system, strconv.Itoa(lenPwd), username, pwd)
			err := os.WriteFile("./"+filename+".txt", data, 0644)
			check(err)
			// d1 := []byte("Hello\ngo\n")
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

func get_menu() {
	fmt.Println("Enter option")
	fmt.Println("1. Generate password")
	fmt.Println("2 Exit")
}

func get_content_email(gree, system, lenPwd, username, pwd string) []byte {
	m := make(map[string]string)
	m["gree"] = gree
	m["p1"] = fmt.Sprintf("Por este medio remito las credenciales de acceso a SIPAP %s", system)
	m["p2"] = "Tener en cuenta las siguientes políricas de seguridad al cambiar la contraseña:"
	m["p3"] = "- No debe ser una palabra que pueda estar en algún diccionario, ni relacionado al nombre de usuario o entidad."
	m["p4"] = "- No debe ser correlativo al anterior, ni similar en las últimas 10 contraseñas."
	m["p5"] = fmt.Sprintf("- Debe contener por los menos una mayúscula, una minúscula, un número y un carácter especial y de longitud mínima de %s.", lenPwd)
	m["p6"] = "- Vigencia mínima de contraseña, de 1(un) días. Un solo cambio por día es posible."
	m["p7"] = "- Escriba su contraseña en un bloc de notas para luego copiar y pegar en el sistema."
	m["p8"] = "\nUSUARIO"
	m["p9"] = username
	m["p10"] = "\nCONTRASEÑA"
	m["p11"] = pwd
	m["p12"] = "\nFavor dar acuse."
	m["p13"] = "Saludos cordiales."

	data := []byte(string(m["gree"]+"\n"+m["p1"]+"\n"+m["p2"]+"\n"+m["p3"]+"\n"+m["p4"]+"\n"+m["p5"]+"\n"+m["p6"]+"\n"+m["p7"]+"\n"+m["p8"]+"\n"+m["p9"]+"\n"+m["p10"]+"\n"+m["p11"]+"\n"+m["p12"]+"\n"+m["p13"]))
	return data
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