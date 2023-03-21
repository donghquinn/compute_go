package libraries

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func MakeMetrics(row int, column int, data []float64) *mat.Dense {
	dense := mat.NewDense(row, column, data)

	fmt.Println("New Dense: ", dense)

	rows, columns := mat.Matrix.Dims(dense)

	lookAroundMatricsElement(dense, rows, columns)
	fmt.Println("Rows: ", rows)
	fmt.Println("Columns: ", columns)

	return dense
}

func lookAroundMatricsElement(dense *mat.Dense, rows int, columns int) {

	for i := 0; i < rows; i += 1 {
		for j := 0; j < columns; j += 1 {
			value := mat.Matrix.At(dense, i, j)

			fmt.Println("Matrics Element Value: ", value)
		}
	}

}
