package repository

import (
	"fmt"

	"github.com/tapvanvn/accountpool/common"
	"github.com/tapvanvn/accountpool/entity"
	engines "github.com/tapvanvn/godbengine"
	"github.com/tapvanvn/godbengine/engine"
)

func PutAccount(account *entity.Account) error {

	eng := engines.GetEngine()

	documentPool := eng.GetDocumentPool()

	return documentPool.Put(CollectionAccount, account)
}

func GetAccount(accountType common.AccountType, accountID string) (*entity.Account, error) {

	eng := engines.GetEngine()

	documentPool := eng.GetDocumentPool()

	doc := &entity.Account{}

	err := documentPool.Get(CollectionAccount, entity.AccountID(accountType, accountID), doc)
	if err != nil {

		return nil, err
	}
	return doc, nil
}

func GetAccounts(accountType common.AccountType, accountIDs []string) ([]*entity.Account, error) {

	eng := engines.GetEngine()

	documentPool := eng.GetDocumentPool()

	query := engine.MakeDBQuery(CollectionAccount, false)

	query.Filter("AccountType", "=", accountType)
	query.Filter("AccountID", "in", accountIDs)

	queryResult := documentPool.Query(query)

	if queryResult.Error() != nil {

		return nil, queryResult.Error()
	}

	defer queryResult.Close()

	accounts := make([]*entity.Account, 0)

	for {
		item := &entity.Account{}

		err := queryResult.Next(&item)

		if err == nil {

			accounts = append(accounts, item)

		} else if documentPool.IsNoRecordError(err) {

			break

		} else {

			fmt.Println((err.Error()))
			return nil, err
		}
	}
	return accounts, nil
}
