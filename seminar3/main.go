package main

import (
	"fmt"
	"math"
)

type Shape struct {
    Name string
}

func (s Shape) GetName() string {
	return s.Name
}


func (s Shape) Area() float64 {
    return 0.0
}

type Rectangle struct {
    Shape
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

type Circle struct {
    Shape
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func main() {
    fmt.Println("Введите стороны прямоугольника в формате ширина высота через пробел и радиус круга: ")
    var Width float64
    var Height float64
    var Radius float64
    _, err := fmt.Scanf("%f %f %f", &Width, &Height, &Radius)	
    if err == nil {
        rect := Rectangle{Shape{"Прямоугольник"}, Width, Height}
        circ := Circle{Shape{"Круг"}, Radius}
        fmt.Printf("%s: Площадь = %.2f\n", rect.GetName(), rect.Area())
        fmt.Printf("%s: Площадь = %.2f\n", circ.GetName(), circ.Area())
    } else {
        fmt.Println("Введите числа")
    }
    fmt.Println("Расчет закончен.")
}