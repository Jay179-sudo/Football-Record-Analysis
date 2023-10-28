package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type PlayerModel struct {
	DB *sql.DB
}

type Player struct {
	Player_ID   int64
	Player_Name string
	DOB         time.Time
	Position    string
}

func (p PlayerModel) Insert(player *Player) error {
	query := `
		INSERT INTO Player (Player_ID, Player_Name, DOB, Position)
		values ($1, $2, $3, $4)
		RETURNING Player_ID, Player_Name, DOB, Position 
	`

	args := []interface{}{player.Player_ID, player.Player_Name, player.DOB, player.Position}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	return p.DB.QueryRowContext(ctx, query, args...).Scan(&player.Player_ID, &player.Player_Name, &player.DOB, &player.Position)
}

func (p PlayerModel) Update(player *Player) error {
	query :=
		`
		UPDATE Player
		SET Player_Name = $2, DOB = $3, Position = $4
		WHERE Player_ID = $1 
		RETURNING Player_ID
	`
	args := []interface{}{
		player.Player_ID,
		player.Player_Name,
		player.DOB,
		player.Position,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := p.DB.QueryRowContext(ctx, query, args...).Scan(&player.Player_ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}

func (p PlayerModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query :=
		`
		DELETE FROM Player
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := p.DB.ExecContext(ctx, query, id)
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

func (p PlayerModel) Get(id int64) (*Player, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query :=
		`
		SELECT Player_ID, Player_Name, DOB, Position
		FROM Player
		WHERE Player_ID = $1
	`

	var player Player
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := p.DB.QueryRowContext(ctx, query, id).Scan(
		&player.Player_ID,
		&player.Player_Name,
		&player.DOB,
		&player.Position,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &player, nil
}
