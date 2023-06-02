package models

type WorkerPostJson struct {
	Owner string `json:"owner"`
	TokenID int32 `json:"tokenId"`
	IpfsURI string `json:"ipfsURI"`
}

type MiniGamePostJson struct {
	Wallet string `json:"wallet"`
	Mark int32 `json:"mark"`
}

type MiniGameWalletJson struct {
	Wallet string `json:"wallet"`
}

type ResponseErrJson struct {
	Status string `json:"status"`
	ErrMsg string `json:"errMsg"`
}

type ResponseSuccessJson struct {
	Status string `json:"status"`
	Data interface{} `json:"data"`
}