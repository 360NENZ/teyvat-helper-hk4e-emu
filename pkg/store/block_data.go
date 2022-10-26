package store

import (
	"context"

	"github.com/uptrace/bun"
)

type BlockDataStore struct{ db *bun.DB }

func (s *BlockDataStore) CreateBlockData(ctx context.Context, record *BlockData) error {
	_, err := s.db.NewInsert().Model(record).Exec(ctx)
	return err
}
func (s *BlockDataStore) DeleteBlockData(ctx context.Context, id int64) error {
	_, err := s.db.NewDelete().Model(&BlockData{ID: id}).WherePK().Exec(ctx)
	return err
}
func (s *BlockDataStore) UpdateBlockData(ctx context.Context, record *BlockData) error {
	_, err := s.db.NewUpdate().Model(record).WherePK().OmitZero().Exec(ctx)
	return err
}
func (s *BlockDataStore) GetBlockData(ctx context.Context, id int64) (*BlockData, error) {
	record := new(BlockData)
	if err := s.db.NewSelect().Model(record).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}
	return record, nil
}
func (s *BlockDataStore) GetBlockDataByPlayerID(ctx context.Context, id, blockID int32) (*BlockData, error) {
	record := new(BlockData)
	if err := s.db.NewSelect().Model(record).Where("player_id = ?", id).Where("block_id = ?", blockID).Scan(ctx); err != nil {
		return nil, err
	}
	return record, nil
}
func (s *BlockDataStore) GetAllBlockData(ctx context.Context, offset, limit int) ([]*BlockData, error) {
	panic("todo")
}
