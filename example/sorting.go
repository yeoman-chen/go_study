package main 

import "fmt"
import "sort"

func main() {
	strs := []string{"c","a","b"}
	sort.Strings(strs)
	fmt.Println("Srings:",strs)

	ints := []int{7,2,4}
	sort.Ints(ints)
	fmt.Println("Ints:",ints)
	//判断是否已排序
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:",s)

}