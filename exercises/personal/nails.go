package expersonal

import "fmt"

type plank struct {
	a int
	b int
}

func Solution(A []int, B []int, C []int) int {
	planks := []plank{}
	nailed := make(map[int]bool)

	for i := 0; i < len(A); i++ {
		newPlank := plank{A[i], B[i]}
		planks = append(planks, newPlank)
	}

	nailIndex := 0
	for _, plank := range planks {

		if plank.a == C[nailIndex] || plank.b == C[nailIndex] {
			nailed[plank.a] = true
		}
		nailIndex++
	}

	fmt.Println(planks)
	return len(nailed)
}

func main() {
	A := []int{1, 4, 5, 8, 1, 1, 1, 3, 3, 3}
	B := []int{4, 5, 9, 10, 4, 4, 4, 9, 9, 9}
	C := []int{4, 6, 7, 10, 2}
	nNail := Solution(A, B, C)
	fmt.Println(nNail)
}
