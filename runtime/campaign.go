package runtime

import (
	"sync"
	"time"

	"github.com/tapvanvn/accountpool/entity"
	"github.com/tapvanvn/accountpool/repository"
	"github.com/tapvanvn/godbengine/engine"
)

var __campaigns map[string]*entity.Campaign = make(map[string]*entity.Campaign)
var __camp_mux sync.Mutex
var __watcher *engine.Watcher = nil

func initCampaignManager(eng *engine.Engine) error {

	__watcher = engine.NewWatcher(10*time.Second, eng.GetDocumentPool())
	__watcher.Run()
	return nil
}

//Get the campaign
func GetCampaign(name string) (*entity.Campaign, error) {
	__camp_mux.Lock()
	if camp, ok := __campaigns[name]; ok {
		__camp_mux.Unlock()
		return camp, nil
	}
	__camp_mux.Unlock()

	camp, err := repository.GetCampaign(name)
	if err != nil {
		return nil, err
	}
	__camp_mux.Lock()
	__campaigns[name] = camp
	__watcher.Watch(repository.CollectionCampaign, camp)
	__camp_mux.Unlock()
	return camp, nil
}
