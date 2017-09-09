/**
* 服务化http://studygolang.com/articles/7176
* 比如我们加载一个网站的时候，例如我们登入新浪微博，我们的消息数据应该来自一个独立的服务，
这个服务只负责 返回某个用户的新的消息提醒。
*/

package main

import "fmt"

func get_notification(user string) chan string {
	/**
	 * 此处可以查询数据库获取新消息等等..
	 */
	notifications := make(chan string)
	go func() { //悬挂一个信道出去
		notifications <- fmt.Sprintf("Hi %s,welcome to weibo.com", user)

	}()

	return notifications
}

func main() {
	jack := get_notification("jack") //获取jack的信息
	joe := get_notification("joe")   //获取joe的消息

	//获取消息的返回
	fmt.Println(<-jack)
	fmt.Println(<-joe)
}
