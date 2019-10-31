package main

import "fmt"

//值传递与指针传递的区别

type Rect struct {
	X, Y, Area, Length float64
}

func (r *Rect) getArea() {
	r.Area = r.X * r.Y
}

func (r Rect) getLength() {
	r.Length = r.X + r.Y
}

func main() {
	r1 := Rect{
		X:      2,
		Y:      3,
		Area:   0,
		Length: 0,
	}
	fmt.Println("before")
	fmt.Printf("r1 area : %f", r1.Area)
	fmt.Printf("   r1 length : %f", r1.Length)

	r1.getLength()
	r1.getArea()

	fmt.Println("\nafter")
	fmt.Printf("r1 area : %f", r1.Area)
	fmt.Printf("  r1 length : %f", r1.Length)
}
