package main

import (
	"time"

	"github.com/Jay179-sudo/FootballRecordAnalysis/internal/data"
)

// For the Player Model

// input a list of players, the generator will generate a player one by one in the channel
func (app *application) generator(done <-chan interface{}, players ...*data.Player) <-chan *data.Player {
	dataStream := make(chan *data.Player)
	go func() {
		defer close(dataStream)
		for _, i := range players {
			select {
			case <-done:
				return
			case dataStream <- i:
			}
		}
	}()
	return dataStream
}

func cleaned(player data.Player) data.Player {
	cleanedPlayer := data.Player{
		Player_ID:   1,
		Player_Name: "Jay",
		Position:    "Midfield",
		DOB:         time.Now(),
	}

	return cleanedPlayer
}

func transformed(player data.Player) data.Player {
	transformedPlayer := data.Player{
		Player_ID:   1,
		Player_Name: "Jay",
		Position:    "Midfield",
		DOB:         time.Now(),
	}

	return transformedPlayer
}
func (app *application) cleaning(done <-chan interface{}, player data.Player) <-chan data.Player {
	dataStream := make(chan data.Player)
	go func() {
		defer close(dataStream)
		select {
		case <-done:
			return
		case dataStream <- cleaned(player):
		}
	}()

	return dataStream
}

func (app *application) transformed(done <-chan interface{}, player data.Player) <-chan data.Player {
	dataStream := make(chan data.Player)
	go func() {
		defer close(dataStream)
		select {
		case <-done:
			return
		case dataStream <- transformed(player):
		}
	}()
	return dataStream
}
