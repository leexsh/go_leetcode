package context

import (
	"context"
	"fmt"
	"time"
)

func EntranceTimeoutContext() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	go handleTimeoutRequest(ctx)
	time.Sleep(10 * time.Second)
	fmt.Println("It's time to stop all sub goroutines")
}

func handleTimeoutRequest(ctx context.Context) {
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
