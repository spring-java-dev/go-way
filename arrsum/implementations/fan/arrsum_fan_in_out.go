package fan

import (
	"sync"
)

// ArrSum demonstrates a Producer-Consumer, fan-out / fan-in, Channel Coordination concurrency patterns.
// Tags: concurrency, channels, fan-in, fan-out, waitgroup, producer, consumer
func ArrSum(numbers ...[]int) int {
	if len(numbers) == 0 {
		return 0
	}

	var wg sync.WaitGroup

	intsCh := make(chan int)
	arrsCh := make(chan []int)

	go func() {
		defer close(arrsCh)
		for _, numberArr := range numbers {
			arrsCh <- numberArr
		}
	}()

	for arr := range arrsCh {
		wg.Add(1)
		a := arr
		go intProducer(&wg, a, intsCh)
	}

	go func() {
		wg.Wait()
		close(intsCh)
	}()

	sum := 0
	for value := range intsCh {
		sum += value
	}

	return sum
}

func intProducer(
	wg *sync.WaitGroup,
	arr []int,
	output chan<- int,
) {
	defer wg.Done()
	for _, n := range arr {
		output <- n
	}

}
