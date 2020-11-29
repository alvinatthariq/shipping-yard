package main

import "fmt"

type PerahuMotor struct {
	Name     string
	Capacity int
}

type PerahuLayar struct {
	Name     string
	Capacity int
}

type KapalPesiar struct {
	Name     string
	Capacity int
}

type Ship interface {
	getShipName() string
	getShipCapacity() int
}

func (p PerahuMotor) getShipName() string {
	return p.Name
}

func (p PerahuMotor) getShipCapacity() int {
	return p.Capacity
}

func (p PerahuLayar) getShipName() string {
	return p.Name
}

func (p PerahuLayar) getShipCapacity() int {
	return p.Capacity
}

func (p KapalPesiar) getShipName() string {
	return p.Name
}

func (p KapalPesiar) getShipCapacity() int {
	return p.Capacity
}

func main() {
	ship1 := KapalPesiar{
		Name:     "Titanic",
		Capacity: 1000,
	}

	ship2 := PerahuMotor{
		Name:     "Queen",
		Capacity: 20,
	}

	ship3 := PerahuLayar{
		Name:     "Sans",
		Capacity: 50,
	}

	ships := []Ship{ship1, ship2, ship3}
	getTotalCapacity(ships)
}

func getTotalCapacity(ships []Ship) {
	var totalCapcity int
	for _, ship := range ships {
		fmt.Printf("%s Capacity : %d\n", ship.getShipName(), ship.getShipCapacity())
		totalCapcity += ship.getShipCapacity()
	}

	fmt.Println("Total all ships capacity :", totalCapcity)
}
