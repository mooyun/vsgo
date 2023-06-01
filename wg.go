package main

import (
	"fmt"
	_ "time"
)
func main(){
	//c:=make(chan int ,1000)
	c:=make(chan bool ,1000)
		for i := 0; i < 1000; i++ {
		go func (i int)  {
		fmt.Println(i)
		c<-true

	}(i)	
		}
		for i := 0; i < 1000; i++ {
			<-c
		}
 
	}