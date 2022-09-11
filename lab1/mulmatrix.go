package main

import (
	"bufio"
	"fmt"
	"os"
)

func calculatematrix(matrix1, matrix2 [][]float64) [][]float64 {
	var result [][]float64
	for k := 0; k < len(matrix1); k++ {
		temp := make([]float64, len(matrix2[0]), len(matrix2[0]))
		result = append(result, temp)
	}
	for k := 0; k < len(matrix1); k++ {
		for j := 0; j < len(matrix2[0]); j++ {
			for i := 0; i < len(matrix2); i++ {
				result[k][j] += matrix1[k][i] * matrix2[i][j]
			}
		}
	}
	return result
}
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter((os.Stdout))
	var a, b, c, d int
	var matrix1, matrix2 [][]float64
	var elem float64
	fmt.Fscanf(in, "%d %d\n", &a, &b)
	for ; a > 0; a-- {
		var temp []float64
		for i := 0; i < b; i++ {
			fmt.Fscanf(in, "%f", &elem)
			temp = append(temp, elem)
		}
		matrix1 = append(matrix1, temp)
		fmt.Fscanf(in, "\n")
	}

	fmt.Fscanf(in, "%d %d\n", &c, &d)
	for ; c > 0; c-- {
		var temp []float64
		for i := 0; i < d; i++ {
			fmt.Fscanf(in, "%f", &elem)
			temp = append(temp, elem)
		}
		matrix2 = append(matrix2, temp)
		fmt.Fscanf(in, "\n")
	}
	if len(matrix2) != len(matrix1[0]) {
		fmt.Fprintf(out, "bad matrix\n")
		out.Flush()
		return
	}
	res := calculatematrix(matrix1, matrix2)
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[0]); j++ {
			fmt.Fprintf(out, "%f ", res[i][j])
		}
		fmt.Fprintf(out, "\n")
	}

	fmt.Fprintf(out, "\n")

	out.Flush()
}
