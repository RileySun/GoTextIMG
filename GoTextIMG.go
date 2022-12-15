package main

import(
	"fmt"
	"os"
	"time"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	
	"bufio"
	"golang.org/x/image/draw"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Println("Enter text to be turned into image")
	in := bufio.NewReader(os.Stdin)
	txt, err := in.ReadString('\n')
	
	if err != nil {
		panic("Something wrong with your text.")
	}
	
	createImage(txt)
}

func createImage(text string) {
	b := []byte(text)
	dim := len(b)/2
	
	img := image.NewRGBA(image.Rect(0, 0, dim, dim))
	
	//rows
	for i := 0; i < dim; i++ {
		//columns
		for j := 0; j < dim; j++ {
			byt := b[i+j]
			
			ran := rand.Intn(3)
			var r, g, b byte = 0, 0, 0
			switch ran {
				case 0:
					r = byt
				case 1:
					g = byt
				case 2:
					b = byt
				default:
					break
			}
			
			cl := color.NRGBA{R: r, G: g, B: b, A: 255}
			img.Set(i, j, cl)
		}
	}
	
	out, err := os.Create("Text.png")
	
	if err != nil {
		panic("Can not save image")
	}
	
	render := image.NewRGBA(image.Rect(0, 0, 640, 640))
	draw.NearestNeighbor.Scale(render, render.Rect, img, img.Bounds(), draw.Over, nil)
	
	png.Encode(out, render)
	out.Close()
}