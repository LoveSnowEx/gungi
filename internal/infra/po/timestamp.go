package po

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type TimeStamp struct {
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

func (u *TimeStamp) BeforeUpdate(ctx context.Context, query *bun.UpdateQuery) error {
	query.Set("updated_at = CURRENT_TIMESTAMP")
	return nil
}
