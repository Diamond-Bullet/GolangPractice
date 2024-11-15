package coding

import (
	"GolangPractice/utils/logger"
	"sync"
	"testing"
)

func TestPipeline(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}

	mapper1 := func(i int) int {
		return i + 1
	}
	mapper2 := func(i int) int {
		return i * 2
	}

	// Method 1: Embedded functions
	result := Work(Work(Generate(source), mapper1), mapper2)
	for v := range result {
		logger.Info(v)
	}

	// Method 2: Functions list
	result = Pipeline(source, Generate[int], Work[int], mapper1, mapper2)
	for v := range result {
		logger.Info(v)
	}
}

func TestConcurrentPipeline(t *testing.T) {
	nums := makeRange(1, 10000)

	mapper1 := func(i int) int {
		return i + 1
	}
	mapper2 := func(i int) int {
		return i * 2
	}

	in := Generate(nums)

	const nProcess = 5
	var chans [nProcess]<-chan int
	for i := range chans {
		chans[i] = Work(Work(in, mapper1), mapper2)
	}

	for n := range sum(merge(chans[:])) {
		logger.Info(n)
	}
}

func Generate[T any](source []T) <-chan T {
	ch := make(chan T)
	go func() {
		for _, v := range source {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

func Work[T any](source <-chan T, mapper func(T) T) <-chan T {
	ch := make(chan T)
	go func() {
		for v := range source {
			ch <- mapper(v)
		}
		close(ch)
	}()
	return ch
}

type GenerateFunc[T any] func(source []T) <-chan T
type WorkFunc[T any] func(source <-chan T, mapper func(T) T) <-chan T
type MapFunc[T any] func(T) T

func Pipeline[T any](source []T, generate GenerateFunc[T], work WorkFunc[T], mappers ...MapFunc[T]) <-chan T {
	ch := generate(source)
	for _, mapper := range mappers {
		ch = work(ch, mapper)
	}
	return ch
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func merge(cs []<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan int) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func sum(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var localSum = 0
		for n := range in {
			localSum += n
		}
		out <- localSum
		close(out)
	}()
	return out
}
