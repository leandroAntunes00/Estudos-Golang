package main

import "fmt"


func main() {
	s := [] int {10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])
	fmt.Printf("len=%d cap=%d %v\n", len(s[:1]), cap(s[:1]), s[:1])
	fmt.Printf("len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s[:2])
	fmt.Printf("len=%d cap=%d %v\n", len(s[:3]), cap(s[:3]), s[:3])
	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])
	fmt.Printf("len=%d cap=%d %v\n", len(s[:5]), cap(s[:5]), s[:5])
	fmt.Printf("len=%d cap=%d %v\n", len(s[:6]), cap(s[:6]), s[:6])

}
