package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/devnandito/genpwd/utils"
)

var (
	minSpecialChar = 1
	minNum = 1
	minUpperCase = 1
	options int
	lenPwd int
	username string
	system string
	filename string
	s1 string
	s2 string
	gree string
	flag = 0
)

func main() {
	for {
		if flag == 0 {
			utils.GetMenu()
			fmt.Print("Option: ")
			fmt.Scanln(&options)
			flag = 1
		}
		
		if options == 1 {
			fmt.Print("Length: ")
			if _, err := fmt.Scanln(&lenPwd); err != nil {
				fmt.Println("Value is not a number")
				utils.Cleanner()
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
				pwd := utils.GeneratePassword(lenPwd, minSpecialChar, minNum, minUpperCase)
				data1 := utils.GetContentEmail(gree, system, strconv.Itoa(lenPwd), username, pwd)
				err := os.WriteFile("./txt/"+filename+".txt", data1, 0644)
				utils.Check(err)
				fmt.Println("TXT generate successfully")
				utils.GetPdf(filename)
				utils.PdfToImage(filename)
				utils.CropImage(filename)
				utils.Cleanner()
				flag = 0
			}
		} else if options == 2 {
			break
		} else {
			fmt.Println("Options not valid")
			utils.Cleanner()
			flag = 0
		}
	}
}