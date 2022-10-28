package store

import (
	"context"
	"database/sql"

	"github.com/teyvat-helper/hk4e-emu/pkg/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

type Store struct {
	config     *config.Config
	db         *bun.DB
	account    *AccountStore
	player     *PlayerStore
	playerData *PlayerDataStore
	blockData  *BlockDataStore
	homeData   *HomeDataStore
}

func NewStore(config *config.Config) *Store {
	s := &Store{config: config}
	s.init()
	return s
}

func (s *Store) init() {
	sqlite, err := sql.Open(sqliteshim.ShimName, s.config.Database.DSN)
	if err != nil {
		panic(err)
	}
	sqlite.SetMaxOpenConns(1)
	s.db = bun.NewDB(sqlite, sqlitedialect.New())
	s.account = &AccountStore{db: s.db}
	s.player = &PlayerStore{db: s.db}
	s.playerData = &PlayerDataStore{db: s.db}
	s.blockData = &BlockDataStore{db: s.db}
	if err := s.install(context.Background()); err != nil {
		panic(err)
	}
}

func (s *Store) Account() *AccountStore       { return s.account }
func (s *Store) Player() *PlayerStore         { return s.player }
func (s *Store) PlayerData() *PlayerDataStore { return s.playerData }
func (s *Store) BlockData() *BlockDataStore   { return s.blockData }
func (s *Store) HomeData() *HomeDataStore     { return s.homeData }
