package main 

import "fmt"
import "os"

/** 
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFilesExist(filename string) (bool) {
	var exist = true;
	if _,err := os.Stat(filename);os.IsNotExist(err){
		exist = false;
	}
	return exist;
}
func main() {
	userFile := "D:/www/gows/file/test.txt"//文件路径
	if checkFilesExist(userFile){
		fin,err := os.Open(userFile)//打开文件,返回File的内存地址
		fmt.Println("文件存在")
	}else {
		fmt.Println("文件不存在")
	}
	
	defer fin.Close() //延迟关闭资源
	if err != nil {
		fmt.Println(userFile,err)
		return
	}
	buf := make([]byte,1024)//创建一个初始容量为1024的slice,作为缓冲容器
	for {
		//循环读取文件数据到缓冲容器中,返回读取到的个数
		n,_ := fin.Read(buf)

		if 0 == n {
			break //如果读到个数为0,则读取完毕,跳出循环
		}
		//从缓冲slice中写出数据,从slice下标0到n,通过os.Stdout写出到控制台
		os.Stdout.Write(buf[:n])
	}

}