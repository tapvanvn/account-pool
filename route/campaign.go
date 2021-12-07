package route

import (
	"github.com/tapvanvn/accountpool/common"
	"github.com/tapvanvn/accountpool/route/form"
	"github.com/tapvanvn/accountpool/route/response"
	"github.com/tapvanvn/accountpool/transaction"
	"github.com/tapvanvn/gorouter/v2"
	"github.com/tapvanvn/goutil"
)

func Campaign(context *gorouter.RouteContext) {

	if context.R.Method == "POST" || context.R.Method == "PUT" {

		campaignName, hasName := context.AnyIndex("campaign_name")

		if !hasName {
			response.BadRequest(context, 0, common.ErrAPInvalidForm.Error(), nil)
			return
		}
		if context.Action == "init_account" {

			selectAccount(campaignName, context)
			return
		}
	}
}

func selectAccount(campaignName string, context *gorouter.RouteContext) {

	frm := &form.FormSelectAccount{}

	if err := goutil.FromRequest(frm, context.R); err != nil {

		response.BadRequest(context, 0, common.ErrAPInvalidForm.Error(), nil)
		return
	}
	accounts, err := transaction.SelectAccount(frm.AccountType, frm.Quantity, campaignName)
	if err != nil {
		response.BadRequest(context, 0, err.Error(), nil)
		return
	}
	response.Success(context, accounts)
}
