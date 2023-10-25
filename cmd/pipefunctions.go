package main

import (
	"github.com/Jay179-sudo/FootballRecordAnalysis/internal/data"
)

// input a list of players, the generator will generate a player one by one in the channel
func (app *application) generator(done <-chan interface{}, players ...*data.Player_Stats) <-chan *data.Player_Stats {
	dataStream := make(chan *data.Player_Stats)
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

// Final function that updates the SQL database

func (app *application) cleaning(done <-chan interface{}, upStream <-chan data.Player_Stats, precomp Precomputations) <-chan data.Player_Stats {
	dataStream := make(chan data.Player_Stats)
	go func() {
		defer close(dataStream)
		select {
		case <-done:
			return
		case dataStream <- cleaned(precomp, <-upStream):
		}
	}()

	return dataStream
}

func (app *application) transformed(done <-chan interface{}, upStream <-chan data.Player_Stats, precomp Precomputations) <-chan data.Player_Stats {
	dataStream := make(chan data.Player_Stats)
	go func() {
		defer close(dataStream)
		select {
		case <-done:
			return
		case dataStream <- transformed(precomp, <-upStream):
		}
	}()
	return dataStream
}

func cleaned(precomp Precomputations, player data.Player_Stats) data.Player_Stats {
	if player.Minutes_Played == -1 {
		return player
	}
	cleanedPlayer := data.Player_Stats{
		Player_ID:         player.Player_ID,
		Current_Club_ID:   player.Current_Club_ID,
		Season:            player.Season,
		Yellow_Cards:      player.Yellow_Cards,
		Red_Cards:         player.Red_Cards,
		Goals:             player.Goals,
		Assists:           player.Assists,
		Minutes_Played:    player.Minutes_Played,
		Player_Valuations: player.Player_Valuations,
	}
	if cleanedPlayer.Minutes_Played <= int32(precomp.MinutesLower) {
		cleanedPlayer.Minutes_Played = -1
	}
	return cleanedPlayer
}

func transformed(precomp Precomputations, player data.Player_Stats) data.Player_Stats {
	if player.Minutes_Played == -1 {
		return player
	}
	transformedPlayer := data.Player_Stats{
		Player_ID:         player.Player_ID,
		Current_Club_ID:   player.Current_Club_ID,
		Season:            player.Season,
		Yellow_Cards:      player.Yellow_Cards,
		Red_Cards:         player.Red_Cards,
		Goals:             player.Goals,
		Assists:           player.Assists,
		Minutes_Played:    player.Minutes_Played,
		Player_Valuations: player.Player_Valuations,
	}

	return transformedPlayer
}
