package main

import (
	elevatormanager "elevator/elevatorManager"
	"elevator/people"
	"strconv"
)

func main() {

	// 創建電梯manager
	elevatorCount := 2
	manager := elevatormanager.New(elevatorCount)
	manager.Start()

	// 創建人數
	peopleCount := 40
	peoples := make([]people.People, peopleCount)

	for i := 0; i < peopleCount; i++ {
		name := "p-" + strconv.Itoa(i)
		peoples[i] = *people.New(name, manager.GetClickChan())
	}

	// 開始點擊電梯
	for _, v := range peoples {
		v.Click()
	}
}
