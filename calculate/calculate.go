package calculate

import (
	"fmt"
	"strconv"
)

func Calcfinal() {
	var calcinfo CalcInfo

	calcinfo.number1 = Input("첫번째")
	calcinfo.number2 = Input("두번째")

	calcinfo.operator = Oper()

	calcinfo = Calc(calcinfo)

	fmt.Print("결과는 [" + strconv.Itoa(calcinfo.result) + "] 입니다.")

}

func Input(inputNum string) int {
	var number int
	fmt.Print(inputNum + " 숫자를 입력해주세요 : ")
	fmt.Scanln(&number)

	return number
}

func Oper() string {
	var oper string
	fmt.Print("어떻게 계산할까요? (+, -, *, /) : ")
	fmt.Scanln(&oper)

	return oper
}

type CalcInfo struct {
	operator string
	number1  int
	number2  int
	result   int
}

func Calc(calcinfo CalcInfo) CalcInfo {

	if calcinfo.operator == "+" {
		calcinfo.result = calcinfo.number1 + calcinfo.number2
	}

	if calcinfo.operator == "-" {
		calcinfo.result = calcinfo.number1 - calcinfo.number2
	}

	if calcinfo.operator == "*" {
		calcinfo.result = calcinfo.number1 * calcinfo.number2
	}

	if calcinfo.operator == "/" {
		calcinfo.result = calcinfo.number1 / calcinfo.number2
	}

	return calcinfo
}
