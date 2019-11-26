package workers

import (
	"fmt"
	"math/rand"
	"time"
)

type Human struct {
	Name     string
	Job      string
	Progress int
}

func NewHuman(name, job string) *Human {
	return &Human{
		Name:     name,
		Job:      job,
		Progress: 0,
	}
}

func (p *Human) IsThere() bool {
	return true
}

func (p *Human) Do() {
	p.Progress += 1
	fmt.Printf("はい、%sです。%sに取り掛かり、終わり次第直帰します。\n", p.Name, p.Job)
}

func (p *Human) Say() {
	p.Progress -= 1
	fmt.Printf("こちら%s、現在仕事中です。\n", p.Name)
}

func (p *Human) Work() {
	rand.Seed(time.Now().UnixNano())
	add := rand.Intn(100)
	p.Progress += add
	if p.Progress < 100 {
		fmt.Printf("こちら%s、現在進捗 %d％です。\n", p.Name, p.Progress)
		return
	}
	fmt.Printf("こちら%s、現在進捗 100％突破しました。\n", p.Name)
}

func (p *Human) ProgressRate() int {
	return p.Progress
}

func (p *Human) GoHome() {
	fmt.Printf("%s、仕事をすべて終えました。帰宅します。\n", p.Name)
}
