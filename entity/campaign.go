package entity

import (
	"time"

	"github.com/tapvanvn/accountpool/common"
)

func NewCampaign(name string, title string) *Campaign {
	return &Campaign{
		CampaignName: name,
		Title:        title,
		CreateTime:   time.Now().Unix(),
		Accounts:     make(map[common.AccountType][]string),
		Done:         make(map[common.Action][]string),
	}
}

type Campaign struct {
	CampaignName string                          `json:"CampaignName" bson:"CampaignName"`
	Title        string                          `json:"Title" bson:"Title"`
	CreateTime   int64                           `json:"CreateTime" bson:"CreateTime"`
	Accounts     map[common.AccountType][]string `json:"Accounts" bson:"Accounts"`
	Done         map[common.Action][]string      `json:"Done" bson:"Done"` //accounts that done the action.
}

func (doc *Campaign) GetAccountOfType(accountType common.AccountType) []string {

	if accounts, has := doc.Accounts[accountType]; has {

		return accounts
	}
	return []string{}
}
func (doc *Campaign) SelectAccount(accountType common.AccountType, accountIDs []string) {

	doc.Accounts[accountType] = accountIDs
}

func (doc *Campaign) GetID() string {

	return doc.CampaignName
}
