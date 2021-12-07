package generator

import (
	"crypto/ecdsa"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tapvanvn/accountpool/common"
	"github.com/tapvanvn/accountpool/entity"
)

func NewEthGenerator() *EthGenerator {
	return &EthGenerator{}
}

type EthGenerator struct {
}

func (gener *EthGenerator) GetAccountType() common.AccountType {

	return common.AccountTypeEth
}
func (gener *EthGenerator) Generate() (*entity.Account, error) {

	privateKey, err := crypto.GenerateKey()

	if err != nil {

		return nil, err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)

	privateKeyString := fmt.Sprintf("%s", hexutil.Encode(privateKeyBytes))

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	account := entity.NewAccount(common.AccountTypeEth, address)

	account.Meta["PrivateKey"] = privateKeyString

	return account, nil
}
