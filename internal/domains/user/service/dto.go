package service

type (
	SignInDTO struct {
		Username string
		Password string
	}

	CreateDTO struct {
		Username string
		Password string
		Email    string
	}

	UpdateDTO struct {
		Username string
		Email    string
	}
)
