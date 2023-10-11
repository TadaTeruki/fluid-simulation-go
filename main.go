package main

import (
	"image"
	"log"

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

	screen.WritePixels(img.Pix)
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
}
