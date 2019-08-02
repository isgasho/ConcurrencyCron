package main

import (
	"ConcurrencyCron"
	"context"
	"fmt"
	"time"
)

/**
 *@author  wxn
 *@project ConcurrencyCron
 *@package ConcurrencyCron
 *@date    19-8-1 下午5:58
 */
func test(num int) {
	fmt.Println("before:im a task", num)
	time.Sleep(10 * time.Second)
	fmt.Println("after:im a task", num, time.Now())
}

func main() {
	scheduler, err := ConcurrencyCron.NewScheduler(200)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 200; i++ {
		scheduler.Every(1).Seconds().Do(test, i)
	}
	ctx, cancel := context.WithCancel(context.Background())
	scheduler.Start(ctx)
	ch := make(chan bool)
	<-ch
	defer cancel()

}
