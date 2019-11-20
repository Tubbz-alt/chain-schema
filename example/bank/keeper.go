package bank

import (
	"context"
	"github.com/regen-network/chain-schema/example"
	"github.com/regen-network/chain-schema/example/account"
	"github.com/regen-network/chain-schema/store"
)

type keeper struct {
	accKeeper account.Keeper
}

type Keeper interface {
	BankServer
}

func NewKeeper(builder store.ModuleBuilder, accKeeper account.Keeper) Keeper {
	res := &keeper{accKeeper: accKeeper}
	builder.RegisterServer(&_Bank_serviceDesc, res)
	return res
}

func (k keeper) Send(context.Context, *SendRequest) (*example.Empty, error) {
	panic("implement me")
}

func (k keeper) Burn(context.Context, *BurnRequest) (*example.Empty, error) {
	panic("implement me")
}

func (k keeper) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	panic("implement me")
}

func (k keeper) GetSupply(context.Context, *GetSupplyRequest) (*GetSupplyResponse, error) {
	panic("implement me")
}
