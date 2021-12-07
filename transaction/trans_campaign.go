package transaction

import (
	"github.com/tapvanvn/accountpool/common"
	"github.com/tapvanvn/accountpool/entity"
	"github.com/tapvanvn/accountpool/repository"
)

func CreateCampaign(campaignName string) (*entity.Campaign, error) {

	camp, err := repository.GetCampaign(campaignName)

	if err == nil || !repository.IsNoRecordError(err) {

		return nil, common.ErrAPCampaignExisted
	}
	camp = entity.NewCampaign(campaignName, "")

	err = repository.PutCampaign(camp)

	if err != nil {

		return nil, err
	}
	return camp, nil
}
