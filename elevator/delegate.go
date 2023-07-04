package elevator

import (
	"elevator/dao"
	"time"
)

// 電梯啟動
func (e *elevator) Start() {

	tick := time.NewTicker(100 * time.Millisecond)
	defer tick.Stop()

	for {
		select {
		case <-tick.C:
			e.action()
		}
	}
}

// 計算使用者點擊後 電梯到達時間
func (e *elevator) GetToPeopleLevelTime(p dao.PeopleClick) int {
	// 假如電梯剛好到與用者此樓層 增加一次等待時間
	// 移動時間
	// 開關門時間
	// 回傳時間
	return 0
}

// 添加進入人員資料
func (e *elevator) JoinPeople(p dao.PeopleClick) {
	e.joinMap[p.ToLevel] = append(e.joinMap[p.ToLevel], p)
}
