package workers

import (
	"fmt"
	"math/rand"
	"time"
)

type Monky struct {
	Name     string
	Job      string
	Progress int
}

func NewMonky(name, job string) *Monky {
	return &Monky{
		Name:     name,
		Job:      job,
		Progress: 0,
	}
}

func (p *Monky) IsThere() bool {
	return true
}

func (p *Monky) Do() {
	p.Progress += 1
	fmt.Printf("うききっ、%sってなに？どうするの？\n", p.Job)
}

func (p *Monky) Say() {
	p.Progress -= 10
	fmt.Println("ウキッ、休憩まだ？ご飯まだ？")
}

func (p *Monky) Work() {
	rand.Seed(time.Now().UnixNano())
	add := rand.Intn(10)
	p.Progress += add
	if p.Progress < 100 {
		fmt.Printf("うきっ？、どうすんねんこれ 進捗%d％\n", p.Progress)
		return
	}
	fmt.Println("進捗100％きたこれぇい")
}

func (p *Monky) ProgressRate() int {
	return p.Progress
}

func (p *Monky) GoHome() {
	fmt.Printf("%s、仕事終わりません、残業します...\n", p.Name)
}
