package elevator

import (
	"elevator/constant"
	"elevator/dao"
	"time"
)

type IElevator interface {
	Start()
	GetToPeopleLevelTime(p dao.PeopleClick) int
	JoinPeople(p dao.PeopleClick)
}

type elevator struct {
	maxLevel       int                       // 最高樓層 大樓共10層樓
	nowLevel       int                       // 當前樓層
	maxCount       int                       // 最大人數 電梯只可容納5人
	nowCount       int                       // 當前人數
	moveTime       time.Duration             // 移動時間 每行經一層樓需耗時1秒
	stopTime       time.Duration             // 停止時間 每停一次處理接人放人需耗1秒
	moveMode       int                       // 移動模式
	joinMap        map[int][]dao.PeopleClick // 進入樓層map[停止樓層]進站人資料
	leaveMap       map[int]int               // 出去樓層map[停止樓層]出站人數
	nextActionTime time.Time                 // 下次動作時間
}

func New() IElevator {

	e := new(elevator)

	e.maxLevel = constant.Max_Level
	e.nowLevel = constant.Min_Level
	e.maxCount = constant.Max_Count
	e.nowCount = 0
	e.moveTime = 1 * time.Second
	e.stopTime = 1 * time.Second
	e.moveMode = constant.MoveMode_Stop
	e.joinMap = make(map[int][]dao.PeopleClick)
	e.leaveMap = make(map[int]int)

	return e
}
