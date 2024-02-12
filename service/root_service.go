package service

import (
	"math/rand"

	utility "github.com/aoskainer/cobra-basics/utility/nocover"
)

type RootService interface {
	Play()
}

func NewRootService(maxNumber int) RootService {
	return &rootServiceImpl{
		maxNumber:     maxNumber,
		target:        rand.Intn(maxNumber + 1),
		userInputUtil: utility.NewUserInputUtility(),
	}
}
