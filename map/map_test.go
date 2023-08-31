package map_test

import "fmt"

func mayWithArrayKey() {
	m := map[[2]int]bool{} // map的key只能用array，不能用slice
	fmt.Println(m)         // map[]

	m2 := map[[2]int]bool{[2]int{8, 7}: true} // 初值
	fmt.Println(m2)                           // map[[8 7]:true]

	m2[[2]int{8, 7}] = true
	fmt.Println(m2) // map[[8 7]:true]

	m2[[2]int{7, 8}] = true // array的值會被當成key，順序是有差的，*不是set的概念*
	fmt.Println(m2)         // map[[7 8]:true [8 7]:true]
}
