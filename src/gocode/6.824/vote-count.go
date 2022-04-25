package main
import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

// raft选举leader demo

func main(){
	rand.Seed(time.Now().UnixNano())

	count := 0
	finished := 0

	// 当共享数据中的某个条件或属性变为true时，通过condition来进行协调
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	for i := 0; i < 10; i++{
		go func(){
			vote := requestVote()
			mu.Lock()
			defer mu.Unlock()
			if vote{
				count++
			}
			finished++
			// 每次获取投票信息后检查是否投票结束
			cond.Broadcast()
		}()
	}

	mu.Lock()
	for count < 5 && finished != 10{
		cond.Wait()
	}
	if count >= 5{
		fmt.Println("receive 5+ votes!")
	}else{
		fmt.Println("lost")
	}
	mu.Unlock()

}

func requestVote() bool {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	return rand.Int() % 2 == 0 
}