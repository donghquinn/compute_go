package mean

import "fmt"

func Mean(parameterArray []float64, count float64) float64 {
	total := 0.0
	fmt.Println()

	for i := 0; i < len(parameterArray); i += 1 {
		// fmt.Println("Each Paramteres", parameterArray[i])
		total += parameterArray[i]
	}

	fmt.Println("총합: ", total)
	fmt.Println("전체 개수: ", count)

	mean := total / count

	fmt.Println("평균: ", mean)

	return mean
}
