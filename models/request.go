package models

type Request struct {
	Action  string
	Message string
	Version string
}

type CreateWalletRequest struct {
	
	WalletName string
	PassPhrase string
	
}
