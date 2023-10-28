package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Player       PlayerModel
	Club         ClubModel
	Player_Stats PlayerStatsModel
	Club_Stats   ClubStatsModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Player:       PlayerModel{DB: db},
		Club:         ClubModel{DB: db},
		Player_Stats: PlayerStatsModel{DB: db},
		Club_Stats:   ClubStatsModel{DB: db},
	}
}
