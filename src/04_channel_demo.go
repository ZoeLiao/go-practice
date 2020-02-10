package main

import (
    // 格式化
    "fmt"
    // 隨機數
    "time"
    "github.com/gookit/color"
)

const TIMEOUT = 5

func producer(header string, channel chan<-string) {
    // 無線循環
    i := 0
    for {
        i += 1
        channel <- fmt.Sprintf("%s says: give me %d second, the program will be stopped if you wait more than %d second", header, i, TIMEOUT)
        time.Sleep(time.Second * time.Duration(i))
    }
}

func producer2(header string, channel chan<-string){
    for {
        channel <- fmt.Sprintf("%s says: hurry up please", header)
        time.Sleep(time.Second)
    }
}

func customer(channel <-chan string) {
    for {
        // 用 select 處理超時或終止以防 DoS attach
        select {
            // 從 channel 取出數據 （阻塞直到 channel 返回數據）
            case <- channel:
                message := <-channel
                color.info.Println(message)
            case <- time.After(time.Second * TIMEOUT):
                fmt.Println("quit")
                return
        }
    }
}

func main() {
    channel := make(chan string)
    go producer("cat", channel)
    go producer2("dog", channel)
    customer(channel)
}
