package main

import (
	//"fmt"
	"math"
	"github.com/hajimehoshi/ebiten/v2"
)


func (g *CellMap) Init() {

	const vradius = 0.

	for iy := range g.colorMap {
		for ix := range g.colorMap[iy] {
			flagX, flagY := 0, 0
			if ix*2 >= cellWidth {
				flagX = 1
			}
			if iy*2 >= cellHeight {
				flagY = 1
			}

			flag := flagX + flagY*2
			var velocity Velocity
			var color Color
			switch flag {
			case 0:
				velocity = Velocity{vradius, -vradius}
				color = Color{0.8, 0.0, 0.0}
			case 1:
				velocity = Velocity{vradius, vradius}
				color = Color{0.0, 0.8, 0.0}
			case 2:
				velocity = Velocity{-vradius, -vradius}
				color = Color{0.0, 0.0, 0.8}
			case 3:
				velocity = Velocity{-vradius, vradius}
				color = Color{0.8, 0.0, 0.0}
			}
			g.colorMap[iy][ix] = color
			g.velocityMap[iy][ix] = velocity
		}
	}
}

func distBetweenPointAndLine(x, y, x1, y1, x2, y2 float64) float64 {
    dx, dy := x2-x1, y2-y1
    px, py := x-x1, y-y1

	if (dx*dx + dy*dy) == 0 {
		return 1e16
	}

    t := (px*dx + py*dy) / (dx*dx + dy*dy)
    if t < 0 {
        return dist(x, y, x1, y1)
    }
    if t > 1 {
        return dist(x, y, x2, y2)
    }
    cx, cy := x1+t*dx, y1+t*dy
    return dist(x, y, cx, cy)
}

func dist(x1, y1, x2, y2 float64) float64 {
    return math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
}
func (g *CellMap) Update() error {

	effectX, effectY := 0.0, 0.0
	effectRadius := 20.0
	
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if g.cursorX != -1 && g.cursorY != -1 {
			effectX = float64(x - g.cursorX)
			effectY = float64(y - g.cursorY)
		}
		g.cursorX, g.cursorY = x, y
	} else {
		g.cursorX, g.cursorY = -1, -1
	}

	newColorMap := make([][]Color, cellHeight)
	newVelocityMap := make([][]Velocity, cellHeight)
	newPressureMap := make([][]float64, cellHeight)

	for iy := range newColorMap {
		newColorMap[iy] = make([]Color, cellWidth)
		newVelocityMap[iy] = make([]Velocity, cellWidth)
		newPressureMap[iy] = make([]float64, cellWidth)
	}

	for iy := range g.colorMap {
		for ix := range g.colorMap[iy] {
			sx := float64(ix) - g.velocityMap[iy][ix].x
			sy := float64(iy) - g.velocityMap[iy][ix].y

			newColorMap[iy][ix] = Sample(&g.colorMap, sx, sy)
			newVelocityMap[iy][ix] = Sample(&g.velocityMap, sx, sy)

			divergence := g.velocityMap[iy][wrap(ix-1, cellWidth)].x - g.velocityMap[iy][wrap(ix+1, cellWidth)].x +
				g.velocityMap[wrap(iy-1, cellHeight)][ix].y - g.velocityMap[wrap(iy+1, cellHeight)][ix].y 

			newPressureMap[iy][ix] = 0.25 *
				(divergence + g.pressureMap[iy][wrap(ix-1, cellWidth)] +
					g.pressureMap[iy][wrap(ix+1, cellWidth)] +
					g.pressureMap[wrap(iy-1, cellHeight)][ix] +
					g.pressureMap[wrap(iy+1, cellHeight)][ix])

			rho := 10.
			pressureDiffX := (g.pressureMap[iy][wrap(ix+1, cellWidth)] - g.pressureMap[iy][wrap(ix-1, cellWidth)]) / 2
			pressureDiffY := (g.pressureMap[wrap(iy+1, cellHeight)][ix] - g.pressureMap[wrap(iy-1, cellHeight)][ix]) / 2
			fx := g.velocityMap[iy][ix].x - pressureDiffX/rho
			fy := g.velocityMap[iy][ix].y - pressureDiffY/rho
			
			newVelocityMap[iy][ix] = Velocity{fx, fy}

			effectStrength := 0.0
			if g.cursorX != -1 && g.cursorY != -1 {
				
				dist := distBetweenPointAndLine(float64(ix), float64(iy), float64(g.cursorX), float64(g.cursorY), float64(g.cursorX+int(effectX)), float64(g.cursorY+int(effectY)))
				effectStrength = (effectRadius-dist)/effectRadius*0.2
				if effectStrength < 0.0 {
					effectStrength = 0.0
				}
			}
			newVelocityMap[iy][ix].add(Velocity{effectX * effectStrength, effectY * effectStrength})
			newVelocityMap[iy][ix].mul(0.99)
		}
	}


	g.colorMap = newColorMap
	g.velocityMap = newVelocityMap
	g.pressureMap = newPressureMap
	return nil
}
