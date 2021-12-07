package repository

import (
	engines "github.com/tapvanvn/godbengine"
	"github.com/tapvanvn/godbengine/engine"
)

var CollectionAccount = "account"
var CollectionCampaign = "campaign"

//IsNoRecordError check if a err is no record error
func IsNoRecordError(err error) bool {

	eng := engines.GetEngine()

	documentPool := eng.GetDocumentPool()

	return documentPool.IsNoRecordError(err)
}

//SaveDataTruck save data truck to database
func SaveDataTruck(dataTruck *engine.DataTruck) error {

	engine := engines.GetEngine()

	documentPool := engine.GetDocumentPool()

	transaction := documentPool.MakeTransaction()

	transaction.Begin()

	for _, item := range dataTruck.Items {

		transaction.Put(item.Collection, item.Document)
	}

	return transaction.Commit()
}
