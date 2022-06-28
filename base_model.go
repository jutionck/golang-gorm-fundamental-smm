package main

import (
	"database/sql"
	"time"
)

type BaseModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
