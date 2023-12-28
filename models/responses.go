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
	Balance int64
}
