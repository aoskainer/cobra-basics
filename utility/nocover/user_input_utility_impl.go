package utility

import "fmt"

type userInputUtilityImpl struct{}

func (utility *userInputUtilityImpl) InputGuessNumber() int {
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
