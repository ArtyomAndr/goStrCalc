package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	/*	var txt1 string
		var txt2 string
		var operator string

		fmt.Scan(&txt1)
		fmt.Scan(&operator)

		// "x" + "y"
		// "a" + "b", "a" - "b", "a" * b, "a" / b
		// Калькулятор должен принимать на вход числа от 1 до 10 включительно, не более,
		// и строки длинной не более 10 символов.

		// scan := bufio.NewReader(os.Stdin)

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		txt1 = scanner

		txt1 = bufio.NewReader(os.Stdin)
		txt2 = bufio.NewReader(os.Stdin)

		// var txt2 = scanner.Reader(os.Stdin) // игнор ""
	*/

	var txt1 string
	var txt2 string
	var operator string
	fmt.Scanf("%q%s", &txt1, &operator)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	txt2 = scanner.Text()

	txt1Run := []rune(txt1)
	txt2Run := []rune(txt2)

	if len(txt1Run) <= 10 && len(txt2Run) <= 10 {

		txt2Run1 := rune(txt2[0])
		txt2Run2 := rune(txt2[len(txt2)-1])

		switch operator {

		case "+":

			if txt2Run1 != '"' && txt2Run2 != '"' {
				panic("не верное значение первого аргумента")
			} else {
				txt2 = txt2[1 : len(txt2)-1]
				fmt.Printf("\"%s%s\"", txt1, txt2)
			}

		case "-":
			if txt2Run1 != '"' && txt2Run2 != '"' {
				panic("не верное значение первого аргумента")
			} else {
				txt2 = txt2[1 : len(txt2)-1]
				txt1 = strings.ReplaceAll(txt1, txt2, "")
				fmt.Printf("%q", txt1)
			}
		case "*":
			if isStrInt, _ := strconv.Atoi(txt2); isStrInt >= 1 && isStrInt <= 10 {

				txt1Run = []rune(strings.Repeat(txt1, isStrInt))

				if len(txt1Run) > 40 {
					txt1Run = append(txt1Run[:40])
					fmt.Printf("\"%s...\"", string(txt1Run))
				} else {
					fmt.Printf("%q", string(txt1Run))
				}
			} else {
				panic("не верное значение второго аргумента")
			}
			// ...
		case "/":
			if isStrInt, _ := strconv.Atoi(txt2); isStrInt >= 1 && isStrInt <= 10 {
				isStrInt, _ := strconv.Atoi(txt2)
				txt1Run = txt1Run[1 : len(txt1Run)/isStrInt]

				fmt.Printf("%q", string(txt1Run))
			} else {
				panic("не верное значение второго аргумента")
			}
			// ...
		default:
			panic("не верное выражение")
		}
	}
}
