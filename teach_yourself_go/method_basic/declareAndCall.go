package main

//there are two lots of stuff in here: movies and then sphere mehthods

import (
	"fmt"
	"math"
	"strconv"
)

type Movie struct {
	Name   string
	Rating float64
}

type Sphere struct {
	Radius float64
}

func (s *Sphere) SurfaceArea() float64 {
	return float64(4) * math.Pi * (s.Radius * s.Radius)
}

func (s *Sphere) Volume() float64 {
	radiusCubed := s.Radius * s.Radius * s.Radius
	return (float64(4) / float64(3)) * math.Pi * radiusCubed

}

func (m *Movie) summary() string {
	//method takes struct and converts and displays it for a reason
	//can basically do what you want with the data and call this
	r := strconv.FormatFloat(m.Rating, 'f', 1, 64)
	return m.Name + ", " + r
}

func main() {
	//can also var m Movie to initalise it
	m := Movie{
		Name:   "Heathers",
		Rating: 3.2,
	}
	fmt.Println(m.summary())

	s := Sphere{
		Radius: 5,
	}
	fmt.Println(s.SurfaceArea())
	fmt.Println(s.Volume())

}
