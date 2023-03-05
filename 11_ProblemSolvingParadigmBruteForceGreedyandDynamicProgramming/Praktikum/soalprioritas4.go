package main

import "fmt"

func SimpleEquations(a, b, c int) {
	found := false
	for x := -1000; x <= 1000; x++ {
		for y := -1000; y <= 1000; y++ {
			for z := -1000; z <= 1000; z++ {
				if x+y+z == a && x*y*z == b && x*x+y*y+z*z == c && x != y && x != z && y != z {
					fmt.Println(x, y, z)
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if found {
			break
		}
	}
	if !found {
		fmt.Println("no solution")
	}
}

func main() {
	SimpleEquations(1, 2, 3) // no solution
	SimpleEquations(6, 6, 14) // 1 2 3
}