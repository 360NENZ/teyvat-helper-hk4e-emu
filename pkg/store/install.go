package store

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *Store) checkInit(ctx context.Context) bool {
	s.db.NewCreateTable().Model((*Config)(nil)).IfNotExists().Exec(ctx)
	config := new(Config)
	err := s.db.NewSelect().Model(config).Where("key = ?", "init").Scan(ctx)
	if err != sql.ErrNoRows && err != nil {
		panic(err)
	}
	if err == sql.ErrNoRows || config.Value != "true" {
		if _, err = s.db.NewInsert().Model(&Config{
			Key:   "init",
			Value: "true",
		}).On("CONFLICT (key) DO UPDATE").Exec(ctx); err != nil {
			panic(err)
		}
		return false
	}
	return true
}

func (s *Store) install(ctx context.Context) error {
	if s.checkInit(ctx) {
		return nil
	}
	s.db.NewDropTable().Model((*HomeData)(nil)).IfExists().Exec(ctx)
	s.db.NewDropTable().Model((*BlockData)(nil)).IfExists().Exec(ctx)
	s.db.NewDropTable().Model((*PlayerData)(nil)).IfExists().Exec(ctx)
	s.db.NewDropTable().Model((*Player)(nil)).IfExists().Exec(ctx)
	s.db.NewDropTable().Model((*Account)(nil)).IfExists().Exec(ctx)
	s.db.NewCreateTable().Model((*Account)(nil)).IfNotExists().Exec(ctx)
	s.db.NewCreateTable().Model((*Player)(nil)).IfNotExists().Exec(ctx)
	s.db.NewCreateTable().Model((*PlayerData)(nil)).IfNotExists().Exec(ctx)
	s.db.NewCreateTable().Model((*BlockData)(nil)).IfNotExists().Exec(ctx)
	s.db.NewCreateTable().Model((*HomeData)(nil)).IfNotExists().Exec(ctx)
	randBytes := make([]byte, 6)
	if _, err := rand.Read(randBytes); err != nil {
		return err
	}
	base64Str := base64.RawStdEncoding.EncodeToString(randBytes)
	hashed, err := bcrypt.GenerateFromPassword([]byte(base64Str), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	log.Info().Str("username", "admin").Str("password", base64Str).Msg("admin account created")
	return s.Account().CreateAccount(ctx, &Account{
		Email:    "admin@localhost",
		Username: "admin",
		Password: string(hashed),
	})
}
