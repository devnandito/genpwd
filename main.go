package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/fogleman/gg"
)

var (
	lowerCharSet = "abcdefghijklmnopqrstuvwxyz"
	upperCharSet = "ABCDEFGHIGKLMNOPQRSTUVWXYZ"
	specialCharSet = "!#$%&*"
	numberSet = "0123456789"
	allCharSet = lowerCharSet + upperCharSet + specialCharSet + numberSet
	options int
	lenPwd int
	username string
	system string
	filename string
	s1 string
	s2 string
	gree string
	data1 []byte
	minSpecialChar = 1
	minNum = 1
	minUpperCase = 1
	i = 5.0
)

const (
	W = 1024
	H = 512
	h = 24
)

func main() {
	// dc := gg.NewContext(1000, 1000)
	// dc.DrawCircle(500, 500, 400)
	// dc.SetRGB(0, 0, 0)
	// dc.Fill()
	// dc.SavePNG("out.png")

	for {
		get_menu()
		fmt.Print("Option: ")
		fmt.Scanln(&options)
		if options == 1 {
			// rand.Seed(time.Now().Unix())
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
			data1, lines := get_content_email(gree, system, strconv.Itoa(lenPwd), username, pwd)
			err := os.WriteFile("./"+filename+".txt", data1, 0644)
			check(err)
			get_draw(filename, lines)
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

func get_draw(filename string, lines []string) {
	dc := gg.NewContext(W, H)
	dc.SetRGB(0, 0, 0)
	//dc.LoadFontFace("/Library/Fonts/Impact.ttf", 128)
	
	for i, line := range lines {
		y := H/2 - h*len(lines)/2 + i*h
		x := float64(W/2)
		dc.DrawStringWrapped(line, x+3, float64(y)+3 , 0.5, 0.5, W, 1.5, gg.AlignLeft)
	}
	
	mask := dc.AsMask()

	dc.SetRGB(1, 1, 1)
	dc.Clear()

	g := gg.NewLinearGradient(0, 0, W, H)
	g.AddColorStop(0, color.RGBA{50, 48, 48, 255})
	g.AddColorStop(1, color.RGBA{31, 27, 27, 255})
	// g.AddColorStop(0, color.RGBA{255, 0, 0, 255})
	// g.AddColorStop(1, color.RGBA{0, 0, 255, 255})
	dc.SetFillStyle(g)

	dc.SetMask(mask)
	dc.DrawRectangle(0, 0, W, H)
	dc.Fill()
	dc.SavePNG(filename+".png")
}

func get_content_email(gree, system, lenPwd, username, pwd string) ([]byte, []string) {
	var lines = []string {
		gree,
		fmt.Sprintf("Por este medio remito las credenciales de acceso a SIPAP %s", system),
		"Tener en cuenta las siguientes políricas de seguridad al cambiar la contraseña:",
		"- No debe ser una palabra que pueda estar en algún diccionario, ni relacionado al nombre de usuario o entidad.",
		"- No debe ser correlativo al anterior, ni similar en las últimas 10 contraseñas.",
		fmt.Sprintf("- Debe contener por los menos una mayúscula, una minúscula, un número y un carácter especial y de longitud mínima de %s.", lenPwd),
		"- Vigencia mínima de contraseña, de 1(un) días. Un solo cambio por día es posible.",
		"- Escriba su contraseña en un bloc de notas para luego copiar y pegar en el sistema.",
		"USUARIO",
		username,
		"CONTRASEÑA",
		pwd,
		"Favor dar acuse.",
		"Saludos cordiales.",

	}

	m := make(map[int]string)
	m[1] = gree
	m[2] = fmt.Sprintf("Por este medio remito las credenciales de acceso a SIPAP %s", system)
	m[3] = "Tener en cuenta las siguientes políricas de seguridad al cambiar la contraseña:"
	m[4] = "- No debe ser una palabra que pueda estar en algún diccionario, ni relacionado al nombre de usuario o entidad."
	m[5] = "- No debe ser correlativo al anterior, ni similar en las últimas 10 contraseñas."
	m[6] = fmt.Sprintf("- Debe contener por los menos una mayúscula, una minúscula, un número y un carácter especial y de longitud mínima de %s.", lenPwd)
	m[7] = "- Vigencia mínima de contraseña, de 1(un) días. Un solo cambio por día es posible."
	m[8] = "- Escriba su contraseña en un bloc de notas para luego copiar y pegar en el sistema."
	m[9] = "\nUSUARIO"
	m[10] = username
	m[11] = "\nCONTRASEÑA"
	m[12] = pwd
	m[13] = "\nFavor dar acuse."
	m[14] = "Saludos cordiales."

	// data1 := []byte(string(m["gree"]+"\n"+m["p1"]+"\n"+m["p2"]+"\n"+m["p3"]+"\n"+m["p4"]+"\n"+m["p5"]+"\n"+m["p6"]+"\n"+m["p7"]+"\n"+m["p8"]+"\n"+m["p9"]+"\n"+m["p10"]+"\n"+m["p11"]+"\n"+m["p12"]+"\n"+m["p13"]))
	data1 := []byte(string(m[1]+"\n"+m[2]+"\n"+m[3]+"\n"+m[4]+"\n"+m[5]+"\n"+m[6]+"\n"+m[7]+"\n"+m[8]+"\n"+m[9]+"\n"+m[10]+"\n"+m[11]+"\n"+m[12]+"\n"+m[13]+"\n"+m[14]))
	return data1, lines
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