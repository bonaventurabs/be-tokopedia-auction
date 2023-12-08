package model

type ClassDetail struct {
	Status string `json:"status"`
}

type ClassResult struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   ClassDetail `json:"data"`
}

const PowerMerchant int = 1
const PowerMerchantPRO int = 2
