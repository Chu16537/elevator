package elevator

import (
	"fmt"
	"time"
)

// 動作
func (e *elevator) action() {

	// 可動作時間未到達
	if e.nextActionTime.After(time.Now()) {
		return
	}

	// 開門
	e.openGate()

	// 移動
	e.move()
}

// 開關門
func (e *elevator) openGate() {
	fmt.Println("openGate start")
	defer fmt.Println("openGate end")
	// 判斷 joinMap leaveMap 是否有人要進出
	// 有 等待進出時間
	// 沒 不等待

	// e.leaveMap 移除 當前樓層出去者資料

	// 假如 e.joinMap 有資料判斷此樓層人員是否可以進入 方向不同暫時無法進入
	// 移除進入者資料

	// e.leaveMap 添加進入者趣網樓層資料

}

// 移動
func (e *elevator) move() {
	// 到 Max_Level 往下 到 Min_Level 往上
	fmt.Println("move start")
	defer fmt.Println("move end")
}
