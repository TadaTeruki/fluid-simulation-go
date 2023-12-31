package main

import "image"

type Velocity struct {
	x float64
	y float64
}

func (v Velocity) lerp(v2 Velocity, t float64) Velocity {
	return Velocity{
		x: v.x*(1-t) + v2.x*t,
		y: v.y*(1-t) + v2.y*t,
	}
}

func (v *Velocity) add(v2 Velocity) {
	v.x += v2.x
	v.y += v2.y
}

func (v *Velocity) mul(s float64) {
	v.x *= s
	v.y *= s
}

type Color struct {
	r float64
	g float64
	b float64
}

func (c Color) lerp(c2 Color, t float64) Color {
	return Color{
		r: c.r*(1-t) + c2.r*t,
		g: c.g*(1-t) + c2.g*t,
		b: c.b*(1-t) + c2.b*t,
	}
}

type CellMap struct {
	colorMap    [][]Color
	velocityMap [][]Velocity
	pressureMap [][]float64
	cursorX     int
	cursorY     int
	imgQueue    []*image.RGBA
}

func NewCellMap() *CellMap {

	colorMap := make([][]Color, cellHeight)
	velocityMap := make([][]Velocity, cellHeight)
	pressureMap := make([][]float64, cellHeight)
	for iy := range colorMap {
		colorMap[iy] = make([]Color, cellWidth)
		velocityMap[iy] = make([]Velocity, cellWidth)
		pressureMap[iy] = make([]float64, cellWidth)
		for ix := range colorMap[iy] {
			colorMap[iy][ix] = Color{}
			velocityMap[iy][ix] = Velocity{}
			pressureMap[iy][ix] = 0.0
		}
	}

	return &CellMap{
		colorMap:    colorMap,
		velocityMap: velocityMap,
		pressureMap: pressureMap,
		cursorX:     -1,
		cursorY:     -1,
		imgQueue:    make([]*image.RGBA, 0),
	}
}
