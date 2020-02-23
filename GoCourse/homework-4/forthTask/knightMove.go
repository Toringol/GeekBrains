package main

import (
	"errors"
	"fmt"
)

type Point struct {
	x int
	y int
}

func (p Point) Get() (x, y int) {
	return p.x, p.y
}

func (p *Point) Set(x, y int) error {
	if x > 7 || x < 0 || y > 7 || y < 0 {
		return errors.New("No such point on board")
	}
	p.x = x
	p.y = y
	return nil
}

type AvailableKnightMoves []Point

func (av *AvailableKnightMoves) FindKnightMoves(startPoint Point) {
	xStart, yStart := startPoint.Get()
	if err := startPoint.Set(xStart+1, yStart+2); err == nil {
		tmpPoint := Point{xStart + 1, yStart + 2}
		*av = append(*av, tmpPoint)
	}
	if err := startPoint.Set(xStart-1, yStart+2); err == nil {
		tmpPoint := Point{xStart - 1, yStart + 2}
		*av = append(*av, tmpPoint)
	}
	if err := startPoint.Set(xStart-2, yStart+1); err == nil {
		tmpPoint := Point{xStart - 2, yStart + 1}
		*av = append(*av, tmpPoint)
	}
	if err := startPoint.Set(xStart+2, yStart+1); err == nil {
		tmpPoint := Point{xStart + 2, yStart + 1}
		*av = append(*av, tmpPoint)
	}
	if err := startPoint.Set(xStart-2, yStart-1); err == nil {
		tmpPoint := Point{xStart - 2, yStart - 1}
		*av = append(*av, tmpPoint)
	}
	if err := startPoint.Set(xStart-1, yStart-2); err == nil {
		tmpPoint := Point{xStart - 1, yStart - 2}
		*av = append(*av, tmpPoint)
	}
	if err := startPoint.Set(xStart+1, yStart-2); err == nil {
		tmpPoint := Point{xStart + 1, yStart - 2}
		*av = append(*av, tmpPoint)
	}
	if err := startPoint.Set(xStart+2, yStart-1); err == nil {
		tmpPoint := Point{xStart + 2, yStart - 1}
		*av = append(*av, tmpPoint)
	}
}

func (av AvailableKnightMoves) PrintKnightMoves() {
	for _, elem := range av {
		fmt.Println(elem.Get())
	}
}

func main() {
	availableKnightMoves := AvailableKnightMoves{}
	startPoint := Point{0, 0}
	availableKnightMoves.FindKnightMoves(startPoint)
	availableKnightMoves.PrintKnightMoves()
}
