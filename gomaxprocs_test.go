package goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	//meliaht jumlah CPU
	totalCpu := runtime.NumCPU()
	fmt.Println("total CPU", totalCpu) //4

	//melihat total Thread
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total Thread", totalThread) //4

	//melihat jumlah Goroutine
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total goroutine", totalGoroutine) //102

	group.Wait()
}

func TestChangeThreadNumber(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	//meliaht jumlah CPU
	totalCpu := runtime.NumCPU()
	fmt.Println("total CPU", totalCpu) //12

	//mengubah jumlah Thread
	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total Thread", totalThread) //20

	//melihat jumlah Goroutine
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total goroutine", totalGoroutine) //102

	group.Wait()
}
