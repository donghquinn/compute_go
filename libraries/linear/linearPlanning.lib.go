package linear

import (
	"fmt"
)

func LinearPlanning(x1Name string, x2Name string, constraint1 int, constraint2 int, minOrMax string) {
	var (
		x1 int
		x2 int
	)

	if x1 < 0 {
		fmt.Println("X1이 음수입니다. Ignore.")

		return
	}

	if x2 < 0 {
		fmt.Println("X2은 음수입니다. Ignore.")

		return
	}

	constraintStatement(constraint1, constraint2, minOrMax)
}

func constraintStatement(constraint1 int, constraint2 int, minOrMax string) {
	fmt.Printf("%s z = %d x1 + %d x2", minOrMax, constraint1, constraint2)
}

func DecideMaxOrMin(maxAnswer string) (int, string) {
	var isMax int
	minOrMax := "Max"

	if maxAnswer == "y" {
		isMax = 1
	} else if maxAnswer == "n" {
		isMax = 0
		minOrMax = "Min"
	} else {
		fmt.Println("유효하지 않은 입력값")

		return -1, ""
	}

	return isMax, minOrMax
}
