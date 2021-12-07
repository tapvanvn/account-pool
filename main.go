package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/tapvanvn/accountpool/route"
	"github.com/tapvanvn/accountpool/runtime"
	"github.com/tapvanvn/accountpool/utility"
	"github.com/tapvanvn/gorouter/v2"
	"github.com/tapvanvn/goutil"
)

type Handles []gorouter.RouteHandle
type Endpoint gorouter.EndpointDefine

func main() {

	var port = goutil.MustGetEnv("PORT")

	rootPath, err := filepath.Abs(filepath.Dir(os.Args[0]))

	configPath := utility.GetGeneralConfigPath(rootPath)

	err = runtime.InitRuntime(configPath)

	if err != nil {

		log.Fatal("init engine fail", err)
	}

	jsonFile, err := os.Open(rootPath + "/route.jsonc")

	if err != nil {

		log.Fatal(err)
	}
	bytes, err := ioutil.ReadAll(jsonFile)
	bytes = goutil.TripJSONComment(bytes)

	var handers = map[string]gorouter.EndpointDefine{

		"":         {Handles: Handles{route.Root}},
		"account":  {Handles: Handles{route.Account}},
		"campaign": {Handles: Handles{route.Campaign}},
	}

	var router = gorouter.Router{}

	router.Init("api/", string(bytes), handers)

	http.Handle("/api/", router)

	cacheFileServer := goutil.NewCacheFileServer(http.Dir(rootPath + "/static"))

	fileServer := http.FileServer(cacheFileServer)

	http.Handle("/", fileServer)

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("i am ok"))
	})

	fmt.Println("listen on port", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
