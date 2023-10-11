package main

import "math"

type lerpables[L any] interface {
	lerp(l2 L, t float64) L
}

func wrap(n, width int) int {
	return (n + width) % width
}

func smallerIndex(n float64, width int) int {
	return int(math.Floor(n+float64(width))) % width
}

func largerIndex(n float64, width int) int {
	return int(math.Ceil(n+float64(width))) % width
}

func Sample[L lerpables[L]](m *[][]L, x, y float64) L {
	ix0, iy0 := smallerIndex(x, cellWidth), smallerIndex(y, cellHeight)
	ix1, iy1 := largerIndex(x, cellWidth), largerIndex(y, cellHeight)
	tx := x - math.Floor(x)
	ty := y - math.Floor(y)

	mixA := (*m)[iy0][ix0].lerp((*m)[iy0][ix1], tx)
	mixB := (*m)[iy1][ix0].lerp((*m)[iy1][ix1], tx)
	mix := mixA.lerp(mixB, ty)

	return mix
}

// set value to sampled area
func SetSampled[L any](m *[][]L, x, y float64, fn func(l1 L, weight float64) L) {
	ix0, iy0 := smallerIndex(x, cellWidth), smallerIndex(y, cellHeight)
	ix1, iy1 := largerIndex(x, cellWidth), largerIndex(y, cellHeight)
	tx := x - math.Floor(x)
	ty := y - math.Floor(y)

	weight00 := (1 - tx) * (1 - ty)
	weight01 := tx * (1 - ty)
	weight10 := (1 - tx) * ty
	weight11 := tx * ty

	(*m)[iy0][ix0] = fn((*m)[iy0][ix0], weight00)
	(*m)[iy0][ix1] = fn((*m)[iy0][ix0], weight01)
	(*m)[iy1][ix0] = fn((*m)[iy0][ix0], weight10)
	(*m)[iy1][ix1] = fn((*m)[iy0][ix0], weight11)
}
