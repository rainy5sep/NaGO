package examples

import (
	"fmt"
	"strconv"
)

func Example1() {
	var name string

	for i := 0; i < 10; i++ {

		fmt.Print("이름 입력 : ")
		fmt.Scanln(&name)

		if name == "나래" {
			name += "천사"
		} else {
			name += "악마"
		}

		fmt.Println(name)
	}

}

func Example2() {
	star := "*"
	for i := 0; i < 5; i++ {

		for j := 0; j < i+1; j++ {
			fmt.Print(star)
		}

		fmt.Println()

	}
}

func Example3() {
	star := "*"
	for i := 5; i > 0; i-- {

		for j := 0; j < i; j++ {
			fmt.Print(star)
		}

		fmt.Println()

	}
}

func Example4() {
	// 숫자 1을 입력받는다.
	var number1 int
	fmt.Print("첫번째 숫자를 입력해주세요 : ")
	fmt.Scanln(&number1)

	// 숫자 2를 입력받는다.
	var number2 int
	fmt.Print("두번째 숫자를 입력해주세요 : ")
	fmt.Scanln(&number2)

	// 연산자를 입력받는다. (+-*/)
	var operator string
	fmt.Print("어떻게 계산할까요? (+, -, *, /) : ")
	fmt.Scanln(&operator)

	// 계산을 한다.
	var result int

	if operator == "+" {
		result = number1 + number2
	}

	if operator == "-" {
		result = number1 - number2
	}

	if operator == "*" {
		result = number1 * number2
	}

	if operator == "/" {
		result = number1 / number2
	}

	// 출력을 한다.
	fmt.Print("결과는 [")
	fmt.Print(result)
	fmt.Print("] 입니다.")
}

func Example5() {

	/*var poo string
	poo = "안녕"

	var koo int // int 양수/음수 모두 표현 가능
	koo = -3

	var koo2 uint // unit는 양수만 표현이 됨
	koo2 = 3

	var loo float32
	loo = 1.2

	var yoo bool
	yoo = true*/

	/*var inputNum int

	fmt.Println("Hello, world")
	fmt.Scanln(&inputNum)

	if inputNum > 10 {
		fmt.Println("10보다 크지롱")
	} else {
		fmt.Println("아니지롱")
	} */

	// 첫번째 숫자 입력하세요

	/*var num1 int
	fmt.Print("첫번째 숫자를 입력하세요 : ")
	fmt.Scanln(&num1)

	var num2 int
	fmt.Print("두번째 숫자를 입력하세요 : ")
	fmt.Scanln(&num2)

	var oper string
	fmt.Print("연산자를 입력하세요 ( +, -, *, /) : ")
	fmt.Scanln(&oper)

	fmt.Print("결과는 ")

	if oper == "+" {
		fmt.Print(num1 + num2)
	}

	if oper == "-" {
		fmt.Print(num1 - num2)
	}

	if oper == "*" {
		fmt.Print(num1 * num2)
	}

	if oper == "/" {
		fmt.Print(num1 / num2)
	}

	fmt.Print(" 입니다.")*/

	for i := 0; i < 5; i++ {

		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}

		fmt.Println("")

	}

	fmt.Println("")

	for i := 5; i > 0; i-- {

		for j := 0; j < i; j++ {
			fmt.Print("*")
		}

		fmt.Println()

	}

	fmt.Println("")

	for i := 5; i > 0; i-- {

		for j := 1; j < i; j++ {
			fmt.Print(" ")

		}

		for k := 5; k > i-1; k-- {
			fmt.Print("*")

		}

		fmt.Println()

	}

	fmt.Println("")

	fmt.Println("")

	for i := 0; i < 6; i++ {

		for j := 0; j < 5; j++ {

			if j >= i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	for i := 0; i < 5; i++ {

		for j := 0; j < 5; j++ {

			if j < 5-i-1 {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}

		fmt.Println()
	}

}

func Example6() {

	for i := 2; i < 10; i++ {

		for j := 1; j < 10; j++ {
			fmt.Print(strconv.Itoa(i) + "X" + strconv.Itoa(j) + "=" + strconv.Itoa(i*j) + " ")
		}

		fmt.Println()

	}
}

func Example7() {

	// var card1 []string
	// var card2 []string
	// var card3 []string

	/*var vanila [3][]string

	var str string
	var str2 string
	var str3 string

	fmt.Scanln(&str)
	fmt.Scanln(&str2)
	fmt.Scanln(&str3)

	arr := []rune(str)
	arr2 := []rune(str2)
	arr3 := []rune(str3)

	vanila[0] = arr
	vanila[1] = arr2
	vanila[2] = arr3

	for i := 0; i < len(arr); i++ {
		fmt.Println(string(arr[i]))
	}

	// vanila[0] = card1
	// vanila[1] = card2
	// vanila[2] = card3

	// fmt.Println("카드값 입력받기")

	// fmt.Println(vanila[0], vanila[1], vanila[2])*/

}
