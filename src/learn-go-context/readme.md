WaitGroup的使用时机：
    当一个工作任务需要拆分成多个子JOb，且全部子JOb处理完成后，
    才能进行下一个阶段程序处理，
    

```go
	var wg = sync.WaitGroup{}
	wg.Add(2)
	go func() {
		target(5, "🐏")
		wg.Done()  //finish goroutine1
	}()
	go func() {
		target(5, "🐂")
		wg.Done()  //finish goroutine1
	}()
//等待全部子任务全部执行完成后，相当于join()
	wg.Wait()    //wait for tasks to be done
```





Channel 使用时机：

​	主线程需要和goroutine进行沟通，通知它执行结束，比如超时结束

```go
func main(){
    exit := make(chan bool)
    go func(){
        for{
            select{
                case <-exit:
                fmt.Println("Exit")
                return
                default:
                fmt.Println("working")
            }
        }
    }()
    
    time.Sleep(5*time.Second) //最大等待时间
    fmt.Println("Notify Exit") //通知goroutine结束
    exit <-true  //keep main goroutine alive
}

```



多个Channel ，Channel 内部还有Channel 。嵌套Channel ，它会一直嵌套通知每一个goroutine，让一个goroutine下的子goroutine也会被通知到，比如超时情况，这样就可以让子goroutine的子goroutine也能跟着父goroutine一起死亡（守护线程一样），这样就不会造成子goroutine的子goroutine还在后台执行占用cpu资源

```go
ctx, cancel := context.WithCancel(context.Background())
//多个goroutine进行工作
go func(ctx context.Context, name string) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(name, "  ---->协程任务被迫结束")
				return
			default:
				fmt.Println(name, " working")
			}
		}
	}(ctx, "goroutine1")
go func(ctx context.Context, name string) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(name, "  ---->协程任务被迫结束")
				return
			default:
				fmt.Println(name, " working")
			}
		}
	}(ctx, "goroutine2")

//设置主线程最大等待时间，到点通知结束
time.Sleep(50 * time.Millisecond)
//当执行到cancel()的时候，即私有协程对象全部执行结束，不会再等待，防止超时处理
cancel()
time.Sleep(5 * time.Millisecond)
```



