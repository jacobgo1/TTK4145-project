package main

import (
	"TTK4145-Project/Config"
	"TTK4145-Project/Elevator"
	"TTK4145-Project/Network"
)

func main() {
	// This is the main function that will be run when you execute your program.
	config.ID = "1"
	// Initialize the elevator
	elevator.InitElevator()

	orders := make(chan [config.NumFloors][config.NumButtons]config.OrderState)

	// Run the elevator
	go elevator.RunElevator(orders)

	go network.Network()

	select {}

}
