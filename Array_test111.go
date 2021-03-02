package main

import (
	"fmt"
	"sort"
)

func main() {
	s1 := GenerateArr(10, 100)
	fmt.Print(s1)
	sort.Ints(s1)
	fmt.Print(s1)

}

//func generateArr(maxNum, scope int) []int {
//	var arr = make([]int, 0)
//	rand.Seed(time.Now().UnixNano()) //使用时间纳秒作为随机种子
//	for i := 1; i <= maxNum; i++ {
//		arr = append(arr, rand.Intn(scope))
//	}
//	return arr
//}
