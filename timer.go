package main

import (
	"fmt"
	"time"
)

func main() {
	
	go func ()  {
		for{
			fmt.Println("我是协程...")
			time.Sleep(time.Second)
		}	
	}()
	<-time.After(time.Second*2)
	fmt.Println("main结束。。。")
}