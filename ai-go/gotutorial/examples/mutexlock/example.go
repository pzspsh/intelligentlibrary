package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Golang中锁有两种，互斥锁Mutex和读写互斥锁RWMutex，互斥锁也叫读锁，读写锁也叫读锁，相互之间的关系为：
写锁需要阻塞写锁：一个协程拥有写锁时，其他协程写锁定需要阻塞
写锁需要阻塞读锁：一个协程拥有写锁时，其他协程读锁定需要阻塞
读锁需要阻塞写锁：一个协程拥有读锁时，其他协程写锁定需要阻塞
读锁不能阻塞读锁：一个协程拥有读锁时，其他协程也可以拥有读锁

使用：互斥锁和读写锁在使用上没有很大区别
1、互斥锁使用Lock()进行加锁，使用Unlock()进行解锁
2、读写锁使用RLock()加读锁，使用RUnlock()进行解读锁；使用Lock()加写锁，使用Unlock解写锁，和互斥锁功能一致；

但两者使用场景不同：
1、互斥锁会将操作串行化，可以保证操作完全有序，适合资源只能由一个协程进行操作的情况，并发能力弱；
2、读写锁适合读多写少的情况，并发能有比较强。
*/
func main() {
	wg := sync.WaitGroup{}

	var mutex sync.Mutex
	fmt.Println("G0抢占中...")
	mutex.Lock()
	fmt.Println("G0已抢占.")
	wg.Add(3)

	for i := 1; i < 4; i++ {
		go func(i int) {
			fmt.Printf("G%d抢占中...\n", i)
			mutex.Lock()
			fmt.Printf("G%d已抢占.\n", i)

			time.Sleep(time.Second * 2)
			mutex.Unlock()
			fmt.Printf("G%d已释放.\n", i)
			wg.Done()
		}(i)
	}
	time.Sleep(time.Second * 5)
	fmt.Println("G0准备释放.")
	mutex.Unlock()
	fmt.Println("G0已释放.")
	wg.Wait()
}
