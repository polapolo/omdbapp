package entity

import "time"

type (
	// Search -> represent 1 row of search table
	Search struct {
		ID        int       `db:"id"`
		Keyword   string    `db:"keyword"`
		Page      int       `db:"page"`
		CreatedAt time.Time `db:"created_at"`
	}
)
