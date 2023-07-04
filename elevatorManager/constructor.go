package elevatormanager

import (
	"elevator/dao"
	"elevator/elevator"
)

type Elevatormanager struct {
	elevators       []elevator.IElevator
	peopleClickChan chan dao.PeopleClick
}

/*
elevatorCount 電梯數量
*/
func New(elevatorCount int) *Elevatormanager {
	e := new(Elevatormanager)

	e.elevators = make([]elevator.IElevator, elevatorCount)
	for i := 0; i < elevatorCount; i++ {
		e.elevators[i] = elevator.New()
	}

	e.peopleClickChan = make(chan dao.PeopleClick, 100)

	go e.enable()
	return e
}
