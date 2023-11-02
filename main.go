package main

import (
	"fmt"
	"image/jpeg"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
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
	options int
	lenPwd int
	username string
	system string
	filename string
	s1 string
	s2 string
	gree string
	// data1 []byte
	minSpecialChar = 1
	minNum = 1
	minUpperCase = 1
	flag = 0
	// i = 5.0
)

// const (
// 	W = 1024
// 	H = 512
// 	h = 24
// )

func main() {
	for {
		if flag == 0 {
			get_menu()
			fmt.Print("Option: ")
			fmt.Scanln(&options)
			flag = 1
		}
		
		if options == 1 {
			fmt.Print("Length: ")
			if _, err := fmt.Scanln(&lenPwd); err != nil {
				fmt.Println("Value is not a number")
				Cleanner()
			} else {
				fmt.Print("Username: ")
				fmt.Scanln(&username)
				fmt.Print("System: ")
				fmt.Scanln(&system)
				fmt.Print("Filename: ")
				fmt.Scanln(&filename)
				fmt.Print("Gretings: ")
				fmt.Scanln(&s1,&s2)
				gree = s1 + " " +s2
				pwd := GeneratePassword(lenPwd, minSpecialChar, minNum, minUpperCase)
				data1 := GetContentEmail(gree, system, strconv.Itoa(lenPwd), username, pwd)
				err := os.WriteFile("./txt/"+filename+".txt", data1, 0644)
				check(err)
				fmt.Println("TXT generate successfully")
				GetPdf(filename)
				PdfToImage(filename)
				CropImage(filename)
				Cleanner()
				flag = 0
			}
		} else if options == 2 {
			break
		} else {
			fmt.Println("Options not valid")
			Cleanner()
			flag = 0
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
		"\nCONTRASEÑA",
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
	check(err)
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
		check(err)
		
		for n := 0; n < doc.NumPage(); n++ {
			img, err := doc.Image(n)
			check(err)
			
			f, err := os.Create(filepath.Join("./img/",  filename+".jpg"))
			check(err)

			err = jpeg.Encode(f, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
			check(err)
			f.Close()
		}
	}
}

func CropImage(filename string) {
	src, err := imaging.Open("./img/"+filename+".jpg")
	check(err)
	src = imaging.CropAnchor(src, 2410, 2120, imaging.TopLeft)
	src = imaging.Resize(src, 700, 616, imaging.Lanczos)
	err = imaging.Save(src, "./img/"+filename+".jpg")
	check(err)
}

// func SaveToImage() {
// 	var files []string
// 	root := "/home/tech/go/src/genpwd/"
// 	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
// 		if filepath.Ext(path) == ".pdf" {
// 			files = append(files, path)
// 		}
// 		return nil
// 	})
// 	check(err)

// 	for _, file := range files {
// 		doc, err := fitz.New(file)
// 		check(err)
// 		folder :=strings.TrimSuffix(path.Base(file), filepath.Ext(path.Base(file)))

// 		for n := 0; n < doc.NumPage(); n++ {
// 			img, err := doc.Image(n)
// 			check(err)
// 			err = os.MkdirAll("img/"+folder, 0755)
// 			check(err)
// 			f, err := os.Create(filepath.Join("img/"+folder+"/", fmt.Sprintf("image-%05d.jpg", n)))
// 			check(err)
// 			err = jpeg.Encode(f, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
// 			check(err)
// 			f.Close()
// 		}
// 	}
// }

// func GetDraw(filename string, lines []string) {
// 	dc := gg.NewContext(W, H)
// 	dc.SetRGB(0, 0, 0)
// 	//dc.LoadFontFace("/Library/Fonts/Impact.ttf", 128)
	
// 	for i, line := range lines {
// 		y := H/2 - h*len(lines)/2 + i*h
// 		x := float64(W/2)
// 		dc.DrawStringWrapped(line, x+3, float64(y)+3 , 0.5, 0.5, W, 1.5, gg.AlignLeft)
// 		// dc.DrawString(line, x+3, float64(y)+3)
// 	}
	
// 	mask := dc.AsMask()

// 	dc.SetRGB(1, 1, 1)
// 	dc.Clear()

// 	g := gg.NewLinearGradient(0, 0, W, H)
// 	g.AddColorStop(0, color.RGBA{50, 48, 48, 255})
// 	g.AddColorStop(1, color.RGBA{31, 27, 27, 255})
// 	// g.AddColorStop(0, color.RGBA{255, 0, 0, 255})
// 	// g.AddColorStop(1, color.RGBA{0, 0, 255, 255})
// 	dc.SetFillStyle(g)

// 	dc.SetMask(mask)
// 	dc.DrawRectangle(0, 0, W, H)
// 	dc.Fill()
// 	dc.SavePNG(filename+".png")
// }