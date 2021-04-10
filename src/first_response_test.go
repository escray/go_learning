package go_learning

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10
	// ch := make(chan string)
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	t.Log("Before: ", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Second * 1)
	t.Log("After: ", runtime.NumGoroutine())
}

// func TestWaitAllResponseDone(t *testing.T) {
//     numOfRunner := 10
//     var wg sync.WaitGroup
//     //这里 length 和 capacity 都是 10，接收返回值
//     strSlice := make([]string, numOfRunner)
//     for i := 0; i < numOfRunner; i++ {
//         wg.Add(1)
//         go func(i int) {
//             //通过i 保证每个协程修改的地方独立，thread-safe
//             strSlice[i] = runTask(i)
//             wg.Done()
//         }(i)
//     }

//     wg.Wait()
//     // 使用 12讲 strings的Join函数
//     fmt.Println("all result:\n", strings.Join(strSlice, ",\n"))
//     //验证 切片的长度和容量
//     fmt.Print(len(strSlice), cap(strSlice))
// }

// 实现所有任务返回且关闭channel
func AllResponseInit() string {
    numOfRoutine := 10
    ch := make(chan string, numOfRoutine)
    for i := 0; i < numOfRoutine; i++ {
        go func(i int) {
            ch <- fmt.Sprintf("The result is from %d", i)
        }(i)
    }
    totalResult := ""
    for i := 0; i < numOfRoutine; i++ {
        // 阻塞式获取channel中的所有消息
        totalResult += <-ch + "\n"
    }
    // 关闭channel
    close(ch)
    return totalResult
}

func AllResponseWithRange() string {
    numOfRoutine := 10
    var wg sync.WaitGroup
    ch := make(chan string, numOfRoutine)
    for i := 0; i < numOfRoutine; i++ {
        wg.Add(1)
        go func(i int) {
            ch <- fmt.Sprintf("The result is from %d", i)
            wg.Done()
        }(i)
    }
    wg.Wait()
    // 关闭channel
    close(ch)
    totalResult := ""
    for v := range ch{
        totalResult += v + "\n"
    }
    return totalResult
}

// func TestGetAllResponse(t *testing.T) {
//     t.Log(AllResponseInit())
//     t.Log(AllResponseWithRange())
// }

func getFirstAgain() string {
	// var wg sync.WaitGroup
    ch := make(chan string, 2)
    for i := 0; i < 2; i++ {
			// wg.Add(1)
        go func(i int) {
            fmt.Println("start:" + strconv.Itoa(i))
            time.Sleep(5 * time.Millisecond)
            ch <- "return:" + strconv.Itoa(i)
            fmt.Println("finish:" + strconv.Itoa(i))
						// wg.Done()
        }(i)
    }
		// wg.Wait()
		// close(ch)
    return <-ch
}

func TestFirstRes(t *testing.T) {

  fmt.Println(getFirstAgain())
  time.Sleep(1 * time.Second)
}
