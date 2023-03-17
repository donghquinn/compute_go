package variance

import (
	"fmt"
	"math"

	lib "github.com/donghquinn/schoolwork/libraries/mean"
)

func VarianceAndDeviance(parameterArray []float64, count float64) (float64, float64) {
	mean := lib.Mean(parameterArray, count)

	var variance float64

	varianceTotal := 0.0

	for i := 0; i < len(parameterArray); i += 1 {
		varianceTotal += (parameterArray[i] - mean) * (parameterArray[i] - mean)
	}
	fmt.Println("Total Count: ", count)

	variance = (varianceTotal / count)
	deviance := math.Sqrt(variance)

	return variance, deviance
}
