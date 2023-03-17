package fibonacci

import "fmt"

var numbersArray []int

func FibonacciNumbers() {
	var count int

	fmt.Println("피보나치 수열")
	fmt.Println("~까지의 피보나치 수열을 구하겠습니다. 마지막 숫자를 입력해 주세요.")

	fmt.Scanln(&count)

	fmt.Println("Final Number", count)

	fibo(count)
}

func fibo(count int) {
	firstNumber := 1

	numbersArray = append(numbersArray, firstNumber)

	for i := 1; i < count; i += 1 {
		fmt.Println("number i: ", i)

		if i == 1 {
			numbersArray = append(numbersArray, numbersArray[i]+0)

			fmt.Println("NumberArray: ", numbersArray)
		}

		numbersArray = append(numbersArray, numbersArray[i-1]+numbersArray[i-2])

		fmt.Println("Number Array: ", numbersArray)
	}

}
