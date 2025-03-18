package config

import (
	"TTK4145-Project/elevio"
)

var (
	ID   string
	Port string
)

const (
	NumFloors    = 4
	NumButtons   = 3
	NumElevators = 3
)

type OrderState int

const (
	NoOrder OrderState = iota
	Unconfirmed
	Confirmed
	Uninitialized = -1
)

type ElevatorState int

const (
	Idle ElevatorState = iota
	Moving
	DoorOpen
)

type Elevator struct {
	ID        string
	State     ElevatorState
	Direction elevio.MotorDirection
	Floor     int
	Orders    [NumFloors][NumButtons]OrderState
}
