package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter((os.Stdout))
	var a, b, c, d int
	rand.Seed(time.Now().UnixNano())
	fmt.Fscanf(in, "%d %d %d %d\n", &a, &b, &c, &d)
	fmt.Fprintf(out, "%d %d\n", a, b)
	for ; a > 0; a-- {
		for i := 0; i < b; i++ {
			fmt.Fprintf(out, "%f ", rand.Float64()*100)
		}

		fmt.Fprintf(out, "\n")
	}
	fmt.Fprintf(out, "%d %d\n", c, d)
	for ; c > 0; c-- {
		for i := 0; i < d; i++ {
			fmt.Fprintf(out, "%f ", rand.Float64()*100)
		}
		fmt.Fprintf(out, "\n")
	}
	out.Flush()
}
