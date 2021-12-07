package transaction

import (
	"github.com/tapvanvn/accountpool/common"
	"github.com/tapvanvn/accountpool/entity"
	"github.com/tapvanvn/accountpool/repository"
	"github.com/tapvanvn/accountpool/runtime"
	"github.com/tapvanvn/godbengine/engine"
)

//SelectAccount - choice a number (quantity) of account type (create if needed) then add to the campain
func SelectAccount(accountType common.AccountType, quantity int, campaignName string) ([]*entity.Account, error) {

	camp, err := repository.GetCampaign(campaignName)

	if err != nil {
		if repository.IsNoRecordError(err) {
			return nil, common.ErrAPCampaignNotExisted
		}
		return nil, err
	}

	accounts, err := repository.GetAccounts(accountType, camp.GetAccountOfType(accountType))
	if err != nil {
		return nil, err
	}

	if len(accounts) < quantity {

		gener := runtime.GetGenerator(accountType)

		if gener == nil {

			return nil, common.ErrAPGeneratorNotExisted
		}
		truck := engine.CreateDataStruck()

		for i := len(accounts); i < quantity; i++ {

			acc, err := gener.Generate()

			if err != nil {

				return nil, err
			}
			if len(truck.Items) == 150 {

				repository.SaveDataTruck(&truck)
				truck = engine.CreateDataStruck()
			}

			accounts = append(accounts, acc)

			truck.Append(repository.CollectionAccount, acc)
		}

		if len(truck.Items) > 0 {

			repository.SaveDataTruck(&truck)
		}
	}
	accountIDs := []string{}
	for _, acc := range accounts {
		accountIDs = append(accountIDs, acc.AccountID)
	}
	camp.SelectAccount(accountType, accountIDs)
	err = repository.PutCampaign(camp)

	if err != nil {

		return nil, err
	}

	return accounts, nil
}
