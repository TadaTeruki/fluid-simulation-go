package main

func (g *CellMap) Init() {

	const vradius = 0.1

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

func (g *CellMap) Update() error {

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

			rho := 0.99
			pressureDiffX := (g.pressureMap[iy][wrap(ix+1, cellWidth)] - g.pressureMap[iy][wrap(ix-1, cellWidth)]) / 2
			pressureDiffY := (g.pressureMap[wrap(iy+1, cellHeight)][ix] - g.pressureMap[wrap(iy-1, cellHeight)][ix]) / 2
			fx := g.velocityMap[iy][ix].x - pressureDiffX/rho
			fy := g.velocityMap[iy][ix].y - pressureDiffY/rho
			newVelocityMap[iy][ix] = Velocity{fx, fy}

		}
	}
	g.colorMap = newColorMap
	g.velocityMap = newVelocityMap
	g.pressureMap = newPressureMap
	return nil
}
