package game

import (
	"context"
	"database/sql"

	"github.com/teyvat-helper/hk4e-emu/pkg/sdk"
	"github.com/teyvat-helper/hk4e-emu/pkg/store"
)

func (s *PlayerSession) TokenLogin(ctx context.Context, id int64, token string) (*Player, error) {
	account, err := s.CheckComboToken(ctx, id, token)
	if err != nil {
		return nil, err
	}
	player, err := s.server.Store().Player().GetPlayerByAccountID(ctx, account.ID)
	if err != sql.ErrNoRows {
		player.Account = account
		return NewPlayer(player), err
	}
	player, err = s.CreatePlayer(ctx, account.ID)
	if err != nil {
		return nil, err
	}
	player.Account = account
	return NewPlayer(player), nil
}

func (s *PlayerSession) CheckComboToken(ctx context.Context, id int64, token string) (*store.Account, error) {
	record, err := s.server.Store().Account().GetAccount(ctx, id)
	if err != nil {
		return nil, err
	}
	if record.ComboToken == "" || record.ComboToken != token {
		return nil, sdk.ErrInvalidComboToken
	}
	return record, nil
}
