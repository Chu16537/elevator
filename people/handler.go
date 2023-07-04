package people

import (
	"elevator/constant"
	"elevator/dao"
	"math/rand"
	"time"
)

/*
設定起始樓層
*/
func (p *People) setStartLevel() {
	p.startLevel = rand.Intn(constant.Max_Level) + 1
}

/*
設定目標樓層
*/
func (p *People) setToLevel() {

	var levels []int

	for i := constant.Min_Level; i <= constant.Max_Level; i++ {
		levels = append(levels, i)
	}

	// 隨機排序
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(levels); i++ {
		j := rand.Intn(i + 1)
		levels[i], levels[j] = levels[j], levels[i]
	}

	for _, v := range levels {
		// 如果目標樓層不是開始樓層
		if v != p.startLevel {
			p.toLevel = v
			return
		}
	}
}

/*
設定等待點擊按鈕時間
*/
func (p *People) setClickWaitTime() {
	rand.Seed(time.Now().UnixNano())

	p.clickWaitTime = time.Duration(rand.Intn(5)) * time.Second
}

// 點擊電梯
func (p *People) Click() {
	time.Sleep(p.clickWaitTime)

	req := dao.PeopleClick{
		Name:     p.name,
		NowLevel: p.startLevel,
		ToLevel:  p.toLevel,
	}

	p.clcikChan <- req
}
