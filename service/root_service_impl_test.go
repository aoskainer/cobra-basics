package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJudgeGuessNumber(t *testing.T) {
	service := rootServiceImpl{
		maxNumber: 10,
		target:    5, // テストのために特定のターゲット値を設定
	}

	cases := []struct {
		name     string
		guess    int
		expected judgementResult
	}{
		{
			name:  "範囲外の値(最小値より1小さい値)を入力した場合、範囲を明示して再入力を促すメッセージとfalseが返却される",
			guess: 0,
			expected: judgementResult{
				result:  false,
				message: "1から10までの数値を入力してください。",
			},
		},
		{
			name:  "範囲外の値(最大値より1大きい値)を入力した場合、範囲を明示して再入力を促すメッセージとfalseが返却される",
			guess: 11,
			expected: judgementResult{
				result:  false,
				message: "1から10までの数値を入力してください。",
			},
		},
		{
			name:  "範囲内かつtargetより小さい値を入力した場合、再入力を促すメッセージとfalseが返却される",
			guess: 1,
			expected: judgementResult{
				result:  false,
				message: "小さすぎます！もう一度試してください。",
			},
		},
		{
			name:  "範囲内かつtargetより大きい値を入力した場合、再入力を促すメッセージとfalseが返却される",
			guess: 10,
			expected: judgementResult{
				result:  false,
				message: "大きすぎます！もう一度試してください。",
			},
		},
		{
			name:  "範囲内かつtargetと同じ値を入力した場合、おめでとうメッセージとtrueが返却される",
			guess: 5,
			expected: judgementResult{
				result:  true,
				message: "おめでとうございます！当たりです！",
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := service.judgeGuessNumber(tt.guess)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
