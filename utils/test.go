package utils

// i = 5.0
// const (
// 	W = 1024
// 	H = 512
// 	h = 24
// )
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