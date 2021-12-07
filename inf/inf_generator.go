package inf

import (
	"github.com/tapvanvn/accountpool/common"
	"github.com/tapvanvn/accountpool/entity"
)

type IGenerator interface {
	GetAccountType() common.AccountType //which account type this generator will generate
	Generate() (*entity.Account, error) //generate new account
}
