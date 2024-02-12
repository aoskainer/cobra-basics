package utility

type UserInputUtility interface {
	InputGuessNumber() int
}

func NewUserInputUtility() UserInputUtility {
	return &userInputUtilityImpl{}
}
