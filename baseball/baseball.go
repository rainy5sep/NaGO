package baseball

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var gameinfo GameInfo

func Bbgame() {

	gameinfo.ComNum = Com()
	fmt.Println(gameinfo.ComNum)

	fmt.Println("야구 게임 시작!")
	fmt.Println("숫자 3개를 골라서 입력해주세요!")
	fmt.Println()

	for {

		gameinfo.userNum = User()
		StrikeOrBall()

		fmt.Println()
		fmt.Println(strconv.Itoa(gameinfo.strike) + " Strike")
		fmt.Println(strconv.Itoa(gameinfo.ball) + " Ball")

		if gameinfo.strike == 3 {
			fmt.Println()
			fmt.Println("삼진아웃! ㅊㅋㅊㅋ 게임 끝 ><")
			fmt.Println()
			break
		}

		fmt.Println()
		fmt.Println("다시 시도해주세요")
		fmt.Println()

	}

}

type GameInfo struct {
	strike  int
	ball    int
	userNum [3]int
	ComNum  [3]int
}

func User() [3]int {

	var userNum [3]int
	var number int

	for i := 0; i < 3; i++ {

		fmt.Print(strconv.Itoa(i+1) + "번 숫자를 입력해주세요 (1부터 9까지) : ")
		fmt.Scanln(&number)

		userNum[i] = number

		if userNum[i] <= 0 || userNum[i] > 9 {
			fmt.Println("1부터 9까지의 숫자만 입력 가능합니다.")
			i--
		} else if i > 0 {
			if userNum[i] == userNum[i-1] || userNum[i] == userNum[0] {
				fmt.Println("이미 입력한 숫자는 입력할 수 없습니다.")
				i--
			}

		}

	}

	return userNum

}

func Com() [3]int {

	var ComNum [3]int

	rand.Seed(time.Now().UnixNano())
	ComNum[0] = rand.Intn(8) + 1

	var tempNum int

	for i := 1; i < 3; i++ {

		rand.Seed(time.Now().UnixNano())
		tempNum = rand.Intn(8) + 1

		if tempNum != ComNum[i-1] && tempNum != ComNum[0] {
			ComNum[i] = tempNum
		} else {
			i--
		}
	}

	return ComNum
}

func StrikeOrBall() (int, int) {

	gameinfo.strike = 0
	gameinfo.ball = 0

	for i, v := range gameinfo.userNum {
		if v == gameinfo.ComNum[i] {
			gameinfo.strike++
			continue
		}

		for _, w := range gameinfo.ComNum {
			if w == gameinfo.userNum[i] {
				gameinfo.ball++
			}
		}
	}

	return gameinfo.strike, gameinfo.ball
}
