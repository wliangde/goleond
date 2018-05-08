/**
@user:          liangde.wld
@createtime:    2018/5/8 下午11:08
@desc:
**/
package design_pattern

import (
	"testing"
	"time"
)

func TestStrategy(t *testing.T) {

	ptWinStrategy := NewWinningStrategy(time.Now().Unix())
	ptProbStrategy := NewProbStrategy(time.Now().Unix() * 2)

	ptPlayer1 := NewPlayer("wld", ptWinStrategy)
	ptPlayer2 := NewPlayer("dxl", ptProbStrategy)

	for i := 0; i < 100; i++ {
		ptHand1 := ptPlayer1.NextHand()
		ptHand2 := ptPlayer2.NextHand()
		t.Log("局数", i, ptHand1.String(), ptHand2.String())
		if ptHand1.isStrongerThan(ptHand2) {
			ptPlayer1.Win()
			ptPlayer2.Lose()
		} else if ptHand1.isWeakThan(ptHand2) {
			ptPlayer2.Win()
			ptPlayer1.Lose()
		} else {
			ptPlayer1.Even()
			ptPlayer2.Even()
		}
	}
	t.Log(ptPlayer1.Report())
	t.Log(ptPlayer2.Report())
}
