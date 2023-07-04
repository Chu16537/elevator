package elevatormanager

import "elevator/dao"

// 取得 按鈕chan
func (e *Elevatormanager) GetClickChan() chan dao.PeopleClick {
	return e.peopleClickChan
}

// 電梯啟動
func (e *Elevatormanager) Start() {

	for _, v := range e.elevators {
		v.Start()
	}
}
