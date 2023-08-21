package context

import (
	"context"
	"fmt"
	"time"
)

func EntranceCancelContext() {
	ctx, cancel := context.WithCancel(context.Background())
	go handleCancelRequest(ctx)
	time.Sleep(5 * time.Second)
	fmt.Println("It's time to stop all sub goroutines")
	cancel()
	time.Sleep(5 * time.Second)
}

func handleCancelRequest(ctx context.Context) {
	go writeRedis(ctx)
	go writeDB(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandleRequest Done")
			return
		default:
			fmt.Println("HandleRequest running")
			time.Sleep(time.Second * 2)
		}
	}
}

func writeRedis(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Wirte Redis Done")
			return
		default:
			fmt.Println("Write Redis Running")
			time.Sleep(1 * time.Second)
		}
	}
}

func writeDB(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Wirte Database Done")
			return
		default:
			fmt.Println("Write Database Running")
			time.Sleep(2 * time.Second)
		}
	}
}
