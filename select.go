package main

import "fmt"

func main() {
	ch1 := make(chan int)
	go func() {
		sum := 0
		for i := 1; i < 1000; i++ {
			sum += i
			ch1 <- sum
		}
		close(ch1) //关闭管道
	}()

	ch2 := make(chan int)
	go func() {
		for i := 1; i < 1000; i++ {
			ch2 <- i + 2
		}
		close(ch2) //关闭管道
	}()

	// for {				//死循环方式读取管道值
	// 	value,ok:=<-ch1
	// 	if ok{
	// 		fmt.Println(" for ch1=",value)
	// 	}else{
	// 		break
	// 	}
	// }

	// for c:=range ch1{		//较为推荐的for-range读取,直到管道关闭就不再接收。
	// 	fmt.Println("range cha1=",c)
	// }

	//如果定义多个管道，那么以上两种就有点吃力了，引入select方法
	for {
		select { //必须为IO操作，要么读要么写。当多个分支条件满足时，会随机选择一个分支执行，本题需要持续接收值，所以需要加for
		case value, ok := <-ch1:
			if ok {
				fmt.Println("select ch1=", value)
			}
		case value, ok := <-ch2:
			if ok {
				fmt.Println("select ch2=", value)
			}
		}
	}

	fmt.Println("main 结束")

}
