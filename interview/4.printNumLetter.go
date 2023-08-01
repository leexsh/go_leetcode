package interview

import (
	"fmt"
	"sync"
)

/*
	交替打印字母和数字
*/

var (
	wg       = sync.WaitGroup{}
	chNum    = make(chan struct{}, 1)
	chLetter = make(chan struct{}, 1)
)

func printNum() {
	for i := 0; i < 10; i++ {
		<-chNum
		fmt.Println(i)
		chLetter <- struct{}{}
	}
	wg.Done()
}

func printLetter() {
	for i := 0; i < 10; i++ {
		<-chLetter
		fmt.Printf("%s\n", string('a'+i))
		chNum <- struct{}{}
	}
	wg.Done()
}

func do() {
	wg.Add(2)
	go printNum()
	go printLetter()
	chNum <- struct{}{}
	wg.Wait()
	close(chLetter)
	close(chNum)
}
