// TODO https://pkg.go.dev/gonum.org/v1/gonum/optimize/convex/lp

package main

import (
	"fmt"

	fibo "github.com/donghquinn/schoolwork/libraries/fibonacci"
	linear "github.com/donghquinn/schoolwork/libraries/linear"
	mean "github.com/donghquinn/schoolwork/libraries/mean"
	variance "github.com/donghquinn/schoolwork/libraries/variance"
)

var (
	method      string
	methodArray []string
)

func main() {
	methodArray = append(methodArray, "피보나치", "분산", "평균", "선형계획법")

	fmt.Println("연산 종류: ", methodArray)
	fmt.Println("위 목록 중 사용할 연산을 선택해 주세요")
	fmt.Scanln(&method)

	decideMethod(method)
}

func decideMethod(method string) {
	if method == "피보나치" {
		fibo.FibonacciNumbers()
	} else if method == "분산" {
		callVarianceAndDeviance()
	} else if method == "평균" {
		callMean()
	} else if method == "선형계획법" {
		callLinearPlanning()
	}
}

func callVarianceAndDeviance() {
	var (
		parameterArray []float64
		parameter      float64
		count          float64
	)

	fmt.Println("'분산'을 골라주셨습니다. 분산에 사용할 값을 하나하나 입력하여 주세요.")

	for {
		fmt.Scanln(&parameter)

		if parameter == -1 {
			fmt.Println("Detected Closing Signal: -1")

			variance, deviance := variance.VarianceAndDeviance(parameterArray, count)

			fmt.Println("분산: ", variance)
			fmt.Println("표준편차: ", deviance)

			break
		}

		parameterArray = append(parameterArray, parameter)

		fmt.Println(parameterArray)

		count += 1
	}
}

func callMean() {
	var (
		parameter      float64
		parameterArray []float64
		count          float64
	)

	fmt.Println("'평균'을 골라주셨습니다. 분산에 사용할 값을 하나하나 입력하여 주세요.")

	for {
		fmt.Scanln(&parameter)

		if parameter == -1 {
			fmt.Println("Detected Closing Signal: -1")

			mean := mean.Mean(parameterArray, count)

			fmt.Println("평균: ", mean)

			break
		}

		parameterArray = append(parameterArray, parameter)

		fmt.Println(parameterArray)

		count += 1
	}
}

func callLinearPlanning() {

	var (
		x1Name       string
		x2Name       string
		maxAnswer    string
		constraints1 int
		constraints2 int
		goal1        int
		goal2        int
	)

	fmt.Println("선형 계획법")

	fmt.Println("최대화 문제인가요? y / n")
	fmt.Scanln(&maxAnswer)

	isMax, minOrMax := linear.DecideMaxOrMin(maxAnswer)

	if isMax == -1 {
		return
	}

	// TODO 조건 받기. 조건 제약 개수 만큼 파라미터 입력 등..

	fmt.Println("Is Max: ", isMax)
	fmt.Println("Min or Max: ", minOrMax)

	fmt.Println("첫번째 의사결정 변수 이름을 알려주세요")
	fmt.Println("예시: 제품 갑의 생산량")
	fmt.Scanln(&x1Name)

	fmt.Println("두번째 의사결정 변수 이름을 알려주세요")
	fmt.Println("예시: 제품 을의 생산량")
	fmt.Scanln(&x2Name)

	fmt.Println("첫번째 목적변수를 설정해 주세요")
	fmt.Scanln(&goal1)

	fmt.Println("두번째 목적변수를 설정해 주세요")
	fmt.Scanln(&goal2)

	fmt.Println("조건1을 입력해 주세요")
	fmt.Scanln(&constraints1)

	fmt.Println("조건2를 입력해 주세요")
	fmt.Scanln(&constraints2)

	linear.LinearPlanning(x1Name, x2Name, constraints1, constraints2, minOrMax, goal1, goal2)
}
