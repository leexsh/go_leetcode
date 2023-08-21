package context

import (
	"context"
	"fmt"
	"time"
)

func EntranceValueContext() {
	ctx := context.WithValue(context.Background(), "parameter", "1")
	go handleValueRequest(ctx)
	time.Sleep(5 * time.Second)
	fmt.Println("It's time to stop all sub goroutines")
}

func handleValueRequest(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandleRequest Done")
			return
		default:
			fmt.Println("HandelRequest running, parameter:", ctx.Value("parameter"))
			time.Sleep(time.Second * 2)
		}
	}
}
