package service

import (
	"fmt"
	"math/rand"
)

type guessGameServiceImpl struct {
	maxNumber int
	target    int
}

func (service *guessGameServiceImpl) Init() {
	service.target = rand.Intn(service.maxNumber + 1)
}

func (service *guessGameServiceImpl) Play() {
	fmt.Printf("1から%dまでの数からランダムな値を選びました。当ててみてください。\n", service.maxNumber)

	for {
		guess := service.inputGuessNumber()
		judgementResult := service.judgeGuessNumber(guess)
		fmt.Println(judgementResult.message)
		if judgementResult.result {
			break
		}
	}
}

func (service *guessGameServiceImpl) inputGuessNumber() int {
	for {
		fmt.Print("あなたの予想: ")
		var guess int
		if _, err := fmt.Scanln(&guess); err != nil {
			fmt.Println("有効な数値を入力してください。")
			continue
		}
		return guess
	}
}

func (service *guessGameServiceImpl) judgeGuessNumber(guess int) judgementResult {
	if guess < 1 || guess > service.maxNumber {
		return judgementResult{
			result:  false,
			message: fmt.Sprintf("1から%dまでの数値を入力してください。", service.maxNumber),
		}
	}
	if guess < service.target {
		return judgementResult{
			result:  false,
			message: "小さすぎます！もう一度試してください。",
		}
	}
	if guess > service.target {
		return judgementResult{
			result:  false,
			message: "大きすぎます！もう一度試してください。",
		}
	}
	return judgementResult{
		result:  true,
		message: "おめでとうございます！当たりです！",
	}
}

type judgementResult struct {
	result  bool
	message string
}
