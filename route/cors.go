package route

import (
	"net/http"

	"github.com/tapvanvn/gorouter/v2"
)

func Cors(context *gorouter.RouteContext) {

	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, X-CSRF-Token, Access-Token, Account-Id"

	if origin := context.R.Header.Get("Origin"); origin != "" {
		context.W.Header().Set("Access-Control-Allow-Origin", "*")
		context.W.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		context.W.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
		context.W.Header().Set("Access-Control-Expose-Headers", "*")
	}
	if context.R.Method == "OPTIONS" {
		context.W.WriteHeader(http.StatusOK)
		context.Handled = true
	}

}
