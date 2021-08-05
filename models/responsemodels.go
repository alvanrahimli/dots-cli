package models

type LoginResponse struct {
	Code    int    `json:"Code"`
	Message string `json:"Message"`
	Data    struct {
		ExpirationDate string `json:"expiration_date"`
		Token          string `json:"token"`
		User           struct {
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"user"`
	} `json:"Data"`
}
