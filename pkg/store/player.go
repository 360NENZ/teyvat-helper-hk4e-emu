package store

import (
	"context"

	"github.com/uptrace/bun"
)

type PlayerStore struct{ db *bun.DB }

func (s *PlayerStore) CreatePlayer(ctx context.Context, record *Player) error {
	_, err := s.db.NewInsert().Model(record).Exec(ctx)
	return err
}
func (s *PlayerStore) DeletePlayer(ctx context.Context, id int32) error {
	_, err := s.db.NewDelete().Model(&Player{ID: id}).WherePK().Exec(ctx)
	return err
}
func (s *PlayerStore) UpdatePlayer(ctx context.Context, record *Player) error {
	_, err := s.db.NewUpdate().Model(record).WherePK().OmitZero().Exec(ctx)
	return err
}
func (s *PlayerStore) GetPlayer(ctx context.Context, id int32) (*Player, error) {
	record := new(Player)
	if err := s.db.NewSelect().Model(record).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}
	return record, nil
}
func (s *PlayerStore) GetPlayerByAccountID(ctx context.Context, id int64) (*Player, error) {
	record := new(Player)
	if err := s.db.NewSelect().Model(record).Where("account_id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}
	return record, nil
}
func (s *PlayerStore) GetPlayers(ctx context.Context, offset, limit int) ([]*Player, error) {
	panic("todo")
}
