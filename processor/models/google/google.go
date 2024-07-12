package googleModels

type LoginRequest struct {
	RedirectURL string `json:"redirect_url"`
}

type User struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}
