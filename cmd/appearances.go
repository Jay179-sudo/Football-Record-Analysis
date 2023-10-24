package main

import (
	"strconv"

	"github.com/Jay179-sudo/FootballRecordAnalysis/internal/data"
)

func getAppRecordTitles() []string {
	record_title := []string{"player_id", "player_club_id", "date", "yellow_cards", "red_cards", "goals", "assists", "minutes_played"}
	return record_title
}

// data from 2017 to 2019
type aggregation struct {
	yellow_cards     int32
	red_cards        int32
	goals            int32
	assists          int32
	minutes_played   int32
	player_valuation int64
}

// ________________________STORES APPEARANCES DATA INTO THE PLAYER_STATS TABLE________________________
func (app *application) GetPlayerStatsData() error {
	// store player valuation into map
	player_val, err := getPlayerValuationRecords()
	if err != nil {
		return err
	}
	records := readCsvFile("./data/appearances.csv")
	// map[string][{yellow_cards, red_cards, goals, assists}].
	// string = player_id + "#" + year(date) + "#" + current_club_id
	mp1 := make(map[string]*aggregation)
	title_rows := getAppRecordTitles()
	title_index := []int{}
	for i, row := range records {
		club_record := &data.Player_Stats{}
		for j, element := range row {

			if i == 0 {
				if search(title_rows, element) {
					title_index = append(title_index, j)
				}
			} else {
				if j == title_index[0] {
					// Store Player_ID
					result, err := stoi64(element)
					if err != nil {
						return err
					}
					club_record.Player_ID = result
				} else if j == title_index[1] {
					// Store Player_Club_ID
					result, err := stoi64(element)
					if err != nil {
						return err
					}
					club_record.Current_Club_ID = result

				} else if j == title_index[2] {
					// Store Date
					yearSlice := element[:4]
					result, err := stoi32(yearSlice)
					if err != nil {
						return err
					}
					if result < 2017 || result >= 2020 {
						break
					}
					club_record.Season = result

					// Get Player Valuation
					player_val_key := strconv.Itoa(int(club_record.Player_ID)) + "#" + yearSlice
					player_value, ok := player_val[player_val_key]
					if !ok {
						player_value = -1
					}

					// Use map for seasonal aggregation
					key := getKey(int(club_record.Player_ID), int(club_record.Season), int(club_record.Current_Club_ID))
					_, exists := mp1[key]
					if exists == false {
						agg := &aggregation{
							yellow_cards:     0,
							red_cards:        0,
							goals:            0,
							assists:          0,
							minutes_played:   0,
							player_valuation: player_value,
						}
						mp1[key] = agg
					}

				} else if j == title_index[3] {
					// aggregate yellow cards
					key := getKey(int(club_record.Player_ID), int(club_record.Season), int(club_record.Current_Club_ID))
					result, err := stoi32(element)
					if err != nil {
						return err
					}
					mp1[key].yellow_cards += result
				} else if j == title_index[4] {
					// aggregate red cards
					key := getKey(int(club_record.Player_ID), int(club_record.Season), int(club_record.Current_Club_ID))
					result, err := stoi32(element)
					if err != nil {
						return err
					}
					mp1[key].red_cards += result
				} else if j == title_index[5] {
					// aggregate goals
					key := getKey(int(club_record.Player_ID), int(club_record.Season), int(club_record.Current_Club_ID))
					result, err := stoi32(element)
					if err != nil {
						return err
					}
					mp1[key].goals += result
				} else if j == title_index[6] {
					// aggregate assists
					key := getKey(int(club_record.Player_ID), int(club_record.Season), int(club_record.Current_Club_ID))
					result, err := stoi32(element)
					if err != nil {
						return err
					}
					mp1[key].assists += result
				} else if j == title_index[7] {
					// aggregate minutes played
					key := getKey(int(club_record.Player_ID), int(club_record.Season), int(club_record.Current_Club_ID))
					result, err := stoi32(element)
					if err != nil {
						return err
					}
					mp1[key].minutes_played += result
				}

			}

		}

	}

	for key, element := range mp1 {
		player_id, season, club_id, err := separateKey(key)
		if err != nil {
			return err
		}
		player := data.Player_Stats{
			Player_ID:         player_id,
			Current_Club_ID:   club_id,
			Season:            season,
			Yellow_Cards:      element.yellow_cards,
			Red_Cards:         element.red_cards,
			Goals:             element.goals,
			Assists:           element.assists,
			Minutes_Played:    element.minutes_played,
			Player_Valuations: element.player_valuation,
		}
		app.models.Player_Stats.Insert(&player)
	}

	return nil
}
