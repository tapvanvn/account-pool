package repository

import (
	"github.com/tapvanvn/accountpool/entity"
	engines "github.com/tapvanvn/godbengine"
)

func PutCampaign(campaign *entity.Campaign) error {

	eng := engines.GetEngine()

	documentPool := eng.GetDocumentPool()

	return documentPool.Put(CollectionCampaign, campaign)
}

func GetCampaign(campaignName string) (*entity.Campaign, error) {

	eng := engines.GetEngine()

	documentPool := eng.GetDocumentPool()

	doc := &entity.Campaign{}
	err := documentPool.Get(CollectionCampaign, campaignName, doc)
	if err != nil {

		return nil, err
	}
	return doc, nil
}
