package utils

import (
	"fmt"
	"image/jpeg"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/disintegration/imaging"
	"github.com/jung-kurt/gofpdf"
	"github.com/karmdip-mi/go-fitz"
)

var (
	lowerCharSet = "abcdefghijklmnopqrstuvwxyz"
	upperCharSet = "ABCDEFGHIGKLMNOPQRSTUVWXYZ"
	specialCharSet = "!#$%&*"
	numberSet = "0123456789"
	allCharSet = lowerCharSet + upperCharSet + numberSet
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetMenu() {
	fmt.Println("Enter option")
	fmt.Println("1. Generate password")
	fmt.Println("2 Exit")
}

func Cleanner() {
	time.Sleep(2 * time.Second)
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}


func GetContentEmail(gree, system, lenPwd, username, pwd string) ([]byte) {
	var lines = []string {
		gree,
		fmt.Sprintf("Por este medio remito las credenciales de acceso a SIPAP %s", system),
		"Tener en cuenta las siguientes políticas de seguridad al cambiar la contraseña:",
		"- No debe ser una palabra que pueda estar en algún diccionario, ni relacionado al nombre de usuario o entidad.",
		"- No debe ser correlativo al anterior, ni similar en las últimas 10 contraseñas.",
		fmt.Sprintf("- Debe contener por los menos una mayúscula, una minúscula, un número y un carácter especial y de longitud mínima de %s.", lenPwd),
		"- Vigencia mínima de contraseña, de 1(un) días. Un solo cambio por día es posible.",
		"- Escriba su contraseña en un bloc de notas para luego copiar y pegar en el sistema.",
		"\nUSUARIO",
		username,
		"CONTRASEÑA",
		pwd,
		"\nFavor dar acuse.",
		"Saludos cordiales.",
	}

	var txt string
	for _, value  := range lines {
		txt = txt + "\n" + value
	}
	data := []byte(string(txt))

	return data
}


func GeneratePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
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

	fmt.Println("Password generate success")
	return string(inRune)
}

func GetPdf(file string) {
	text, err := os.ReadFile("./txt/"+file+".txt")
	Check(err)
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	// pdf.SetFont("Arial", "B", 16)
	// pdf.MoveTo(0, 10)
	// pdf.Cell(1, 1, "Title")
	pdf.SetFont("Arial", "", 10)
	pdf.MoveTo(0, 0)
	tr := pdf.UnicodeTranslatorFromDescriptor("")
	width, _ := pdf.GetPageSize()
	pdf.MultiCell(width, 10, tr(string(text)), "" , "", false)
	err = pdf.OutputFileAndClose("./pdf/"+file+".pdf")

	if err == nil {
		fmt.Println("PDF generate successfully")
	}
}

func PdfToImage(filename string) {
	var files []string
	path := "./pdf/"+filename+".pdf"
	files = append(files, path)

	for _, file := range files {
		doc, err := fitz.New(file)
		Check(err)
		
		for n := 0; n < doc.NumPage(); n++ {
			img, err := doc.Image(n)
			Check(err)
			
			f, err := os.Create(filepath.Join("./img/",  filename+".jpg"))
			Check(err)

			err = jpeg.Encode(f, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
			Check(err)
			f.Close()
		}
	}
}

func CropImage(filename string) {
	src, err := imaging.Open("./img/"+filename+".jpg")
	Check(err)
	// src = imaging.CropAnchor(src, 2410, 2120, imaging.TopLeft)
	src = imaging.CropAnchor(src, 350, 1100, imaging.Left)
	src = imaging.CropAnchor(src, 350, 450, imaging.TopLeft)
	src = imaging.Resize(src, 250, 350, imaging.Lanczos)
	// src = imaging.Resize(src, 700, 616, imaging.Lanczos)
	err = imaging.Save(src, "./img/"+filename+".jpg")
	Check(err)
}