package store

import (
	"context"

	"github.com/uptrace/bun"
)

type AccountStore struct{ db *bun.DB }

func (s *AccountStore) CreateAccount(ctx context.Context, record *Account) error {
	_, err := s.db.NewInsert().Model(record).Exec(ctx)
	return err
}
func (s *AccountStore) DeleteAccount(ctx context.Context, id int64) error {
	_, err := s.db.NewDelete().Model(&Account{ID: id}).WherePK().Exec(ctx)
	return err
}
func (s *AccountStore) UpdateAccount(ctx context.Context, record *Account) error {
	_, err := s.db.NewUpdate().Model(record).WherePK().OmitZero().Exec(ctx)
	return err
}
func (s *AccountStore) UpdateAccountPassword(ctx context.Context, id int64, password string) error {
	_, err := s.db.NewUpdate().Model(&Account{ID: id, Password: password}).WherePK().OmitZero().Exec(ctx)
	return err
}
func (s *AccountStore) UpdateAccountLoginToken(ctx context.Context, id int64, token string) error {
	_, err := s.db.NewUpdate().Model(&Account{ID: id, LoginToken: token}).WherePK().OmitZero().Exec(ctx)
	return err
}
func (s *AccountStore) UpdateAccountComboToken(ctx context.Context, id int64, token string) error {
	_, err := s.db.NewUpdate().Model(&Account{ID: id, ComboToken: token}).WherePK().OmitZero().Exec(ctx)
	return err
}
func (s *AccountStore) GetAccount(ctx context.Context, id int64) (*Account, error) {
	record := new(Account)
	if err := s.db.NewSelect().Model(record).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}
	return record, nil
}
func (s *AccountStore) GetAccountByEmail(ctx context.Context, email string) (*Account, error) {
	record := new(Account)
	if err := s.db.NewSelect().Model(record).Where("email = ?", email).Scan(ctx); err != nil {
		return nil, err
	}
	return record, nil
}
func (s *AccountStore) GetAccountByUsername(ctx context.Context, username string) (*Account, error) {
	record := new(Account)
	if err := s.db.NewSelect().Model(record).Where("username = ?", username).Scan(ctx); err != nil {
		return nil, err
	}
	return record, nil
}
func (s *AccountStore) GetAccounts(ctx context.Context, offset, limit int) ([]*Account, error) {
	panic("todo")
}
