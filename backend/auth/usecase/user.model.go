package usecase

// SignInInput represents the input for a sign-in request.
type SignInInput struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
