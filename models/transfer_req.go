package models


// TransferReq ... 
type TransferReq struct{
	SenderPrivateKey string
	SenderAddress string
	RecieverAddress string
	RecieverPrivateKey string
	Amount int
}