package po

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	TimeStamp
	ID   uint   `bun:",pk,autoincrement"`
	Name string `bun:",nullzero,notnull"`
}
