package route

import (
	"github.com/tapvanvn/accountpool/common"
	"github.com/tapvanvn/accountpool/route/form"
	"github.com/tapvanvn/accountpool/route/response"
	"github.com/tapvanvn/accountpool/transaction"
	"github.com/tapvanvn/gorouter/v2"
	"github.com/tapvanvn/goutil"
)

func Root(context *gorouter.RouteContext) {

	if context.R.Method == "POST" || context.R.Method == "PUT" {

		if context.Action == "create_campaign" {

			createCampaign(context)
			return
		}
	}
}

func createCampaign(context *gorouter.RouteContext) {

	frm := &form.FormCreateCampaign{}

	if err := goutil.FromRequest(frm, context.R); err != nil {

		response.BadRequest(context, 0, common.ErrAPInvalidForm.Error(), nil)
		return
	}
	camp, err := transaction.CreateCampaign(frm.CampaignName)
	if err != nil {
		response.BadRequest(context, 0, err.Error(), nil)
		return
	}
	response.Success(context, camp)
}
