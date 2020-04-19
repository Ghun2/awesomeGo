package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	numbers := MakeNumbers()
	cnt := 0
	for {
		cnt++
		inputNumbers := InputNumbers()

		strike, ball := CompareNumbers(numbers, inputNumbers)
		PrintResult(strike, ball)
		if IsGameEnd(strike) {
			break
		}
	}
	fmt.Printf("%d 번만에 정답을 맞췄습니다.\n", cnt)

}


func MakeNumbers() [3]int {
	rn := [3]int{}
	for i, value := range rand.Perm(10)[:3] {rn[i] = value}
	return rn
}

func InputNumbers() [3]int {
	in := [3]int{}

	for {
		var no string
		_, err := fmt.Scanf("%s\n", &no)
		if err != nil {
			fmt.Println("잘못 입력했음.")
			continue
		}
		if len(no) < 3 {
			fmt.Println("3자리를 입력")
			continue
		}
		in[0], _ = strconv.Atoi(string(no[0]))
		in[1], _ = strconv.Atoi(string(no[1]))
		in[2], _ = strconv.Atoi(string(no[2]))

		if in[0] == in[1] || in[0] == in[2] || in[1] == in[2] {
			fmt.Println("서로 다른 3자리 입력")
			in = [3]int{}
			continue
		}

		break
	}

	return in
}

func CompareNumbers(rnum , inum [3]int) (int, int) {
	strike := 0
	ball := 0
	for i:=0;i<3;i++{
		if rnum[i] == inum[i] {
			strike++
			continue
		}
		for j:=0;j<3;j++{
			if rnum[i] == inum[j] {
				ball++
			}
		}
	}
	return strike, ball
}

func PrintResult(strike, ball int) {
	fmt.Printf("strike : %d, ball : %d\n", strike, ball)
}

func IsGameEnd(strike int) bool {
	if strike == 3 {
		return true
	}
	return false
}