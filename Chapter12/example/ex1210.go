package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

type PressureGauge struct {
	ch chan struct{}
}

func New(limit int) *PressureGauge {
	return &PressureGauge{
		ch: make(chan struct{}, limit),
	}
}

func (pg *PressureGauge) Process(f func()) error {
	select {
	case pg.ch <- struct{}{}:
		f()
		<-pg.ch
		return nil
	default:
		return errors.New("キャパシティに余裕がありません。")
	}
}

func doThingThatShouldBeLimited() string {
	time.Sleep(2 * time.Second)
	return "完了"
}

func main() {
	pg := New(5)
	http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		err := pg.Process(func() {
			w.Write([]byte(doThingThatShouldBeLimited()))
		})

		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("リクエストが多すぎてさばききれません\n"))
		}
	})

	fmt.Println("ブラウザで次を開いてください：'http://localhost:8080/request")
	fmt.Println("あるいは'sh ex1210.shを実行してみてください。")
	http.ListenAndServe(":8080", nil)
}
