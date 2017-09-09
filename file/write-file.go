package main 

import "fmt"
import "os"

func main () {
	userFile := "D:/www/gows/file/test.txt"//文件路径
	fout,err := os.Create(userFile) //根据路径创建File的内存地址
	defer fout.Close() //延迟关闭资源
	if err != nil {
		fmt.Println(userFile,err)
		return 
	}
	for i := 0; i < 10; i++ {
		fout.WriteString("Hello world!\r\n")//写入字符串
		fout.Write([]byte("abcd!\r\n"))//强转成byte slice后再写入
	}

}