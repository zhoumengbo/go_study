package main
import (
	"fmt"
	"sync"
	"time"
)

// 定时任务，直到某事发生

var done bool // 多个线程使用的共享变量
var mu sync.Mutex // 使得线程对 done 的访问互斥

func main(){
	time.Sleep(1 * time.Second)
	fmt.Println("started")
	go periodic()
	// 等一会儿，我们可以观察ticker的作用
	time.Sleep(5 * time.Second)
	mu.Lock()
	done = true
	mu.Unlock()
	fmt.Println("cancelled")
	// 观察无输出
	time.Sleep(3 * time.Second)
}

func periodic(){
	for{
		fmt.Println("tick")
		time.Sleep(1 * time.Second)
		mu.Lock()
		if done{
			mu.Unlock()
			return
		}
		mu.Unlock()
	}
}