package route

import (
	"github.com/tapvanvn/accountpool/common"
	"github.com/tapvanvn/accountpool/repository"
	"github.com/tapvanvn/accountpool/route/form"
	"github.com/tapvanvn/accountpool/route/response"
	"github.com/tapvanvn/accountpool/transaction"
	"github.com/tapvanvn/gorouter/v2"
	"github.com/tapvanvn/goutil"
)

func Campaign(context *gorouter.RouteContext) {

	campaignName, hasName := context.AnyIndex("campaign_name")

	if !hasName {
		response.BadRequest(context, 0, common.ErrAPInvalidForm.Error(), nil)
		return
	}
	if context.R.Method == "POST" || context.R.Method == "PUT" {

		if context.Action == "init_account" {

			initAccount(campaignName, context)
			return

		} else if context.Action == "select" {

			selectAccount(campaignName, context)
			return
		}
	}
}

func initAccount(campaignName string, context *gorouter.RouteContext) {

	frm := &form.FormInitAccount{}

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

func selectAccount(campaignName string, context *gorouter.RouteContext) {

	frm := &form.FormSelectAccount{}

	if err := goutil.FromRequest(frm, context.R); err != nil {

		response.BadRequest(context, 0, common.ErrAPInvalidForm.Error(), nil)
		return
	}
	camp, err := repository.GetCampaign(campaignName)
	if err != nil {
		response.BadRequest(context, 0, err.Error(), nil)
		return
	}

	response.Success(context, camp.Accounts[frm.AccountType])
}
