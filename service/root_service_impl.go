package service

import (
	"fmt"
	"math/rand"

	utility "github.com/aoskainer/cobra-basics/utility/nocover"
)

type rootServiceImpl struct {
	maxNumber     int
	target        int
	userInputUtil utility.UserInputUtility
}

func (service *rootServiceImpl) Init() {
	service.target = rand.Intn(service.maxNumber + 1)
	service.userInputUtil = utility.NewUserInputUtility()
}

func (service *rootServiceImpl) Play() {
	fmt.Printf("1から%dまでの数からランダムな値を選びました。当ててみてください。\n", service.maxNumber)

	for {
		guess := service.userInputUtil.InputGuessNumber()

		judgementResult := service.judgeGuessNumber(guess)

		fmt.Println(judgementResult.message)
		if judgementResult.result {
			break
		}
	}
}

func (service *rootServiceImpl) judgeGuessNumber(guess int) judgementResult {
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
