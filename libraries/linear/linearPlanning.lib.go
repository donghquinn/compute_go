package linear

import (
	"fmt"
)

func LinearPlanning(x1Name string, x2Name string, constraint1 int, constraint2 int, minOrMax string, goal1 int, goal2 int) {
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

	constraintStatement(constraint1, constraint2, goal1, goal2, minOrMax)
}

func constraintStatement(constraint1 int, constraint2 int, goal1 int, goal2 int, minOrMax string) {
	fmt.Println("제약조건식")

	fmt.Printf("%d x1 + %d x2 <= %d ", constraint1, constraint2, goal1)
	fmt.Printf("%d x1 + %d x2 <= %d ", constraint1, constraint2, goal2)

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
