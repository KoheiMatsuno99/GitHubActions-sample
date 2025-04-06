package sample

import "fmt"

func main() {
	a := 10
	b := 20

	// Go 1.21 以降では、min 関数が組み込みとして利用可能
	minVal := min(a, b)

	fmt.Println("Minimum:", minVal)
}
