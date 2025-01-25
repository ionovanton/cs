## Progress
### 0 / 1

# Goroutines

Synchronous work
```go
package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const (
	iterationsNum = 4
	goroutinesNum = 4
)

func doWork(th int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 0; j < iterationsNum; j++ {
		fmt.Printf(formatWork(th, j))
		time.Sleep(time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < goroutinesNum; i++ {
		wg.Add(1)
		doWork(i, &wg)
	}
	wg.Wait()
}

func formatWork(in, j int) string {
	return fmt.Sprintln(strings.Repeat("  ", in), "█",
		strings.Repeat("  ", goroutinesNum-in),
		"th", in,
		"iter", j, strings.Repeat("■", j))
}

/Users/ayionov/Library/Caches/JetBrains/GoLand2023.1/tmp/GoLand/___go_build_main
█          th 0 iter 0
█          th 0 iter 1 ■
█          th 0 iter 2 ■■
█          th 0 iter 3 ■■■
  █        th 1 iter 0
  █        th 1 iter 1 ■
  █        th 1 iter 2 ■■
  █        th 1 iter 3 ■■■
    █      th 2 iter 0
    █      th 2 iter 1 ■
    █      th 2 iter 2 ■■
    █      th 2 iter 3 ■■■
      █    th 3 iter 0
      █    th 3 iter 1 ■
      █    th 3 iter 2 ■■
      █    th 3 iter 3 ■■■
```

For asynchronous parallel work add `go` keyword – `go doWork(i, &wg)`
```
       █    th 3 iter 0 
 █          th 0 iter 0 
   █        th 1 iter 0 
     █      th 2 iter 0 
     █      th 2 iter 1 ■
       █    th 3 iter 1 ■
 █          th 0 iter 1 ■
   █        th 1 iter 1 ■
   █        th 1 iter 2 ■■
     █      th 2 iter 2 ■■
       █    th 3 iter 2 ■■
 █          th 0 iter 2 ■■
 █          th 0 iter 3 ■■■
     █      th 2 iter 3 ■■■
   █        th 1 iter 3 ■■■
       █    th 3 iter 3 ■■■
```

Let's imitate network call via runtime.Gosched()
```go
func doWork(th int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 0; j < iterationsNum; j++ {
		fmt.Printf(formatWork(th, j))
		runtime.Gosched()
	}
}
```

Now one goroutine makes one operation and passes work to another goroutine via `runtime.Gosched()`
```
       █    th 3 iter 0 
 █          th 0 iter 0 
   █        th 1 iter 0 
     █      th 2 iter 0 
       █    th 3 iter 1 ■
 █          th 0 iter 1 ■
   █        th 1 iter 1 ■
     █      th 2 iter 1 ■
       █    th 3 iter 2 ■■
 █          th 0 iter 2 ■■
   █        th 1 iter 2 ■■
     █      th 2 iter 2 ■■
       █    th 3 iter 3 ■■■
 █          th 0 iter 3 ■■■
   █        th 1 iter 3 ■■■
     █      th 2 iter 3 ■■■
```

# Channels
## Unbuffered channels
Channels are made for passing data ownership between goroutines.

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// Unbuffered channel
	ch := make(chan int)

	wg.Add(1)
	go func(in chan int) {
		defer wg.Done()
		val := <-in
		fmt.Println("GO: get from chan:", val)
		fmt.Println("GO: after read from chan")
	}(ch)

	ch <- 42

	fmt.Println("MAIN: after put to chan")
	wg.Wait()
}
```
```
GO: get from chan: 42
GO: after read from chan
MAIN: after put to chan
```

Let's write something else to channel
```go
...
	ch <- 42
	ch <- 21

	fmt.Println("MAIN: after put to chan")
...
```
```
GO: get from chan: 42
GO: after read from chan
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        go_practice/main.go:22 +0xf8

```

## Buffered channels

In buffered channels writer locks until someone reads from channel.
```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// Buffered channel
	ch := make(chan int, 1)

	wg.Add(1)
	go func(in chan int) {
		defer wg.Done()
		val := <-in
		fmt.Println("GO: get from chan:", val)
		fmt.Println("GO: after read from chan")
	}(ch)

	ch <- 42
	ch <- 21

	fmt.Println("MAIN: after put to chan")
	wg.Wait()
}
```

No error this time, because there is a place where we can put the value without being blocked.
```
GO: get from chan: 42
GO: after read from chan
MAIN: after put to chan
```

## Loop and channels

There will be an error because of channeled is not closed on loop end
```go
package main

import (
	"fmt"
)

func main() {
	in := make(chan int)

	go func(out chan<- int) {
		for i := 0; i <= 4; i++ {
			fmt.Println("before", i)
			out <- i
			fmt.Println("after", i)
		}
		// crucial for loop end
		//close(out)
		fmt.Println("generator finish")
	}(in)

	// still waits for some data from channel
	for i := range in {
		fmt.Println("\tget", i)
	}
}
```
```
before 0
after 0
before 1
        get 0
        get 1
after 1
before 2
after 2
before 3
        get 2
        get 3
after 3
before 4
after 4
generator finish
        get 4
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
        go_practice/main.go:21 +0xe8
```

Uncommenting `close(out)` will solve this problem
```go
...
    for i := 0; i <= 4; i++ {
        fmt.Println("before", i)
        out <- i
        fmt.Println("after", i)
    }
    // crucial for loop end
    close(out)
    fmt.Println("generator finish")
...
```
```
before 0
after 0
before 1
        get 0
        get 1
after 1
before 2
after 2
before 3
        get 2
        get 3
after 3
before 4
after 4
generator finish
        get 4
```

## Channel multiplexing using `select`

```go
package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	select {
	// nothing to read from ch1
	case val := <-ch1:
		fmt.Println("ch1 val", val) 
	// no one read from channel, so this case does not hold
	case ch2 <- 1:
		fmt.Println("put val to ch2")
	// default case holds
	default:
		fmt.Println("default case")
	}
}
```
```
default case
```

Let's put something to ch1
```go
ch1 := make(chan int, 1)
...
ch1 <- 1
...
```
```
ch1 val 1
```

Let's read something from ch2
```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	//ch1 <- 1
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		val := <-ch2
		fmt.Println("read from ch2 val", val)
	}()

	// Cannot wait here due to no one writes to ch2, so it will be deadlock
	//wg.Wait()

	select {
	case val := <-ch1:
		fmt.Println("ch1 val", val)
	case ch2 <- 1:
		fmt.Println("put val to ch2")
	default:
		fmt.Println("default case")
	}

	// Can wait here due to there is goroutine that reads from ch2 and main goroutine that puts value on ch2
	wg.Wait()
}
```
```
put val to ch2
read from ch2 val 1
```

If no default case is set, there will be a deadlock
```go
package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int)

	//ch1 <- 1

	select {
	case val := <-ch1:
		fmt.Println("ch1 val", val)
	case ch2 <- 1:
		fmt.Println("put val to ch2")
		//default:
		//	fmt.Println("default case")
	}
}
```
```
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [select]:
main.main()
        go_practice/main.go:13 +0x84
```

Select picks random case from multiple to choose from
```go
package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)

	ch1 <- 1
	ch1 <- 2
	ch2 <- 3

LOOP:
	for {
		select {
		case v1 := <-ch1:
			fmt.Println("chan1 val", v1)
		case v2 := <-ch2:
			fmt.Println("chan2 val", v2)
		default:
			break LOOP
		}
	}
}
```
```
chan2 val 3
chan1 val 1
chan1 val 2
```

## Cancel channel
```go
package main

import (
	"fmt"
)

func main() {
	cancelChannel := make(chan struct{})
	dataChannel := make(chan int)

	go func(cancelChannel chan struct{}, dataChannel chan int) {
		val := 0
		for {
			select {
			case <-cancelChannel:
				return
			case dataChannel <- val:
				val++
			}
		}
	}(cancelChannel, dataChannel)

	for curVal := range dataChannel {
		fmt.Println("read", curVal)
		if curVal > 3 {
			fmt.Println("send cancel")
			cancelChannel <- struct{}{}
			break
		}
	}
}
```
```
read 0
read 1
read 2
read 3
read 4
send cancel
```

# Multiprocessing programming

## Timeouts
```go
package main

import (
	"fmt"
	"time"
)

func longSQLQuery() chan string {
	ch := make(chan string, 1)
	go func() {
		// some long operation
		time.Sleep(time.Second * 2)
		ch <- "executed query"
	}()
	return ch
}

func main() {
	timer := time.NewTimer(time.Second)
	select {
	case <-timer.C:
		fmt.Println("timer.C timeout happened")
	case <-time.After(time.Minute):
		fmt.Println("time.After timeout happened")
	case result := <-longSQLQuery():
		// freeing resource
		if !timer.Stop() {
			<-timer.C
		}
		fmt.Println("operation result:", result)
	}
}
```
```
timer.C timeout happened
```

Raising timeout up to 3 seconds will yield in the result
```go
...
timer := time.NewTimer(time.Second * 3)
...
```
```
operation result: executed query
```

## Ticker
### Periodic events
`time.Ticker` is used to periodic events
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second) // tick every second
	i := 0
	// ticker.C sends tick time via channel
	for tickTime := range ticker.C {
		i++
		fmt.Println("step", i, "time", tickTime)
		if i > 5 {
			ticker.Stop()
			break
		}
	}
	fmt.Println("total", i)
}
```
```
step 1 time 2025-01-19 05:07:35.920129208 +0300 MSK m=+1.000179251
step 2 time 2025-01-19 05:07:36.9201225 +0300 MSK m=+2.000180501
step 3 time 2025-01-19 05:07:37.920112791 +0300 MSK m=+3.000179084
step 4 time 2025-01-19 05:07:38.920104833 +0300 MSK m=+4.000179709
step 5 time 2025-01-19 05:07:39.920095333 +0300 MSK m=+5.000179168
step 6 time 2025-01-19 05:07:40.920086791 +0300 MSK m=+6.000179417
total 6
```

### Endless periodic events
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	c := time.Tick(time.Second)
	i := 0
	for tickTime := range c {
		i++
		fmt.Println("step", i, "time", tickTime)
		if i > 5 {
			break
		}
	}
}
```
```
step 1 time 2025-01-19 05:10:40.980095792 +0300 MSK m=+1.000172668
step 2 time 2025-01-19 05:10:41.980085667 +0300 MSK m=+2.000171668
step 3 time 2025-01-19 05:10:42.9800775 +0300 MSK m=+3.000171709
step 4 time 2025-01-19 05:10:43.980068917 +0300 MSK m=+4.000171876
step 5 time 2025-01-19 05:10:44.980062292 +0300 MSK m=+5.000172960
step 6 time 2025-01-19 05:10:45.980053792 +0300 MSK m=+6.000173043
```

## Afterfunc
`Afterfunc` makes callback after some period of time

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.AfterFunc(time.Second, func() {
		fmt.Println("Hello world")
	})

	time.Sleep(time.Second * 2)
	timer.Stop()
	// can't put here because timer will be stopped before AfterFunc can execute its function
	//time.Sleep(time.Second * 2)
}
```
```
Hello world
```

# Context
## Cancel context

```go
package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(ctx context.Context, i int, result chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	waitTime := time.Duration(rand.Intn(100)+10) * time.Millisecond
	fmt.Println("worker", i, "start with time", waitTime)
	select {
	case <-ctx.Done():
		return
	case <-time.After(waitTime):
		fmt.Println("worker", i, "done")
		result <- i
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	result := make(chan int, 1)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ctx, i, result, &wg)
	}

	// can't put here because other goroutines should know that
	// work is cancelled via cancel context ctx.Done() channel
	//wg.Wait()

	foundBy := <-result
	fmt.Println("result found by worker", foundBy)
	cancel()

	wg.Wait()
}
```
```
worker 9 start with time 64ms
worker 7 start with time 94ms
worker 1 start with time 66ms
worker 4 start with time 104ms
worker 6 start with time 87ms
worker 5 start with time 106ms
worker 2 start with time 31ms
worker 3 start with time 12ms
worker 8 start with time 21ms
worker 0 start with time 24ms
worker 3 done
```

## Timeout context

```go
package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func worker(ctx context.Context, i int, result chan int) {
	waitTime := time.Duration(rand.Intn(100)+10) * time.Millisecond
	fmt.Println("worker", i, "start with time", waitTime)
	select {
	case <-ctx.Done():
		return
	case <-time.After(waitTime):
		fmt.Println("worker", i, "done")
		result <- i
	}
}

func main() {
	timeout := time.Duration(30) * time.Millisecond
	fmt.Println("timeout value is set to", timeout)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := make(chan int, 1)

	for i := 0; i < 10; i++ {
		go worker(ctx, i, result)
	}

	for {
		select {
		case <-ctx.Done():
			return
		case foundBy := <-result:
			fmt.Println("result found by worker", foundBy)
		}
	}
}
```

Only two goroutines workers executes work under 30 ms: 1 and 4. Other workers are cancelled due to timeout.
```
timeout value is set to 30ms
worker 9 start with time 59ms
worker 7 start with time 58ms
worker 8 start with time 63ms
worker 5 start with time 81ms
worker 4 start with time 15ms
worker 1 start with time 23ms
worker 0 start with time 45ms
worker 6 start with time 53ms
worker 3 start with time 52ms
worker 2 start with time 41ms
worker 4 done
result found by worker 4
worker 1 done
result found by worker 1
```

## Parallelism in Go
```go
package main

import (
	"errors"
	"fmt"
	"time"
)

func heavySqlWork() chan string {
	// should be buffered channel
	// write will be synchronous call
	// if error occur before read from channel
	// then no one read from channel
	// this will be goroutine leak
	result := make(chan string, 1)
	go func(out chan<- string) {
		time.Sleep(time.Second * 3)
		out <- "query executed"
	}(result)
	return result
}

func anotherSqlWork() error {
	return errors.New("sql error")
}

func main() {
	resultCh := heavySqlWork()

	err := anotherSqlWork()
	if err != nil {
		return
	}

	heavySqlResult := <-resultCh
	fmt.Println(heavySqlResult)
}
```

## Worker pool
```go
package main

import (
	"fmt"
	"strings"
	"time"
)

var goroutinesNum = 3

func formatWork(in int, input string) string {
	return fmt.Sprint(strings.Repeat("  ", in), "█",
		strings.Repeat("  ", goroutinesNum-in),
		in, " recieved work ", input)
}

func startWorker(workerNum int, workerInput <-chan string) {
	for input := range workerInput {
		fmt.Println(formatWork(workerNum, input))
	}
	fmt.Println(workerNum, "finished work")
}

func main() {
	workerInput := make(chan string, 1)
	for i := 0; i < goroutinesNum; i++ {
		go startWorker(i, workerInput)
	}

	months := []string{
		"Январь", "Февраль", "Март", "Апрель",
		"Май", "Июнь", "Июль", "Август",
		"Сентябрь", "Октябрь", "Ноябрь", "Декабрь",
	}

	for _, monthName := range months {
		workerInput <- monthName
	}
	
	close(workerInput)

	time.Sleep(time.Millisecond)
}
```
```
█      0 recieved work Февраль
█      0 recieved work Апрель
█      0 recieved work Май
  █    1 recieved work Март
    █  2 recieved work Январь
  █    1 recieved work Июль
    █  2 recieved work Август
  █    1 recieved work Сентябрь
    █  2 recieved work Октябрь
    █  2 recieved work Декабрь
2 finished work
  █    1 recieved work Ноябрь
1 finished work
█      0 recieved work Июнь
0 finished work
```

It is essential to close channel as it waits to workers to complete their jobs
```go
// close(workerInput)
```

No workers finished
```
  █    1 recieved work Март
  █    1 recieved work Апрель
  █    1 recieved work Май
    █  2 recieved work Январь
    █  2 recieved work Июль
    █  2 recieved work Август
  █    1 recieved work Июнь
█      0 recieved work Февраль
    █  2 recieved work Сентябрь
    █  2 recieved work Декабрь
█      0 recieved work Ноябрь
  █    1 recieved work Октябрь
```

