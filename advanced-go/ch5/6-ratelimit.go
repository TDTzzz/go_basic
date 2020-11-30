package main

import (
	"fmt"
	"time"
)

func main() {
	var fillInterval = time.Millisecond * 1000
	var capacity = 100
	var tokenBucket = make(chan struct{}, capacity)

	fillToken := func() {
		ticker := time.NewTicker(fillInterval)
		for {
			select {
			case <-ticker.C:
				select {
				case tokenBucket <- struct{}{}:
				default:
				}
				fmt.Println("current token cnt:", len(tokenBucket), time.Now())
			}
		}
	}

	go fillToken()
	time.Sleep(time.Hour)
}

//消耗桶令牌
//func TakeAvailable(block bool) bool {
//	var takenResult bool
//	if block {
//		select {
//		case <-tokenBucket:
//			takenResult = true
//		}
//	} else {
//		select {
//		case <-tokenBucket:
//			takenResult = true
//		default:
//			takenResult = false
//		}
//	}
//
//	return takenResult
//}
