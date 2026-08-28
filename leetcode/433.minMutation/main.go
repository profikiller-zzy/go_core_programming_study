package main

import "fmt"

type state struct {
	step int
	gene string
}

func minMutation(startGene string, endGene string, bank []string) int {
	if startGene == endGene {
		return 0
	}
	validGene := make(map[string]bool)
	choice := []byte{'A', 'C', 'G', 'T'}
	for _, gene := range bank {
		validGene[gene] = true
	}
	if !validGene[endGene] {
		return -1
	}
	queue := []state{{0, startGene}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		gene := []byte(cur.gene)
		for index := 0; index < 8; index++ {
			curCh := gene[index]
			for _, ch := range choice {
				if ch == curCh {
					continue
				}
				gene[index] = ch
				next := string(gene)
				if validGene[next] {
					if next == endGene {
						return cur.step + 1
					}
					queue = append(queue, state{cur.step + 1, next})
					delete(validGene, next)
				}
			}
			gene[index] = curCh
		}
	}
	return -1
}
func main() {
	startGene := "AACCGGTT"
	endGene := "AACCGGTA"
	bank := []string{"AACCGGTA"}
	fmt.Println(minMutation(startGene, endGene, bank))
}
