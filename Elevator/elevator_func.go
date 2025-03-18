package elevator

import (
	"TTK4145-Project/Config"
	"TTK4145-Project/elevio"
)

func InitElevator() {
	orders := [config.NumFloors][config.NumButtons]config.OrderState{}
	for i := 0; i < config.NumFloors; i++ {
		for j := 0; j < config.NumButtons; j++ {
			orders[i][j] = config.Uninitialized
		}
	}
	ElevatorInstance = config.Elevator{
		ID:        config.ID,
		State:     config.Idle,
		Direction: elevio.MD_Stop,
		Floor:     0,
		Orders:    orders,
	}
}

func updateElevator(state chan config.ElevatorState, dir chan elevio.MotorDirection, floor chan int, orders chan [config.NumFloors][config.NumButtons]config.OrderState) {
	for {
		select {
		case f := <-floor:
			ElevatorInstance.Floor = f
		case s := <-state:
			ElevatorInstance.State = s
		case d := <-dir:
			ElevatorInstance.Direction = d
		case ord := <-orders:
			ElevatorInstance.Orders = ord
		}

	}

}
