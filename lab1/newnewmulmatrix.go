package main

import (
	"bufio"
	"fmt"
	"os"
)

var result [][]float64
var out *bufio.Writer

func summatrix1(matrix1, matrix2 [][]float64) [][]float64 {
	var result1 [][]float64
	z := len(matrix1)
	for k := 0; k < z; k++ {
		temp := make([]float64, z, z)
		result1 = append(result1, temp)
	}
	for i := 0; i < z; i++ {
		for j := 0; j < z; j++ {

			result1[i][j] += matrix2[i][j] + matrix1[i][j]
		}
	}
	return result1
}
func cutmatrix1(in [][]float64, starta, stopa, startb, stopb int, c chan [][]float64) {
	var temp [][]float64
	var z float64
	for i := starta; i < stopa; i++ {
		var tempp []float64
		for j := startb; j < stopb; j++ {
			z = in[i][j]
			tempp = append(tempp, z)
		}
		temp = append(temp, tempp)
	}
	c <- temp
}
func printmatrix1(res [][]float64) {
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			fmt.Fprintf(out, "%f ", res[i][j])
		}
		fmt.Fprintf(out, "\n")
	}

	fmt.Fprintf(out, "\n")
}
func calculatematrixnew1(matrix1, matrix2 [][]float64, c chan [][]float64) {

	var result1 [][]float64
	d := len(matrix1)
	for k := 0; k < d; k++ {
		temp := make([]float64, d, d)
		result1 = append(result1, temp)
	}

	for k := 0; k < d; k++ {
		for j := 0; j < d; j++ {
			for i := 0; i < d; i++ {
				a, b := 0.0, 0.0
				if i < len(matrix1) && k < len(matrix1[0]) {
					a = matrix1[i][k]
				}
				if k < len(matrix2) && j < len(matrix2[0]) {
					b = matrix2[k][j]
				}
				//println(i, j, k, a, b)
				result1[i][j] += a * b
			}
		}
	}
	c <- result1
	//printmatrix(result1)
}
func newcalculatematrix1(matrix1, matrix2 [][]float64) [][]float64 {

	var result1 [][]float64
	for k := 0; k < len(matrix1); k++ {
		temp := make([]float64, len(matrix2[0]), len(matrix2[0]))
		result1 = append(result1, temp)
	}

	A, B, C, D := make(chan [][]float64), make(chan [][]float64), make(chan [][]float64), make(chan [][]float64)
	go cutmatrix1(matrix1, 0, len(matrix1), 0, len(matrix1[0])/2, A)
	go cutmatrix1(matrix1, 0, len(matrix1), len(matrix1[0])/2, len(matrix1[0]), B)
	go cutmatrix1(matrix2, 0, len(matrix2)/2, 0, len(matrix2[0]), C)
	go cutmatrix1(matrix2, len(matrix2)/2, len(matrix2), 0, len(matrix2[0]), D)

	t1, t2 := make(chan [][]float64), make(chan [][]float64)

	//printmatrix1(A)

	//printmatrix1(C)
	go calculatematrixnew1(<-A, <-C, t1)

	go calculatematrixnew1(<-B, <-D, t2)
	m1, m2 := <-t1, <-t2

	result1 = summatrix1(m1, m2)

	return result1
}
func main() {
	in := bufio.NewReader(os.Stdin)
	out = bufio.NewWriter((os.Stdout))
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
	for k := 0; k < len(matrix1); k++ {
		temp := make([]float64, len(matrix2[0]), len(matrix2[0]))
		result = append(result, temp)
	}

	printmatrix1(newcalculatematrix1(matrix1, matrix2))
	out.Flush()
}
