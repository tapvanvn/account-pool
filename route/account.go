package route

import (
	"github.com/tapvanvn/gorouter/v2"
)

func Account(context *gorouter.RouteContext) {

	if context.R.Method == "POST" || context.R.Method == "PUT" {

		if context.Action == "select" {

			return
		}
	}
}
