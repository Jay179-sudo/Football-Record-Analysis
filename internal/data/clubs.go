package data

import (
	"context"
	"database/sql"
	"time"
)

type ClubModel struct {
	DB *sql.DB
}

type Club struct {
	Club_ID   int64
	Team_Name string
	Country   string
}

func (c ClubModel) Insert(club *Club) error {
	query :=
		`
		INSERT INTO Club
		VALUES ($1, $2, $3)
		RETURNING Club_ID, Team_Name, Country
	`

	args := []interface{}{club.Club_ID, club.Team_Name, club.Country}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := c.DB.QueryRowContext(ctx, query, args...).Scan(&club.Club_ID, &club.Team_Name, &club.Country)
	if err != nil {
		return err
	}
	return nil
}

func (c ClubModel) Update(club *Club) error {
	query :=
		`
		UPDATE Club
		SET Team_Name = $2, Country = $3
		WHERE Club_ID = $1
		RETURNING Club_ID
	`

	args := []interface{}{
		club.Club_ID,
		club.Country,
		club.Team_Name,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := c.DB.QueryRowContext(ctx, query, args...).Scan(&club.Club_ID)
	if err != nil {
		return err
	}

	return nil
}

func (c ClubModel) Delete(Club_ID int64) error {
	if Club_ID < 1 {
		return ErrRecordNotFound
	}

	query :=
		`
		DELETE FROM Club
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
