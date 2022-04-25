package main

import (
	"fmt"
	"sync"
)

// 循环发送 RPC （并行）  使用 WaitGroup

func main(){
	var wg sync.WaitGroup
	for i := 0; i < 5; i++{
		wg.Add(1)
		go func(x int){
			sengRPC(x)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func sengRPC(i int){
	fmt.Println(i)
}