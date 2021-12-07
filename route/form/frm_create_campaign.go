package form

import "github.com/tapvanvn/accountpool/common"

type FormCreateCampaign struct {
	CampaignName string                     `json:"CampaignName"`
	Accounts     map[common.AccountType]int `json:"Accounts"`
}
