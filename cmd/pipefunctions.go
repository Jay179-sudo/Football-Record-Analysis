package main

import (
	"log"

	"github.com/Jay179-sudo/FootballRecordAnalysis/internal/data"
)

// input a list of players, the generator will generate a player one by one in the channel
func (app *application) generator(done <-chan interface{}, players ...data.Player_Stats) <-chan data.Player_Stats {
	dataStream := make(chan data.Player_Stats)
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
		for i := range upStream {
			select {
			case <-done:
				return
			case dataStream <- cleaned(precomp, i):
			}
		}

	}()
	return dataStream
}

func (app *application) transformed(done <-chan interface{}, upStream <-chan data.Player_Stats, precomp Precomputations) <-chan data.Player_Stats {
	dataStream := make(chan data.Player_Stats)
	go func() {
		defer close(dataStream)
		for i := range upStream {
			select {
			case <-done:
				return
			case dataStream <- transformed(precomp, i):
			}

		}
	}()
	return dataStream
}

func (app *application) pipelineEnd(done <-chan interface{}, upStream <-chan data.Player_Stats) {
	dataStream := make(chan data.Player_Stats)
	go func() {
		defer close(dataStream)
		for dataVal := range upStream {
			select {
			case <-done:
				return
			default:
				// Store to database. Update or delete
				var err error
				if dataVal.Minutes_Played == -1 {
					err = app.models.Player_Stats.Delete(dataVal.Player_ID, dataVal.Current_Club_ID, dataVal.Season)
				} else {
					err = app.models.Player_Stats.Update(dataVal)
				}
				if err != nil {
					log.Fatal("Pipeline crashed.", err)
				}

			}
		}

	}()
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
		Yellow_Cards:      float32(player.Yellow_Cards) / float32(player.Minutes_Played) * 1000,
		Red_Cards:         float32(player.Red_Cards) / float32(player.Minutes_Played) * 1000,
		Goals:             float32(player.Goals) / float32(player.Minutes_Played) * 1000,
		Assists:           float32(player.Assists) / float32(player.Minutes_Played) * 1000,
		Minutes_Played:    player.Minutes_Played,
		Player_Valuations: player.Player_Valuations,
	}
	return transformedPlayer
}
