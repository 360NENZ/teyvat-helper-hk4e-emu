package store

import (
	"context"

	"github.com/uptrace/bun"
)

type HomeDataStore struct{ db *bun.DB }

func (s *HomeDataStore) CreateHomeData(ctx context.Context, record *HomeData) error {
	_, err := s.db.NewInsert().Model(record).Exec(ctx)
	return err
}
func (s *HomeDataStore) DeleteHomeData(ctx context.Context, id int32) error {
	_, err := s.db.NewDelete().Model(&HomeData{ID: id}).WherePK().Exec(ctx)
	return err
}
func (s *HomeDataStore) UpdateHomeData(ctx context.Context, record *HomeData) error {
	_, err := s.db.NewUpdate().Model(record).WherePK().OmitZero().Exec(ctx)
	return err
}
func (s *HomeDataStore) GetHomeData(ctx context.Context, id int32) (*HomeData, error) {
	record := new(HomeData)
	if err := s.db.NewSelect().Model(record).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}
	return record, nil
}
func (s *HomeDataStore) GetAllHomeData(ctx context.Context, offset, limit int) ([]*HomeData, error) {
	panic("todo")
}
