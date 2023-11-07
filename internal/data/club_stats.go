package data

import (
	"context"
	"database/sql"
	"time"
)

type ClubStatsModel struct {
	DB *sql.DB
}

type Club_Stats struct {
	Club_ID int64
	Season  int64
	Wins    int32
	Losses  int32
	Draws   int32
}

func (c ClubStatsModel) Insert(Club_Stats Club_Stats) error {
	query :=
		`
		INSERT INTO Club_Stats
		VALUES ($1, $2, $3, $4, $5)
		RETURNING Club_ID, Season, Wins, Losses, Draws
	`

	args := []interface{}{Club_Stats.Club_ID, Club_Stats.Season, Club_Stats.Wins, Club_Stats.Losses, Club_Stats.Draws}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := c.DB.QueryRowContext(ctx, query, args...).Scan(&Club_Stats.Club_ID, &Club_Stats.Season, &Club_Stats.Wins, &Club_Stats.Losses, &Club_Stats.Draws)
	if err != nil {
		return err
	}

	return nil
}

func (c ClubStatsModel) Update(Club_Stats *Club_Stats) error {
	query :=

		`
		UPDATE Club_Stats
		SET Average_Age = $3, Squad_Size = $4
		WHERE Club_ID = $1 AND Season = $2
		RETURNING Club_ID
	`

	args := []interface{}{Club_Stats.Club_ID, Club_Stats.Season, Club_Stats.Wins, Club_Stats.Losses, Club_Stats.Draws}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := c.DB.QueryRowContext(ctx, query, args...).Scan(&Club_Stats.Club_ID)
	if err != nil {
		return err
	}

	return nil
}

func (c ClubStatsModel) Delete(Club_ID int64) error {
	query :=

		`
		DELETE FROM Club_Stats
		WHERE Club_ID = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := c.DB.ExecContext(ctx, query, Club_ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}
