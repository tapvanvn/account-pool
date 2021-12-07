package form

import "github.com/tapvanvn/accountpool/common"

type FormInitAccount struct {
	AccountType common.AccountType `json:"AccountType"`
	Quantity    int                `json:"Quantity"`
}
