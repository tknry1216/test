package usecase

type createAccountUsecase struct {
}

func NewCreateAccountUsecase() *createAccountUsecase {
	return &createAccountUsecase{}
}

type CreateAccountRequestDTO struct {
	ExternalAccountID string
	TenantID          string
}

type CreateAccountResponseDTO struct {
	ID string
}

func (uc *createAccountUsecase) Execute(req *CreateAccountRequestDTO) (*CreateAccountResponseDTO, error) {
	return &CreateAccountResponseDTO{ID: "1234567890"}, nil
}
