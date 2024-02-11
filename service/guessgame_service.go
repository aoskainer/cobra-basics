package service

type GuessGameService interface {
	Init()
	Play()
}

func NewGuessGameService(maxNumber int) GuessGameService {
	return &guessGameServiceImpl{
		maxNumber: maxNumber,
	}
}
