package models

type Request struct {
	Action  string
	Message string
	Version string
}

type CreateWalletRequest struct {
	Address       string
	PassPhrase    string
	IsBroadcasted bool 
}

type GetBalanceReq struct {
	WalletName string
}
type TransferReq struct {
	Sender        string
	Receiver      string
	Amount        string
	IsBroadcasted bool
}
