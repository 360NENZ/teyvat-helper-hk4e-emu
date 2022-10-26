package store

import (
	"context"

	"github.com/uptrace/bun"
)

type PlayerDataStore struct{ db *bun.DB }

func (s *PlayerDataStore) CreatePlayerData(ctx context.Context, record *PlayerData) error {
	_, err := s.db.NewInsert().Model(record).Exec(ctx)
	return err
}
func (s *PlayerDataStore) DeletePlayerData(ctx context.Context, id int32) error {
	_, err := s.db.NewDelete().Model(&PlayerData{ID: id}).WherePK().Exec(ctx)
	return err
}
func (s *PlayerDataStore) UpdatePlayerData(ctx context.Context, record *PlayerData) error {
	_, err := s.db.NewUpdate().Model(record).WherePK().OmitZero().Exec(ctx)
	return err
}
func (s *PlayerDataStore) GetPlayerData(ctx context.Context, id int32) (*PlayerData, error) {
	record := new(PlayerData)
	if err := s.db.NewSelect().Model(record).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}
	return record, nil
}
func (s *PlayerDataStore) GetAllPlayerData(ctx context.Context, offset, limit int) ([]*PlayerData, error) {
	panic("todo")
}
