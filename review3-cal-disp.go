package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func input(message []string, total int) []float64 {

	inputReader := bufio.NewReader(os.Stdin)
	var arrSlice []float64
	count := 0

	for {
		fmt.Println(message[count])
		input, _ := inputReader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		num, _ := strconv.ParseFloat(input, 64)
		arrSlice = append(arrSlice, num)

		if count == total {
			break
		}
		count += 1
	}

	return arrSlice
}

func GenDisplaceFn(v0 float64, a float64, s0 float64) func(t float64) float64 {
	return func(t float64) float64 {
		//		return ((1.0 / 2.0) * a * t * t) + v0*t + s0
		return (a * t * t / 2) + v0*t + s0
	}
}

func main() {
	res := input([]string{"Enter Velocity", "Enter Acceleration", "Enter Displacement"}, 2)
	// for debug :
	fmt.Println(res)
	cal_disp := GenDisplaceFn(res[0], res[1], res[2])
	// for debug :
	fmt.Println(cal_disp)
	t := input([]string{"Enter Time"}, 0)
	// for debug :
	fmt.Println("t=", t, "t[0]=", t[0])

	fmt.Println("Displacement 1: ", cal_disp(1))
	fmt.Println("Displacement 2: ", cal_disp(2))
	fmt.Println("Displacement 10: ", cal_disp(10))
	fmt.Println("Displacement ", t[0], ":", cal_disp(t[0]))
}
