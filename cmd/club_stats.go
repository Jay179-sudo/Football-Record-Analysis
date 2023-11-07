package main

import (
	"strconv"
	"strings"

	"github.com/Jay179-sudo/FootballRecordAnalysis/internal/data"
)

func getClubStatsRecordTitles() []string {
	record_title := []string{"date", "home_club_id", "away_club_id", "home_club_goals", "away_club_goals"}
	return record_title
}

// Database Struct
type club_stats struct {
	homeTeam  int32
	awayTeam  int32
	season    int32
	homeGoals int32
	awayGoals int32
}

// Aggregation struct to store in a map
type agg_club_stats struct {
	wins   int32
	losses int32
	draws  int32
}

// ________________ GET CLUB AGGREGATED STATISTICS ________________
func (app *application) GetClubStatsData() error {
	records := readCsvFile("./datasource/games.csv")
	title_rows := getClubStatsRecordTitles()
	title_index := []int{}
	mp1 := make(map[string]*agg_club_stats)
	for i, row := range records {
		var row_stats club_stats
		for j, element := range row {
			if i == 0 {
				if search(title_rows, element) {
					title_index = append(title_index, j)
				}
			} else {
				if j == title_index[0] {
					// Store Date
					yearSlice := element[:4]
					result, err := stoi32(yearSlice)
					if err != nil {
						return err
					}
					if result < 2017 || result >= 2023 {
						break
					}
					row_stats.season = result
				} else if j == title_index[1] {
					// Store away_club_id
					result, err := stoi32(element)
					if err != nil {
						return err
					}
					row_stats.homeTeam = result

				} else if j == title_index[2] {
					// Store home_club_id
					result, err := stoi32(element)
					if err != nil {
						return err
					}

					row_stats.awayTeam = result

				} else if j == title_index[3] {
					// home_club_goals
					result, err := stoi32(element)
					if err != nil {
						return err
					}
					row_stats.homeGoals = result

				} else if j == title_index[4] {
					// away_club_goals
					result, err := stoi32(element)
					if err != nil {
						return err
					}
					row_stats.awayGoals = result

				}

			}

		}
		if i > 0 {
			homeKey := strconv.Itoa(int(row_stats.homeTeam)) + "#" + strconv.Itoa(int(row_stats.season))
			awayKey := strconv.Itoa(int(row_stats.awayTeam)) + "#" + strconv.Itoa(int(row_stats.season))
			_, ok := mp1[homeKey]
			if !ok {
				agg := agg_club_stats{
					wins:   0,
					losses: 0,
					draws:  0,
				}
				mp1[homeKey] = &agg
			}
			_, ok = mp1[awayKey]
			if !ok {
				agg := agg_club_stats{
					wins:   0,
					losses: 0,
					draws:  0,
				}
				mp1[awayKey] = &agg
			}
			if row_stats.homeGoals > row_stats.awayGoals {
				mp1[homeKey].wins += 1
				mp1[awayKey].losses += 1
			} else if row_stats.awayGoals > row_stats.homeGoals {
				mp1[homeKey].losses += 1
				mp1[awayKey].wins += 1
			} else {
				mp1[homeKey].draws += 1
				mp1[awayKey].draws += 1
			}
		}

	}

	for key, element := range mp1 {
		separateKey := strings.Split(key, "#")
		club_id, err := strconv.Atoi(separateKey[0])
		if err != nil {
			return err
		}
		season, err := strconv.Atoi(separateKey[1])
		if err != nil {
			return err
		}
		club_stats := data.Club_Stats{
			Club_ID: int64(club_id),
			Season:  int64(season),
			Wins:    element.wins,
			Losses:  element.losses,
			Draws:   element.draws,
		}
		err = app.models.Club_Stats.Insert(club_stats)
		if err != nil {
			return err
		}
	}

	return nil

}
