package data

import (
	"context"
	"database/sql"
	"time"
)

type PlayerStatsModel struct {
	DB *sql.DB
}

type Player_Stats struct {
	Player_ID         int64
	Current_Club_ID   int64
	Season            int32
	Yellow_Cards      int32
	Red_Cards         int32
	Goals             int32
	Assists           int32
	Minutes_Played    int32
	Player_Valuations int64
}

func (p PlayerStatsModel) Insert(Player *Player_Stats) error {
	query :=
		`
			INSERT INTO Player_Stats
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			RETURNING Player_ID, Current_Club_ID, Season, Yellow_Cards, Red_Cards, Goals, Assists, Minutes_Played, Player_Valuations
		`
	args := []interface{}{Player.Player_ID, Player.Current_Club_ID, Player.Season, Player.Yellow_Cards, Player.Red_Cards, Player.Goals, Player.Assists, Player.Minutes_Played, Player.Player_Valuations}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return p.DB.QueryRowContext(ctx, query, args...).Scan(&Player.Player_ID, &Player.Current_Club_ID, &Player.Season, &Player.Yellow_Cards, &Player.Red_Cards, &Player.Goals, &Player.Assists, &Player.Minutes_Played, &Player.Player_Valuations)
}

func (p PlayerStatsModel) Update(Player Player_Stats) error {
	query :=
		`
		UPDATE Player_Stats
		SET Yellow_Cards = $4, Red_Cards = $5, Goals = $6, Assists = $7, Minutes_Played = $8, Player_Valuations = $9
		WHERE Player_ID = $1 AND Current_Club_ID = $2 AND Season = $3 
		RETURNING Player_ID
	`

	args := []interface{}{
		Player.Player_ID,
		Player.Current_Club_ID,
		Player.Season,
		Player.Yellow_Cards,
		Player.Red_Cards,
		Player.Goals,
		Player.Assists,
		Player.Minutes_Played,
		Player.Player_Valuations,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := p.DB.QueryRowContext(ctx, query, args...).Scan(&Player.Player_ID)
	if err != nil {
		return err
	}
	return nil
}

func (p PlayerStatsModel) Delete(Player_ID int64, Current_Club_ID int64, Season int32) error {
	query :=
		`
		DELETE FROM Player_Stats
		WHERE Player_ID = $1 AND Current_Club_ID = $2 AND Season = $3
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := p.DB.ExecContext(ctx, query, Player_ID, Current_Club_ID, Season)
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}

type Minutes struct {
	Player_ID       int64
	Current_Club_ID int64
	Season          int32
	Minutes_Played  int64
}

func (p PlayerStatsModel) GetMinutes() ([]Minutes, error) {
	query :=
		`
		SELECT player_id, current_club_id, season, minutes_played
		FROM player_stats
		ORDER BY minutes_played;
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := p.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	result := []Minutes{}
	temp := Minutes{}
	for rows.Next() {
		err := rows.Scan(&temp.Player_ID, &temp.Current_Club_ID, &temp.Season, &temp.Minutes_Played)
		if err != nil {
			return nil, err
		}
		result = append(result, temp)
	}

	return result, nil

}

func (p PlayerStatsModel) GetAllStats() ([]Player_Stats, error) {
	query := `
		SELECT * FROM player_stats;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := p.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	stats := []Player_Stats{}
	for rows.Next() {
		var stat Player_Stats
		err := rows.Scan(
			&stat.Player_ID,
			&stat.Current_Club_ID,
			&stat.Season,
			&stat.Yellow_Cards,
			&stat.Red_Cards,
			&stat.Goals,
			&stat.Assists,
			&stat.Minutes_Played,
			&stat.Player_Valuations,
		)
		if err != nil {
			return nil, err
		}

		stats = append(stats, stat)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return stats, err
}
