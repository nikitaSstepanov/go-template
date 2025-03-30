package middleware

type Middleware struct {
	auth AuthUseCase
}

func New(uc AuthUseCase) *Middleware {
	return &Middleware{
		auth: uc,
	}
}
