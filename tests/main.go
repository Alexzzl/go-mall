package main

import "fmt"

type FF func(int, int)
type A interface {
	F(int, int)
}
type St struct {
	Mem   string
	Point FF
}

func de_func(int, int) {
	fmt.Println("success")
}
func (s *St) init_in() {
	s.Point = de_func
}

func main() {
	var s St
	s.init_in()
	c := s.Point
	c(1, 2)
}
