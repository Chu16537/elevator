package elevatormanager

import (
	"elevator/dao"
	"fmt"
)

func (e *Elevatormanager) enable() {

	for p := range e.peopleClickChan {
		e.distribute(p)
	}
}

/*
實作分配策略
*/
func (e *Elevatormanager) distribute(p dao.PeopleClick) {
	// 實作分配策略
	fmt.Println("distribute start")
	defer fmt.Println("distribute end")

	eMap := map[int]int{}

	//預設給第一台電梯
	tagElevator := -1

	// for 判斷 並給每台電梯時間 等待時間低者 分配使用者給此電梯
	for i, v := range e.elevators {

		// 取得 此電梯 移動到 使用者 時間
		waitTime := v.GetToPeopleLevelTime(p)

		eMap[i] = waitTime

		//　假如是第一台 直接預設
		if tagElevator == -1 {
			tagElevator = i
		}

		// 假如已經預設別台電梯 判斷等待時間是否比 tagElevator 還低
		if waitTime < eMap[tagElevator] {
			tagElevator = i
		}
	}

	// 分配給電梯
	e.elevators[tagElevator].JoinPeople(p)
}
