package runtime

import (
	"sync"

	"github.com/tapvanvn/accountpool/common"
	"github.com/tapvanvn/accountpool/inf"
)

var __generators map[common.AccountType]inf.IGenerator = make(map[common.AccountType]inf.IGenerator)
var __mux sync.Mutex

func RegGenerator(accountType common.AccountType, generator inf.IGenerator) error {
	__mux.Lock()
	__generators[accountType] = generator
	__mux.Unlock()
	return nil
}

func GetGenerator(accountType common.AccountType) inf.IGenerator {

	__mux.Lock()
	defer __mux.Unlock()

	if gener, has := __generators[accountType]; has {

		return gener
	}
	return nil
}
