package models

type TrxData struct{
	IsBroadcasted bool
	SenderBlockID string
	RecieverBlockID string
	Amount float64
	Sender string
	Reciever string
	SPrivateKey string
	RPrivateKey string
	SenderLastBlockAmount int
	RecieverLastBlockAmount int
}