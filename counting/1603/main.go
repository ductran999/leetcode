package main

import "fmt"

type ParkingSystem struct {
	placed [4]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	place := [4]int{
		0, big, medium, small,
	}
	p := ParkingSystem{
		placed: place,
	}
	return p
}

func (p *ParkingSystem) AddCar(carType int) bool {
	if p.placed[carType] == 0 {
		return false
	}
	p.placed[carType]--
	return true
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */

func main() {
	obj := Constructor(1, 1, 0)
	fmt.Println(obj.AddCar(1))
	fmt.Println(obj.AddCar(2))
	fmt.Println(obj.AddCar(3))
	fmt.Println(obj.AddCar(1))
}
