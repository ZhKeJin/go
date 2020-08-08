package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var(
	myMap = make(map[int]int,10)

	lock sync.Mutex
)

func testq(n int){
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//加锁
	lock.Lock()
	myMap[n] = res
    lock.Unlock()
}

func main() {
	//获取当前 系统的cpu  数量
	num := runtime.NumCPU()

	runtime.GOMAXPROCS(num)
	fmt.Println("num=",num)

	//map自动扩容
	map1 := make(map[string]string,1)

	map1["name"] = "zhangkejin"
	map1["age"] = "zhangkejin"
	fmt.Println(map1)

	//goruntine
	//计算各个数字的阶乘
	//fatal error: concurrent map writes   出现抢占 报错
	for i := 1; i < 100; i++ {
		go testq(i)
	}

	time.Sleep(time.Second*10)


	for i, val := range myMap {
		fmt.Printf("i:%v,val:%v\n",i,val)
	}


    //加锁进行互斥


}