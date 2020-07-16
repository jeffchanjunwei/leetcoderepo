package main

import "fmt"

func main() {
	fmt.Println(numTrees(3))
}

func numTrees(n int) int {
	result := make([]int, n+1)
	result[0] = 1
	result[1] = 1
	if n == 0 || n == 1 {
		return result[n]
	}
	for i := 2; i <= n; i++ {
		temp := 0
		for j := 0; j < i; j++ {
			temp = temp + result[j]*result[i-j-1]
		}
		result[i] = temp
	}
	return result[n]
}

// 该题需要知道二叉搜索树定义，并且看出计算时后的递归逻辑，但直接使用递归方法的话会报超时错误，所以需要使用"空间换时间"的方法来优化递归算法
