package semaphore

import (
	"sync"
)

// ArrSum demonstrates a Producer-Consumer, fan-out / fan-in, Channel Coordination concurrency patterns.
// Tags: semaphore, concurrency, channels, fan-in, fan-out, waitgroup, producer, consumer
func ArrSum(numbers ...[]int) int {
	if len(numbers) == 0 {
		return 0
	}

	var wg sync.WaitGroup

	sumCh := make(chan int)

	intsCh := make(chan int)
	arrsCh := make(chan []int)

	//semSize := runtime.NumCPU()
	semSize := 3
	semaphore := make(chan struct{}, semSize)

	go func() { //produce chan with input arrays
		defer close(arrsCh)
		for _, numberArr := range numbers {
			arrsCh <- numberArr
		}
	}()

	go func() { // consume ints flow and produce outcome
		sum := 0
		for value := range intsCh {
			sum += value
		}
		sumCh <- sum
	}()

	for arr := range arrsCh {
		select {
		case semaphore <- struct{}{}:
			wg.Add(1)
			a := arr
			go intProducer(&wg, a, intsCh, semaphore)
		}
	}

	go func() {
		wg.Wait()
		close(intsCh)
	}()

	return <-sumCh
}

func intProducer(
	wg *sync.WaitGroup,
	arr []int,
	output chan<- int,
	sem chan struct{},
) {
	defer wg.Done()
	defer func() { <-sem }()

	for _, n := range arr {
		output <- n
	}

}
