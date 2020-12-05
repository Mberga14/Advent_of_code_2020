package day3

import (
	"bufio"
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	treeMap := readinput()

	var down = []int{1, 1, 1, 1, 2}
	var right = []int{1, 3, 5, 7, 1}
	var numberOfTrees []int
	for index := 0; index < len(down); index++ {
		var trees = 0
		var limit = 0
		for i := down[index]; i < len(treeMap); i += down[index] {
			limit = (limit + right[index]) % len(treeMap[0])

			fmt.Println(i, limit)

			if treeMap[i][limit] == '#' {
				trees++
			}
		}

		numberOfTrees = append(numberOfTrees, trees)

	}
	count := 1
	for i := range numberOfTrees {
		count *= numberOfTrees[i]
		fmt.Println(numberOfTrees[i])
	}
	fmt.Println(numberOfTrees)
	fmt.Println(count)

}

func readinput() [][]byte {
	file, err := os.Open("input.txt")
	checkError(err)
	var treemap [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []byte(scanner.Text())
		treemap = append(treemap, line)
	}

	return treemap
}
