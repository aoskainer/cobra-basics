package service

import (
	"testing"
)

func TestJudgeGuessNumber(t *testing.T) {
	service := guessGameServiceImpl{
		maxNumber: 10,
		target:    5, // テストのために特定のターゲット値を設定
	}

	tests := []struct {
		name  string
		guess int
		want  judgementResult
	}{
		{
			name:  "範囲外の値(最小値より1小さい値)を入力した場合、範囲を明示して再入力を促すメッセージとfalseが返却される",
			guess: 0, want: judgementResult{result: false, message: "1から10までの数値を入力してください。"},
		},
		{
			name:  "範囲外の値(最大値より1大きい値)を入力した場合、範囲を明示して再入力を促すメッセージとfalseが返却される",
			guess: 11, want: judgementResult{result: false, message: "1から10までの数値を入力してください。"},
		},
		{
			name:  "範囲内かつtargetより小さい値を入力した場合、再入力を促すメッセージとfalseが返却される",
			guess: 1, want: judgementResult{result: false, message: "小さすぎます！もう一度試してください。"},
		},
		{
			name:  "範囲内かつtargetより大きい値を入力した場合、再入力を促すメッセージとfalseが返却される",
			guess: 10, want: judgementResult{result: false, message: "大きすぎます！もう一度試してください。"},
		},
		{
			name:  "範囲内かつtargetと同じ値を入力した場合、おめでとうメッセージとtrueが返却される",
			guess: 5, want: judgementResult{result: true, message: "おめでとうございます！当たりです！"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := service.judgeGuessNumber(tt.guess)
			if got.result != tt.want.result || got.message != tt.want.message {
				t.Errorf("judgeGuessNumber(%d) = {%v, %q}; want {%v, %q}", tt.guess, got.result, got.message, tt.want.result, tt.want.message)
			}
		})
	}
}
