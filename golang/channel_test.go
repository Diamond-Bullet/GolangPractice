package golang

import (
	"fmt"
	"testing"
	"time"
)

// goroutine src/runtime/runtime2.go

// channel src/runtime/chan.go
// refer to `sudog`

// Sending to a full channel shall block the goroutine, until receiver draws one out of the channel so the channel is no longer full.
func TestSend2Full(t *testing.T) {
	ch := make(chan int, 3)
	go func() {
		time.Sleep(5 * time.Second)
		<-ch
	}()

	for i := 0; i < 4; i++ {
		ch <- 1
	}

	println("channel block stops")
}

// fetching data from empty channel, resulting in blocking the receiver goroutine.
func TestFetchFromEmpty(t *testing.T) {
	ch := make(chan int)
	go func() {
		for {
			fmt.Printf("go routine1: %d\n", <-ch)
		}
	}()
	go func() {
		for {
			fmt.Printf("go routine2: %d\n", <-ch)
		}
	}()

	for i := 0; ; i = 1 - i {
		ch <- i
		time.Sleep(time.Second)
	}
}

func TestForRange(t *testing.T) {
	ch := make(chan int, 2)

	// `for-range` traverses channel. if channel is empty, it's blocked. if channel is closed, it's over.
	go func() {
		for x := range ch {
			fmt.Printf("go routine: %d\n", x)
			time.Sleep(time.Second)
		}
		fmt.Println("go routine: channel has been closed")
	}()

	for i := 0; i < 5; i++ {
		ch <- 1
		fmt.Printf("main thread: %d\n", i)
	}
	close(ch)
	time.Sleep(10 * time.Second)
}

func TestFetchFromClosed(t *testing.T) {
	ch := make(chan int, 2)

	// fetch data from closed channel. the received are default values of channel's underlying type.
	go func() {
		for {
			fmt.Printf("go routine: %d\n", <-ch)
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 3; i++ {
		ch <- 1
		fmt.Printf("main thread: %d\n", i)
	}
	close(ch)
	time.Sleep(10 * time.Second)
}

func TestSend2Closed(t *testing.T) {
	ch := make(chan int, 2)

	// Sending datum to closed channel, causing panic.
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("go routine: %d\n", <-ch)
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for i := 0; i < 10; i++ {
		ch <- 1
		fmt.Printf("main thread: %d\n", i)
	}

	time.Sleep(10 * time.Second)
}

func TestOkJudgement(t *testing.T) {
	ch := make(chan int, 2)

	// `ok` semantics, helping you judge whether channel closes.
	go func() {
		for {
			x, ok := <-ch
			if ok {
				fmt.Printf("go routine: %d\n", x)
				time.Sleep(time.Second)
			} else {
				fmt.Println("go routine: channel has been closed")
				break
			}
		}
	}()

	for i := 0; i < 3; i++ {
		ch <- 1
		fmt.Printf("main thread: %d\n", i)
	}
	close(ch)
	time.Sleep(10 * time.Second)
}

// `Select` source code: runtime/select.go

// `for-select`, https://programs.wiki/wiki/analysis-of-go-bottom-series-select-principle.html
// select does not cycle(circulate, recur) itself.
func TestSelectGet(t *testing.T) {
	ch1 := make(chan int, 2)

	// randomly select a case from ready ones, get data from it and execute. if no default statement here and no case ready, it's blocked.
	// quit a goroutine through `done` channel. mention that a closed channel bring out zero-value.
	done := make(chan int)
	go func() {
		for {
			select {
			case <-done:
				fmt.Printf("go routine done, exit")
				return
			case x := <-ch1:
				fmt.Printf("go routine ch1: %d\n", x)
			case x := <-time.After(4 * time.Second):
				fmt.Printf("go routine time out: %v\n", x)
				//default:
				//	fmt.Println("go routine nothing from channel")
				//	time.Sleep(500 *time.Millisecond)
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch1 <- i
		time.Sleep(time.Second)
	}
	close(done)
	time.Sleep(5 * time.Second)
}

func TestSelectSend(t *testing.T) {
	// randomly select a channel and send data to it.
	// if the channel is full, this branch can not proceed now. other branches are tested if they can be executed.
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)
	go func() {
		for {
			select {
			case ch1 <- 1:
				fmt.Printf("go routine ch1: %d\n", 1)
			case ch2 <- 1:
				fmt.Printf("go routine ch2: %d\n", 1)
			default:
				fmt.Println("channel are filled with data!")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	for i := 0; i < 5; i++ {
		<-ch1
		<-ch2
		time.Sleep(time.Second)
	}
	time.Sleep(10 * time.Second)
}

// single-direction(one-way) channel, generally applied to receiver and sender respectively.
// so there is stricter restriction on both sides to prevent misoperations(misuse).

// random number based on channel, just for fun.
func TestRandomOnChannel(t *testing.T) {
	random(100)
}

func random(n int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 0; i < n; i++ {
			select {
			case c <- 0:
			case c <- 1:
			}
		}
	}()
	return c
}
