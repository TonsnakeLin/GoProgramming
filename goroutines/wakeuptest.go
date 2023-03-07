package goroutines

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

func TestGoroutineWakeUp() {
	wg := new(sync.WaitGroup)
	dataChan := make(chan time.Time, 1024)
	allStats := make([]time.Duration, 0, 100)

	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			begin, ok := <-dataChan
			if !ok {
				return
			}
			allStats = append(allStats, time.Since(begin))
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			dataChan <- time.Now()
		}
		time.Sleep(time.Microsecond * 1)
		close(dataChan)
	}()

	wg.Wait()
	sort.Slice(allStats, func(i, j int) bool {
		return allStats[i] <= allStats[j]
	})
	// fmt.Println("allStats:", allStats)
	sum := int64(0)
	for _, t := range allStats {
		sum += t.Microseconds()
	}
	// fmt.Println("test times:", len(allStats))
	avg := float64(sum) / float64(len(allStats))
	fmt.Println("avg:", avg, "median:", allStats[49], "p99:", allStats[98], "pmax:", allStats[99])
}

func Test10times() {
	for i := 0; i < 10; i++ {
		TestGoroutineWakeUp()
	}
}
