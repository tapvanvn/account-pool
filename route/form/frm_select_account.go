package form

import "github.com/tapvanvn/accountpool/common"

type FormSelectAccount struct {
	AccountType common.AccountType `json:"AccountType"`
	Quantity    int                `json:"Quantity"`
}
