package algorithm

import "fmt"

//输入一个正数n，输出所有和为n 连续正数序列
func Algo1(a int) {
	l := a/2 + a%2
	for i := 1; i <= l; i++ {
		sum := 0
		for j := i; j <= l; j++ {
			sum += j
			if sum == a {
				fmt.Printf("%d -- %d \n", i, j)
			}
			if sum > a {
				break
			}
		}
	}
}
