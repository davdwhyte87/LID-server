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

type GetBalanceReq struct {
	WalletName string
}
type TransferReq struct{
	Sender string
	Receiver string 
	Amount string

}
