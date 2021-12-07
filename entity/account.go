package entity

import (
	"fmt"

	"github.com/tapvanvn/accountpool/common"
)

func AccountID(accountType common.AccountType, accountID string) string {
	return fmt.Sprintf("%s_%s", accountType, accountID)
}

func NewAccount(accountType common.AccountType, accountID string) *Account {
	return &Account{
		AccountType: accountType,
		AccountID:   accountID,
		Meta:        map[string]string{},
	}
}

type Account struct {
	AccountID   string             `json:"AccountID" bson:"AccountID"`
	AccountType common.AccountType `json:"AccountType" bson:"AccountType"`
	Meta        map[string]string  `json:"Meta" bson:"meta"`
}

func (doc *Account) GetID() string {

	return AccountID(doc.AccountType, doc.AccountID)
}
