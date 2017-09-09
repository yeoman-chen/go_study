/**生成器
* 在Python中我们可以使用yield关键字来让一个函数成为生成器，在Go中我们可以使用信道来制造生成器(一种lazy load类似的东西)。
* 当然我们的信道并不是简单的做阻塞主线的功能来使用的哦。
* 下面是一个制作自增整数生成器的例子，直到主线向信道索要数据，我们才添加数据到信道：
 */
package main

import (
	"fmt"
)

func xrange() chan int { //xrange用来生成自增的整数
	var ch chan int = make(chan int)

	go func() { //开出一个goroutine
		for i := 0; ; i++ {
			ch <- i //直到信道索要数据，才把i添加进信道
		}

	}()
	return ch
}

func main() {
	generator := xrange()

	for i := 0; i < 100; i++ { //生成100个自增的整数
		fmt.Println(<-generator)
	}
}
