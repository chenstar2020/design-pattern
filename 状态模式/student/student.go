package student

import "fmt"

type AbstractState interface {
	checkState(int)
	setState(sc *ScoreContext)
}

// ScoreContext 环境类
type ScoreContext struct {
	state AbstractState
}

func (d *ScoreContext) checkState(score int) {
	d.state = &HighState{}
	d.state.checkState(score)
}

func (d *ScoreContext) setState() {
	d.state.setState(d)
}

// HighState 优秀状态
type HighState struct{
	score int
}

func (h *HighState) checkState(score int){
	fmt.Println("加上：", score,  "当前分数:", h.score + score, "状态：High")
	h.score += score
}

func (h *HighState) setState(sc *ScoreContext){
	if h.score < 60 {
		sc.state = &LowState{}
	}else if h.score >= 90{
		sc.state = &HighState{}
	}else{
		sc.state = &MiddleState{}
	}
}

// MiddleState 中等状态
type MiddleState struct{
	score int
}

func (m *MiddleState) checkState(score int){
	fmt.Println("加上：", score,  "当前分数:", m.score + score, "状态：Middle")
	m.score += score
}

func (m *MiddleState) setState(sc *ScoreContext){
	if m.score < 60 {
		sc.state = &LowState{}
	}else if m.score >= 90{
		sc.state = &HighState{}
	}else{
		sc.state = &MiddleState{}
	}
}

var _ AbstractState = (*MiddleState)(nil)

// LowState 不及格状态
type LowState struct {
	score int
}

func (l *LowState) checkState(score int){
	fmt.Println("加上：", score,  "当前分数:", l.score + score, "状态：Low")
	l.score += score
}

func (l *LowState) setState(sc *ScoreContext){
	if l.score < 60 {
		sc.state = &LowState{}
	}else if l.score >= 90{
		sc.state = &HighState{}
	}else{
		sc.state = &MiddleState{}
	}
}
