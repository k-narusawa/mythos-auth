package services

type RegisterUserService struct{}

func NewRegisterUserInteractor() *RegisterUserService {
	return new(RegisterUserService)
}

func (r *RegisterUserService) Invoke(in *RegisterUserInputData) (out *RegisterUserOutputData, err error) {
	return
}
