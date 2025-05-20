package response

type Login struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SuccessAuthResponse struct {
	Meta
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
}
