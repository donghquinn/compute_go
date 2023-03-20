package main

import (
	"fmt"

	fibo "github.com/donghquinn/schoolwork/libraries/fibonacci"
	mean "github.com/donghquinn/schoolwork/libraries/mean"
	variance "github.com/donghquinn/schoolwork/libraries/variance"
)

var (
	method      string
	methodArray []string
)

func main() {
	methodArray = append(methodArray, "피보나치", "분산", "평균")

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
