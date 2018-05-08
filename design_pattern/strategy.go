/**
User:       wliangde
CreateTime: 2018/5/8 下午8:38
@desc:
策略模式
 Strategy 策略模式：
        它定义了算法家族，分别封装起来，让它们可以相互替换，
		此模式让算法的变化，不会影响到使用算法的客户。
 个人想法：UML图很相似，策略模式是用在对多个做同样事情（统一接口）的类对象的选择上，
         而状态模式是：将对某个事情的处理过程抽象成接口和实现类的形式，
		由context保存一份state，在state实现类处理事情时，修改状态传递给context，
		由context继续传递到下一个状态处理中

说白了就是定义一个接口暴露算法的api，
然后在实现里不同的算法做不同的处理。
这样不会影响到算法的使用者。

**/
package design_pattern

import (
	"fmt"
	"math/rand"
)

/**
例子：
猜拳策略，有两个策略
1、下次出拳出的是上次赢得那个
2、从上一次的出拳，概率算下次的出拳
*/

//手势
type THandType int

const (
	THandType_SHITOU  = 0
	THandType_JIANDAO = 1
	THandType_BU      = 2
)

type THand struct {
	nValue int
}

func NewHand(nValue int) *THand {
	return &THand{
		nValue: nValue,
	}
}

func (this *THand) isStrongerThan(ptOther *THand) bool {
	return this.fight(ptOther) == 1
}

func (this *THand) isWeakThan(ptOther *THand) bool {
	return this.fight(ptOther) == -1
}

func (this *THand) fight(ptOther *THand) int {
	if this.nValue == ptOther.nValue {
		return 0
	}

	if (this.nValue+1)%3 == ptOther.nValue {
		return 1
	} else {
		return -1
	}
}

func (this *THand) String() string {
	if this.nValue == 0 {
		return "石头"
	} else if this.nValue == 1 {
		return "剪刀"
	} else {
		return "布"
	}
}

type IStrategy interface {
	GetName() string
	NextHand() *THand
	Study(bWin bool)
}

type TWinningStrategy struct {
	bWon      bool
	ptPreHand *THand
	ptRand    *rand.Rand
}

func NewWinningStrategy(ldwSeed int64) *TWinningStrategy {
	return &TWinningStrategy{
		ptRand: rand.New(rand.NewSource(ldwSeed)),
	}
}

func (this *TWinningStrategy) GetName() string {
	return "TWinningStrategy"
}

func (this *TWinningStrategy) NextHand() *THand {
	if !this.bWon {
		nHandValue := this.ptRand.Intn(3)
		this.ptPreHand = NewHand(nHandValue)
	}

	return this.ptPreHand
}

func (this *TWinningStrategy) Study(bWin bool) {
	this.bWon = bWin
}

/**
TProbStrategy
*/
type TProbStrategy struct {
	aryRateTable [3][3]int
	ptRand       *rand.Rand
	ptPreHand    *THand
	ptCurHand    *THand
}

func NewProbStrategy(ldwSeed int64) *TProbStrategy {
	return &TProbStrategy{
		ptRand: rand.New(rand.NewSource(ldwSeed)),
		aryRateTable: [3][3]int{
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		},
	}
}

func (this *TProbStrategy) GetName() string {
	return "TProbStrategy"
}

func (this *TProbStrategy) NextHand() *THand {
	if this.ptPreHand == nil {
		nHandValue := this.ptRand.Intn(3)
		this.ptPreHand = NewHand(nHandValue)
		this.ptCurHand = this.ptPreHand
		return this.ptCurHand
	}

	this.ptCurHand = this.getHand(this.ptPreHand.nValue)

	return this.ptCurHand
}

func (this *TProbStrategy) getHand(nValue int) *THand {
	tempAry := this.aryRateTable[nValue]

	nTotal := 0
	for v := range tempAry {
		nTotal += v
	}

	nRand := this.ptRand.Intn(nTotal)

	nOffset := 0
	for k, v := range tempAry {
		if nRand < v+nOffset {
			return NewHand(k)
		}
		nOffset += v
	}

	nValue = this.ptRand.Intn(3)
	return NewHand(nValue)

}

func (this *TProbStrategy) Study(bWin bool) {
	if bWin {
		this.aryRateTable[this.ptPreHand.nValue][this.ptCurHand.nValue]++
	} else {
		this.aryRateTable[this.ptPreHand.nValue][(this.ptCurHand.nValue+1)%3]++
		this.aryRateTable[this.ptPreHand.nValue][(this.ptCurHand.nValue+2)%3]++
	}
}

/**
TPlayer
*/
type TPlayer struct {
	strName   string
	iStrategy IStrategy
	nWinCnt   int
	nLoseCnt  int
	nGameCnt  int
}

func NewPlayer(strName string, i IStrategy) *TPlayer {
	return &TPlayer{
		strName:   strName,
		iStrategy: i,
	}
}

func (this *TPlayer) NextHand() *THand {
	return this.iStrategy.NextHand()
}

func (this *TPlayer) Win() {
	this.iStrategy.Study(true)
	this.nWinCnt++
	this.nGameCnt++

}

func (this *TPlayer) Lose() {
	this.iStrategy.Study(false)
	this.nLoseCnt++
	this.nGameCnt++
}

//平手
func (this *TPlayer) Even() {
	this.nGameCnt++
}

func (this *TPlayer) Report() string {
	return fmt.Sprintf("[name:%s 战斗:%d 赢了:%d 输了:%d]\n", this.strName, this.nGameCnt, this.nWinCnt, this.nLoseCnt)
}
