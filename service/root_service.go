package service

type RootService interface {
	Init()
	Play()
}

func NewRootService(maxNumber int) RootService {
	return &rootServiceImpl{
		maxNumber: maxNumber,
	}
}
