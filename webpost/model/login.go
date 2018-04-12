package model

// LoginRequest data type
type LoginRequest struct {
	Phone string `json:"phone"`
	Token string `json:"token"`
}

// LoginResponse data type
type LoginResponse struct {
	Token      string `json:"token"`
	ContractID string `json:"contratId"`
	Plan       string `json:"plan"`
}
