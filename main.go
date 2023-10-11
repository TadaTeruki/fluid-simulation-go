package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	//"os"
	//"time"
	//"image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	cellWidth  = 320
	cellHeight = 320
	cellSize   = 2
)

func (g *CellMap) Draw(screen *ebiten.Image) {
	img := image.NewRGBA(image.Rect(0, 0, cellWidth, cellHeight))

	for iy := range g.colorMap {
		for ix := range g.colorMap[iy] {
			i := ix + iy*cellWidth
			img.Pix[4*i] = uint8(g.colorMap[iy][ix].r * 255)
			img.Pix[4*i+1] = uint8(g.colorMap[iy][ix].g * 255)
			img.Pix[4*i+2] = uint8(g.colorMap[iy][ix].b * 255)
			img.Pix[4*i+3] = 0xff
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.imgQueue = append(g.imgQueue, img)
	} else {
		// draw white rectangle at (20x20)
		for iy := 0; iy < 20; iy++ {
			for ix := 0; ix < 20; ix++ {
				i := ix + iy*cellWidth
				img.Pix[4*i] = 0xff
				img.Pix[4*i+1] = 0xff
				img.Pix[4*i+2] = 0xff
				img.Pix[4*i+3] = 0xff
			}
		}
	}

	screen.WritePixels(img.Pix)

	//g.imgQueue = append(g.imgQueue, img)
}

func (g *CellMap) Layout(outsideWidth, outsideHeight int) (int, int) {
	return cellWidth, cellHeight
}

func main() {
	ebiten.SetWindowSize(cellWidth*cellSize, cellHeight*cellSize)
	ebiten.SetWindowTitle("Fluid Simulation")
	g := NewCellMap()
	g.Init()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

	// write images to file
	for i, img := range g.imgQueue {
		f, _ := os.Create(fmt.Sprintf("images/%03d.png", i))
		png.Encode(f, img)
		f.Close()
	}

}
