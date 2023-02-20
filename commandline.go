package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check_length(command []string, num_arguments int, symbol string) bool {
	num_arguments += 1

	if len(command) < num_arguments {
		fmt.Printf("Too few arguments for `%s`\n", symbol)
		return false
	} else if len(command) > num_arguments {
		fmt.Printf("Too many arguments for `%s`\n", symbol)
		return false
	}

	return true
}

func parse_float(command []string) (float64, bool) {
	x, err_x := strconv.ParseFloat(command[1], 64)
	if err_x != nil {
		fmt.Printf("'%s' is not a float\n", command[1])
		return 0, false
	}

	return x, true
}

func parse_floats(command []string) (float64, float64, bool) {
	x, err_x := strconv.ParseFloat(command[1], 64)
	if err_x != nil {
		fmt.Printf("'%s' is not a float\n", command[1])
		return 0, 0, false
	}

	y, err_y := strconv.ParseFloat(command[2], 64)
	if err_y != nil {
		fmt.Printf("'%s' is not a float\n", command[2])
		return 0, 0, false
	}

	return x, y, true
}

func parse_int(command []string) (int64, bool) {
	x, err_x := strconv.ParseInt(command[1], 10, 64)
	if err_x != nil {
		fmt.Printf("'%s' is not a float\n", command[1])
		return 0, false
	}

	return x, true
}

const PHI = 1.618033988749894848205
const PSI = -0.6180339887498948482046
const SQRT5 = 2.236067977499789696409

func fibn(n int64) int64 {
	power := float64(n)
	num1 := math.Pow(PHI, power)
	num2 := math.Pow(PSI, power)

	return int64((num1 - num2) / SQRT5)
}

func fibr(n int64) int64 {
	switch {
	case n <= 0:
		return 0
	case n == 1:
		return 1
	default:
		return fibr(n-1) + fibr(n-2)
	}
}

func main() {
	running := true

	for running {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("An error occurred!")
			continue
		}

		input = strings.TrimSpace(input)
		command := strings.Split(input, " ")

		if len(command) == 0 {
			continue
		}

		switch command[0] {
		case "quit":
			running = false
		case "help":
			fmt.Println("`quit`: quit the program")
			fmt.Println("`echo` ..: echo the arguments back")
			fmt.Println("`+` x y: print the result of x + y")
			fmt.Println("`-` x y: print the result of x - y")
			fmt.Println("`*` x y: print the result of x * y")
			fmt.Println("`/` x y: print the result of x / y")
			fmt.Println("`fibn` n: calculate the n'th element in the fibonnaci sequence with math")
			fmt.Println("`fibr` n: calculate the n'th element in the fibonnaci sequence with recursion")
		case "echo":
			fmt.Println(strings.Join(command[1:], " "))
		case "+":
			if check_length(command, 2, "+") {
				x, y, success := parse_floats(command)

				if success {
					fmt.Printf("%g + %g = %g\n", x, y, x+y)
				}
			}
		case "-":
			if check_length(command, 2, "-") {
				x, y, success := parse_floats(command)

				if success {
					fmt.Printf("%g - %g = %g\n", x, y, x-y)
				}
			}
		case "*":
			if check_length(command, 2, "*") {
				x, y, success := parse_floats(command)

				if success {
					fmt.Printf("%g * %g = %g\n", x, y, x*y)
				}
			}
		case "/":
			if check_length(command, 2, "/") {
				x, y, success := parse_floats(command)

				if y == 0 {
					fmt.Println("Can not divide by 0")
				} else if success {
					fmt.Printf("%g / %g = %g\n", x, y, x/y)
				}
			}
		case "sin":
			if check_length(command, 1, "sin") {
				x, success := parse_float(command)

				if success {
					fmt.Printf("sin %g = %g\n", x, math.Sin(x))
				}
			}
		case "cos":
			if check_length(command, 1, "cos") {
				x, success := parse_float(command)

				if success {
					fmt.Printf("cos %g = %g\n", x, math.Cos(x))
				}
			}
		case "exp":
			if check_length(command, 1, "exp") {
				x, success := parse_float(command)

				if success {
					fmt.Printf("exp %g = %g\n", x, math.Exp(x))
				}
			}

		case "fibn":
			if check_length(command, 1, "fibn") {
				x, success := parse_int(command)

				if success {
					if x > 50 {
						fmt.Printf("%d is too large, so it can be inaccurate or wrong\n", x)
					} else if x < 0 {
						fmt.Printf("%d is too small, so it can be inaccurate or wrong\n", x)
					}

					fmt.Printf("fibn %d = %d\n", x, fibn(x))
				}
			}
		case "fibr":
			if check_length(command, 1, "fibr") {
				x, success := parse_int(command)

				if success {
					if x > 40 {
						fmt.Printf("%d is too large\n", x)
						continue
					} else if x < 0 {
						fmt.Printf("%d can't be negative\n", x)
						continue
					}

					fmt.Printf("fibr %d = %d\n", x, fibr(x))
				}
			}
		default:
			fmt.Printf("%s is not a valid command, type `help` for a list of commands\n", command[0])
		}
	}
}
