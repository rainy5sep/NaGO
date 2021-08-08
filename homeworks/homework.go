package homeworks

import (
	"fmt"
	"math/rand"
	"strconv"
)

func Homework1() {
	star := "*"

	for i := 0; i < 9; i++ {

		if i >= 0 && i < 5 {

			for j := 0; j < i+1; j++ {
				fmt.Print(star)
			}

		} else {

			for j := 9; j > i; j-- {
				fmt.Print(star)
			}

		}

		fmt.Println()
	}
}

func Homework1a() {
	star := "*"
	for i := 0; i < 5; i++ {

		for j := 0; j < i+1; j++ {
			fmt.Print(star)

		}

		fmt.Println()

	}

	for k := 4; k > 0; k-- {

		for l := 0; l < k; l++ {
			fmt.Print(star)
		}

		fmt.Println()

	}

}

func Homework2() {
	star := "*"
	blank := " "

	for i := 5; i > 0; i-- {

		for j := 1; j < i; j++ {
			fmt.Print(blank)

		}

		for k := 5; k > i-1; k-- {
			fmt.Print(star)

		}

		fmt.Println()

	}

}

func Homework2a() {

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

func Homework3() {
	star := "*"
	blank := " "

	for i := 0; i < 6; i++ {

		for j := 0; j < i; j++ {
			fmt.Print(blank)

		}

		for k := 5; k > i; k-- {
			fmt.Print(star)

		}

		fmt.Println()

	}

}

func Homework3a() {
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

}

func Homework4() {

	for i := 0; i < 5; i++ {

		for j := 4; j > i; j-- {
			fmt.Print(" ")

		}

		for k := 0; k <= 2*i; k++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}

func Homework4a() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < i*2+1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func Homework5() {
	// 무작위 숫자 3개 만들고??
	cn1 := rand.Intn(8) + 1
	var cn2 int
	var cn3 int

	for {
		cn2 = rand.Intn(8) + 1

		if cn2 == cn1 {
			continue

		} else {
			break
		}
	}

	for {
		cn3 = rand.Intn(8) + 1

		if cn3 == cn2 || cn3 == cn1 {
			continue
		} else {
			break
		}
	}

	fmt.Println()
	//입력
	var num1 int
	var num2 int
	var num3 int

	for {
		fmt.Println("첫번째 숫자를 입력해주세요 (1부터 9까지) : ")
		fmt.Scanln(&num1)

		if num1 > 10 {
			fmt.Println("9보다 큰 숫자는 입력할 수 없습니다.")
			continue
		} else if num1 == 0 {
			fmt.Println("1보다 작은 숫자는 입력할 수 없습니다.")
			continue
		} else {
			break
		}
	}

	fmt.Println()

	for {
		fmt.Println("두번째 숫자를 입력해주세요 (1부터 9까지) : ")
		fmt.Scanln(&num2)
		if num1 == num2 {
			fmt.Println("이미 입력한 숫자는 입력할 수 없습니다.")
			fmt.Println()
			continue
		} else if num2 > 10 {
			fmt.Println("9보다 큰 숫자는 입력할 수 없습니다")
			fmt.Println()
			continue
		} else if num2 == 0 {
			fmt.Println("1보다 작은 숫자는 입력할 수 없습니다.")
			continue
		} else {
			break
		}
	}

	fmt.Println()

	for {
		fmt.Println("세번째 숫자를 입력해주세요 (1부터 9까지) : ")
		fmt.Scanln(&num3)
		if num3 == num1 || num3 == num2 {
			fmt.Println("이미 입력한 숫자는 입력할 수 없습니다.")
			fmt.Println()
			continue
		} else if num3 > 10 {
			fmt.Println("9보다 큰 숫자는 입력할 수 없습니다")
			fmt.Println()
			continue
		} else if num3 == 0 {
			fmt.Println("1보다 작은 숫자는 입력할 수 없습니다.")
			continue
		} else {
			break
		}
	}

	fmt.Println()
	fmt.Println(cn1, cn2, cn3)
	fmt.Println(num1, num2, num3)
	fmt.Println()

	if num1 == cn1 && num2 == cn2 && num3 == cn3 {
		fmt.Print("3 Strike ")
		return

	} else if (num1 == cn1 && num2 == cn2) || (num1 == cn1 && num3 == cn3) || (num2 == cn2 && num3 == cn3) {
		fmt.Print("2 Strike ")
	} else if num1 == cn1 {
		fmt.Print("1 Strike ")
	} else {
		fmt.Print("0 Strike ")
	}

	if num1 == cn2 || num1 == cn3 || num2 == cn1 || num2 == cn3 || num3 == cn1 || num3 == cn2 {
		fmt.Print("3 ball")
	} else if num1 == cn2 || num1 == cn3 || num2 == cn1 || num2 == cn3 {
		fmt.Print("2 ball")
	} else if num1 == cn2 || num1 == cn3 {
		fmt.Print("1 ball")
	} else {
		fmt.Print("0 ball")
	}
}

func Homework6() {
	var inputNum [3]int
	inputNum[0] = rand.Intn(8) + 1
	var tempNum int
	for i := 1; i < 3; i++ {
		tempNum = rand.Intn(8) + 1
		if tempNum != inputNum[i-1] {
			inputNum[i] = tempNum
		} else {
			i--
		}
	}
	fmt.Println(inputNum)
	var userNum [3]int
	for {
		for i := 0; i < 3; i++ {
			fmt.Println(strconv.Itoa(i+1) + "번 숫자를 입력해주세요 (1부터 9까지) : ")
			fmt.Scanln(&userNum[i])
			if userNum[i] == 0 || userNum[i] > 9 {
				fmt.Println("1부터 9까지의 숫자만 입력 가능합니다.")
				i--
			} else if i > 0 {
				if userNum[i] == userNum[i-1] {
					fmt.Println("이미 입력한 숫자는 입력할 수 없습니다.")
					i--
				} else if userNum[i] == userNum[0] {
					fmt.Println("이미 입력한 숫자는 입력할 수 없습니다.")
					i--
				}
			}
		}
		var strike int = 0
		var ball int = 0
		for i, v := range userNum {
			if v == inputNum[i] {
				strike++
				continue
			}
			for _, w := range inputNum {
				if w == userNum[i] {
					ball++
				}
			}
		}
		fmt.Println(strconv.Itoa(strike) + " Strike")
		fmt.Println(strconv.Itoa(ball) + " Ball")

		if strike == 3 {
			fmt.Println("삼진아웃! ㅊㅋㅊㅋ 게임 끝 ><")
			break
		}

		fmt.Println()
		fmt.Println("다시 시도해주세요")

	}
}
