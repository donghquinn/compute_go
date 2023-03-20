package linear

import "fmt"

var (
	x1Name string
	x2Name string

	maxAnswer string
	// x1        float64
	// x2        float64
)

func LinearPlanning() {
	var isMax int

	fmt.Println("선형 계획법")

	fmt.Println("최대화 문제인가요? y / n")
	fmt.Scanln(&maxAnswer)

	isMax = decideMaxOrMin(maxAnswer, isMax)

	fmt.Println("Is Max: ", isMax)

	fmt.Println("첫번째 의사결정 변수 이름을 알려주세요")
	fmt.Println("예시: 제품 갑의 생산량")
	fmt.Scanln(&x1Name)

	fmt.Println("두번째 의사결정 변수 이름을 알려주세요")
	fmt.Println("예시: 제품 을의 생산량")
	fmt.Scanln(&x2Name)
}

func decideMaxOrMin(maxAnswer string, isMax int) int {
	if maxAnswer == "y" {
		isMax = 1
	} else if maxAnswer == "n" {
		isMax = 0
	} else {
		fmt.Println("유효하지 않은 입력값")
	}

	return isMax
}
