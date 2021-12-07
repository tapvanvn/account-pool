package runtime

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/tapvanvn/accountpool/common"
	"github.com/tapvanvn/accountpool/entity"
	"github.com/tapvanvn/accountpool/generator"
	engines "github.com/tapvanvn/godbengine"
	"github.com/tapvanvn/godbengine/engine"
	"github.com/tapvanvn/godbengine/engine/adapter"
	"github.com/tapvanvn/goutil"
)

const (
	ENV_DEV = iota
	ENV_PROD
	ENV_LOCAL
)

var NodeName string = ""
var Environment = ENV_PROD

func IsDev() bool {
	return Environment == ENV_DEV
}

func IsProd() bool {
	return Environment == ENV_PROD
}

func IsLocal() bool {
	return Environment == ENV_LOCAL
}

var Config *entity.Config = nil

//Start start engine
func startEngine(eng *engine.Engine) {

	env := Config.Environment

	Environment = ENV_DEV

	if env == "prod" {

		Environment = ENV_PROD

	} else if env == "local" {

		Environment = ENV_LOCAL
	}

	var documentDB engine.DocumentPool = nil

	if Config.DocDB != nil {

		connectString := Config.DocDB.ConnectionString
		databaseName := Config.DocDB.DatabaseName

		if Config.DocDB.Provider == "mongodb" {
			mongoPool := &adapter.MongoPool{}
			err := mongoPool.InitWithDatabase(connectString, databaseName)
			if err != nil {
				log.Fatal("cannot init mongo")
			}
			documentDB = mongoPool

		} else {

			firestorePool := adapter.FirestorePool{}
			firestorePool.Init(connectString)
			documentDB = &firestorePool
		}
	}

	eng.Init(nil, documentDB, nil)
}

//InitRuntime seting up the runtime.
//TODO: maybe config is needed
func InitRuntime(configPath string) error {

	NodeName = goutil.GenVerifyCode(5)

	//MARK: init system config
	jsonFile, err := os.Open(configPath)

	if err != nil {

		return err
	}
	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {

		return err
	}
	bytes = goutil.TripJSONComment(bytes)

	systemConfig := &entity.Config{}

	if err != nil {

		return err
	}

	err = json.Unmarshal(bytes, systemConfig)

	if err != nil {
		fmt.Println(string(bytes))
		return err
	}

	Config = systemConfig

	engines.InitEngineFunc = startEngine
	eng := engines.GetEngine()

	RegGenerator(common.AccountTypeEth, generator.NewEthGenerator())

	err = initCampaignManager(eng)

	if err != nil {

		return err
	}
	return nil
}
