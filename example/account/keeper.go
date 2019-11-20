package account

import (
	"context"
	"github.com/regen-network/chain-schema/example"
	"github.com/regen-network/chain-schema/store"
)

type Keeper interface {
	GetAccount(ctx context.Context,address example.Address) (Account, error)
	SetAccount(ctx context.Context, account Account) error
}

type keeper struct {
	accountTable store.NaturalKeyTable
}

func (k keeper) GetAccount(ctx context.Context, address example.Address) (Account, error) {
	var res Account
	_, err := k.accountTable.GetOne(ctx, address, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (k keeper) SetAccount(ctx context.Context, account Account) error {
	return k.accountTable.Save(ctx, account)
}

func (m Account) ID() []byte {
	return m.Address
}

func NewKeeper() Keeper {
	return &keeper{}
}

