package main;

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

// Конструктор Point
func NewPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}
// Функция вычисления расстояния
func Distance(point1 *Point, point2 *Point) float64 {
	dx := point2.x - point1.x;
	dy := point2.y - point1.y;

	return math.Sqrt(dx * dx + dy * dy);
}

func main() {
	// Иниц. точек
	pa := NewPoint(0.0, 4.2);
	pb := NewPoint(3.2, 6.8);

	fmt.Printf("%.2f\n", Distance(pa, pb));
}