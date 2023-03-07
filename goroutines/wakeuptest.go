package goroutines

import (
	"fmt"
	"sort"
	"time"
)

func TestGoroutineWakeUp() {
	
	// wg := new(sync.WaitGroup)
	dataChan := make(chan time.Time, 1024)
    errCh := make(chan int)
	allStats := make([]time.Duration, 0, 100)

	go func() {
		for {
			select {
			case begin, ok := <-dataChan:
				if !ok {
					return
				}
				allStats = append(allStats,time.Since(begin))
			case <- errCh:
				return
			}
		}
	}()

	go func() {
		for i:=0; i< 100; i++ {
			dataChan <- time.Now()
		}
		errCh <- 0
	}()

	sort.Slice(allStats, func(i, j int) bool {
		return i<=j
	})

	sum := int64(0)
	for _, t := range allStats {
		sum += t.Microseconds()
	}
	fmt.Println("test times:", len(allStats))
	avg := float64(sum)/float64(len(allStats))
	fmt.Println("avg:",avg, "p80:", allStats[79], "p99:", allStats[98])
}