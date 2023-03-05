package main

import (
 "fmt"
)

func getBinaryArray(n int) []int {
 ans := make([]int, n+1)
 for i := 0; i <= n; i++ {
  ans[i] = i
 }
 return ans
}

func main() {
 fmt.Println(getBinaryArray(2))  // [0 1 2]
 fmt.Println(getBinaryArray(3))  // [0 1 2 3]
}