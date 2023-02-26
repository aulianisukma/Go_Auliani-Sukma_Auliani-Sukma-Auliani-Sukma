package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	merged := append(arrayA, arrayB...)
    unique := make(map[string]bool)
    result := []string{}

    for _, item := range merged {
        if !unique[item] {
            unique[item] = true
            result = append(result, item)
        }

}
}
func main() {
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", "bryan"]

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devi; jin", "yoshimitsu", "alisa", "law"}))
	// ["alisa", "yoshimitsu", "devil jin", "law"]

	fmt.Println(ArayyMerge([]string{}, []string{"devil jin", "sergei"}))
	//["devi; jin", "sergei"]

	fmt.Println(ArrayMerge([]string{"howrang"}, []string{}))
	// ["howrang"]

	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []

}
