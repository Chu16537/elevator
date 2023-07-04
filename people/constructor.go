package people

import (
	"elevator/dao"
	"time"
)

type People struct {
	name          string               // 名稱
	startLevel    int                  // 起始樓層
	toLevel       int                  // 目標樓層
	clickWaitTime time.Duration        // 等待點擊時間
	clcikChan     chan dao.PeopleClick // 傳送點擊資料給manager
}

func New(name string, clcikChan chan dao.PeopleClick) *People {

	p := new(People)
	p.name = name
	p.clcikChan = clcikChan
	p.setStartLevel()
	p.setToLevel()
	p.setClickWaitTime()

	return p
}
