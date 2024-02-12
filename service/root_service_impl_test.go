package service

import (
	"testing"

	mock_utility "github.com/aoskainer/cobra-basics/utility/nocover/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestPlay(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserInputUtil := mock_utility.NewMockUserInputUtility(ctrl)
	gomock.InOrder(
		// 1回目の入力。4を入力して間違えたと仮定する。
		mockUserInputUtil.EXPECT().InputGuessNumber().Return(4).Times(1),
		// 2回目の入力。5を入力して正解したと仮定する。
		mockUserInputUtil.EXPECT().InputGuessNumber().Return(5).Times(1),
	)

	service := rootServiceImpl{maxNumber: 10}
	service.target = 5
	service.userInputUtil = mockUserInputUtil

	assert.NotPanics(t, func() {
		service.Play()
	})
}

func TestJudgeGuessNumber(t *testing.T) {
	service := rootServiceImpl{maxNumber: 10}
	service.target = 5

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
