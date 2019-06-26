package main


import (
    "fmt"
    "time"
    "sync"
)


func runConcurrentTasks(tasks []func() error, concurrency, errs int) {
    if concurrency < 1 {
        return
    }

    taskIdxs := make(chan int, concurrency)
    var wg sync.WaitGroup

    for i, t := range tasks {
        taskIdxs <- i
        wg.Add(1)
        go func(f func() error) {
            f()
            defer wg.Done()
            <- taskIdxs
        }(t)
    }
    wg.Wait()
}


func main() {
    fmt.Println("Start")

    tasks := []func() error {
        func() error {
            fmt.Println("1, start", time.Now())
            time.Sleep(1*time.Second)
            fmt.Println("1, end", time.Now())
            return nil
        },
        func() error {
            fmt.Println("2, start", time.Now())
            time.Sleep(2*time.Second)
            fmt.Println("2, end", time.Now())
            return nil
        },
        func() error {
            fmt.Println("3, start", time.Now())
            time.Sleep(3*time.Second)
            fmt.Println("3, end", time.Now())
            return nil
        },
        func() error {
            fmt.Println("4, start", time.Now())
            time.Sleep(4*time.Second)
            fmt.Println("4, end", time.Now())
            return nil
        },
    }

    runConcurrentTasks(tasks, 3, 1)
    fmt.Println("End")
}
