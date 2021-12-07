package route

import (
	"github.com/tapvanvn/accountpool/common"
	"github.com/tapvanvn/accountpool/repository"
	"github.com/tapvanvn/accountpool/route/response"
	"github.com/tapvanvn/gorouter/v2"
)

func Account(context *gorouter.RouteContext) {

	accountType, hasType := context.AnyIndex("account_type")
	accountID, hasID := context.AnyIndex("account_id")

	if !hasType || !hasID {

		response.BadRequest(context, 0, common.ErrAPInvalidForm.Error(), nil)
		return
	}

	if context.R.Method == "GET" {

		if context.Action == "index" {

			getAccount(common.AccountType(accountType), accountID, context)
		}
	}
}

func getAccount(accountType common.AccountType, accountID string, context *gorouter.RouteContext) {

	acc, err := repository.GetAccount(common.AccountType(accountType), accountID)
	if err != nil {

		response.BadRequest(context, 0, common.ErrAPInvalidForm.Error(), nil)
		return
	}
	response.Success(context, acc)
}
