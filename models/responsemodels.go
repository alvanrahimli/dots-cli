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

type GetPackagesResponse struct {
	Code    int    `json:"Code"`
	Message string `json:"Message"`
	Data    struct {
		Packages []struct {
			Id          int    `json:"Id"`
			Name        string `json:"Name"`
			Version     string `json:"Version"`
			ArchiveName string `json:"ArchiveName"`
			UserId      int    `json:"UserId"`
		} `json:"Packages"`
	} `json:"Data"`
}
