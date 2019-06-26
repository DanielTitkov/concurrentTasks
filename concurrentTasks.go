package main


import (
    "log"
    "time"
    "sync"
    "errors"
)


func resolveTask(
    f func() error, taskIdxs chan int,
    errsCount *int, mx *sync.Mutex,
    wg *sync.WaitGroup,
) {
    if err := f(); err != nil {
        mx.Lock()
        *errsCount++
        mx.Unlock()
    }
    <- taskIdxs
    wg.Done()
}


func runConcurrentTasks(tasks []func() error, concurrency, maxErrs int) {
    if concurrency < 1 {
        log.Println("Zero and bellow concurrency makes no sense, quiting.")
        return
    }

    taskIdxs := make(chan int, concurrency)
    errsCount := 0
    var wg sync.WaitGroup
    var mutex sync.Mutex

    for i, t := range tasks {
        taskIdxs <- i
        if errsCount > maxErrs {
            log.Println("Got", errsCount, "errors, soft stop.")
            return
        } else {
            wg.Add(1)
            go resolveTask(t, taskIdxs, &errsCount, &mutex, &wg)
        }
    }
    wg.Wait()
    log.Println("All tasks are done, got", errsCount, "errors.")
}


func main() {
    log.Println("Start")

    tasks := []func() error {
        func() error {
            log.Println("1, start")
            time.Sleep(6*time.Second)
            log.Println("1, end")
            return nil
        },
        func() error {
            log.Println("2, start")
            time.Sleep(2*time.Second)
            log.Println("2, end")
            return nil
        },
        func() error {
            log.Println("3, start")
            time.Sleep(3*time.Second)
            log.Println("3, end")
            return nil
        },
        func() error {
            log.Println("4, start")
            time.Sleep(4*time.Second)
            log.Println("4, end")
            return nil
        },
        func() error {
            log.Println("5, start")
            time.Sleep(1*time.Second)
            log.Println("5, end")
            err := errors.New("random error")
            return err
        },
        func() error {
            log.Println("6, start")
            time.Sleep(1*time.Second)
            log.Println("6, end")
            err := errors.New("random error")
            return err
        },
        func() error {
            log.Println("7, start")
            time.Sleep(1*time.Second)
            log.Println("7, end")
            err := errors.New("random error")
            return err
        },
        func() error {
            log.Println("8, start")
            time.Sleep(1*time.Second)
            log.Println("8, end")
            err := errors.New("random error")
            return err
        },
        func() error {
            log.Println("9, start")
            time.Sleep(1*time.Second)
            log.Println("9, end")
            return nil
        },
    }

    runConcurrentTasks(tasks, 3, 1)
    log.Println("End")
}
