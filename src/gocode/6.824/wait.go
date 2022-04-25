package main

import (
	"fmt"
)

// 循环发送 RPC （并行）  使用 channel 代替 waitGroup

func main(){
	done := make(chan bool)
	for i := 0; i < 5; i++{
		go func(x int){
			sengRPC(x)
			done <- true
		}(i)
	}
	for i := 0; i < 5; i++ {
		<- done
	}
}

func sengRPC(i int){
	fmt.Println(i)
}