package store

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type Timestamp struct {
	CreatedAt time.Time  `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time  `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt *time.Time `bun:",soft_delete,nullzero"`
}

var _ bun.BeforeAppendModelHook = (*Timestamp)(nil)

func (x *Timestamp) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		x.UpdatedAt = time.Now()
		x.CreatedAt = x.UpdatedAt
	case *bun.UpdateQuery:
		x.UpdatedAt = time.Now()
	}
	return nil
}

type Config struct {
	bun.BaseModel `bun:"table:hk4e_config"`

	Key   string `bun:"type:varchar(255),pk"`
	Value string `bun:"type:text,nullzero"`

	Timestamp
}

type Account struct {
	bun.BaseModel `bun:"table:hk4e_accounts"`

	ID       int64  `bun:",pk,autoincrement"`
	Email    string `bun:",notnull,unique"`
	Username string `bun:",notnull,unique"`
	Password string `bun:",nullzero"`

	LoginToken string `bun:",nullzero"`
	ComboToken string `bun:",nullzero"`

	Timestamp
}

type Player struct {
	bun.BaseModel `bun:"table:hk4e_players"`

	ID        int32    `bun:",pk,autoincrement"`
	AccountID int64    `bun:",notnull"`
	Account   *Account `bun:"rel:belongs-to,join:account_id=id"`

	Timestamp
}

type BlockData struct {
	bun.BaseModel `bun:"table:hk4e_block_data"`

	ID       int64  `bun:",pk,autoincrement"`
	PlayerID int32  `bun:",notnull"`
	BlockID  int32  `bun:",notnull"`
	BinData  []byte `bun:"type:mediumblob,nullzero"`
	Version  int32  `bun:",notnull"`

	Timestamp
}

type HomeData struct {
	bun.BaseModel `bun:"table:hk4e_home_data"`

	ID      int32  `bun:",pk,autoincrement"`
	BinData []byte `bun:"type:mediumblob,nullzero"`
	Version int32  `bun:",notnull"`

	Timestamp
}

type PlayerData struct {
	bun.BaseModel `bun:"table:hk4e_player_data"`

	ID      int32  `bun:",pk,autoincrement"`
	BinData []byte `bun:"type:mediumblob,nullzero"`
	Version int32  `bun:",notnull"`

	Timestamp
}
