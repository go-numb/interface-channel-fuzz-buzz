package main

import (
	"fmt"
	"time"

	"github.com/go-numb/interface-channel-fuzz-buzz/workers"

	"github.com/labstack/gommon/log"
	"golang.org/x/sync/errgroup"
)

type Client struct {
	Results chan interface{}
}

func New() *Client {
	return &Client{
		Results: make(chan interface{}),
	}
}

func main() {
	go getChannels()
	structWorkers()
}

func structWorkers() {
	eachWorkers := map[string]workers.Worker{
		"salary0": workers.NewHuman("鈴木", "営業周り"),
		"salary1": workers.NewHuman("高橋", "資料整理"),
		"salary2": workers.NewHuman("後藤", "プレゼン資料作り"),
		"salary3": workers.NewMonky("猿如", "ネットスケープ"),
	}

	fmt.Println("就業開始")

	for _, worker := range eachWorkers {
		go func(worker workers.Worker) {
			if !worker.IsThere() {
				return
			}

			worker.Do()
			time.Sleep(2 * time.Second)
			worker.Say()
			time.Sleep(2 * time.Second)
			for {
				if 100 < worker.ProgressRate() {
					break
				}
				worker.Work()
				time.Sleep(2 * time.Second)
			}
			worker.GoHome()
		}(worker)
	}

	time.Sleep(30 * time.Second)
	fmt.Println("終業、はいここまで本日時間切れ")
}

func getChannels() {
	client := New()

	var eg errgroup.Group

	eg.Go(func() error { // 送信チャンネル
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				client.Results <- workers.NewHuman("鈴木", "営業周り")
			}
		}
	})

	eg.Go(func() error { // 受信チャンネル
		for {
			select {
			case result := <-client.Results:
				// channelからInterfaceを満たすstructがくれば、型チェックして型による動的な処理が可能
				worker, ok := result.(workers.Worker)
				if !ok {
					continue
				}
				worker.Do()
				worker.Say()
				worker.Work()
			}
		}
	})

	if err := eg.Wait(); err != nil {
		log.Error(err)
	}
}
