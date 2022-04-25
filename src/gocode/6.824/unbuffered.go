package main
import (
	"fmt"
	"time"
)

// channels 实现

func main(){
	c := make(chan bool)
	go func(){
		time.Sleep(1 * time.Second)
		<-c // 从 channel 中读取数据
	}()
	start := time.Now()
	c <- true // 写数据到 channel 后阻塞，直到另一个 goroutine 接收到 channel 中的数据
	fmt.Printf("send took %v\n", time.Since(start))
}