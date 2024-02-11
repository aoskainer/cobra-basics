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
		var guess int
		fmt.Print("あなたの予想: ")
		if _, err := fmt.Scanln(&guess); err != nil {
			fmt.Println("有効な数値を入力してください。")
			continue
		}

		if guess < 1 || guess > service.maxNumber {
			fmt.Printf("1から%dまでの数値を入力してください。\n", service.maxNumber)
			continue
		}

		if guess < service.target {
			fmt.Println("小さすぎます！もう一度試してください。")
		} else if guess > service.target {
			fmt.Println("大きすぎます！もう一度試してください。")
		} else {
			fmt.Println("おめでとうございます！当たりです！")
			break
		}
	}
}
