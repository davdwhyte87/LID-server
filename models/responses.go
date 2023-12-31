package models

type CreateWalletResponse struct {
	//PrivateKey string
	PublicKey string
}

type ErrorResponse struct {
	Code    int
	Message string
}

type GetBalanceResp struct {
	Code    int
	Balance float64
}

type GeneralResponse struct {
	Code    int
	Message string
}
