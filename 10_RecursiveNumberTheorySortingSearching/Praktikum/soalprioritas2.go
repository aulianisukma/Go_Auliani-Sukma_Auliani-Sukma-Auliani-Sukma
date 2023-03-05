package main

import "fmt"

type pair struct {
	name string
	count int
}

func MostAppearItem(items []string) []pair {
	// membuat map untuk menghitung kemunculan setiap item
	counts := make(map[string]int)
	for _, item := range items {
		counts[item]++
	}

	// membuat slice untuk pasangan item-kemunculan
	var pairs []pair
	for item, count := range counts {
		pairs = append(pairs, pair{item, count})
	}

	// mengurutkan pasangan item-kemunculan berdasarkan kemunculan yang paling banyak
	for i := range pairs {
		for j := i + 1; j < len(pairs); j++ {
			if pairs[j].count > pairs[i].count {
				pairs[i], pairs[j] = pairs[j], pairs[i]
			}
		}
	}

	return pairs
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	// golang-> ruby->2 js->4
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	// C->1 D->2 B->3 A->4
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tebis"}))
	// football->1 basketball->1 tenis->1
}