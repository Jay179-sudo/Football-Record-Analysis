package main

import (
	"strconv"
	"strings"

	"github.com/Jay179-sudo/FootballRecordAnalysis/internal/data"
)

// __________________________ Generic Search _____________________
func search[T comparable](list []T, searchVal T) bool {
	for _, element := range list {
		if element == searchVal {
			return true
		}
	}
	return false
}

// ______________________ CONVERT STRING TO A int64 ________________
func stoi64(element string) (int64, error) {
	result, err := strconv.ParseInt(element, 10, 64)
	if err != nil {
		return -1, err
	}
	return result, nil
}

// ______________________ CONVERT STRING TO A int32 ________________
func stoi32(element string) (int32, error) {
	result, err := strconv.ParseInt(element, 10, 32)
	if err != nil {
		return -1, err
	}
	return int32(result), nil
}

// ______________________ CONVERT INPUT STRINGS TO A # SEPARATED KEY ________________
func getKey(playerId, Season, clubID int) string {
	key := strconv.Itoa(int(playerId))
	key += "#"
	key += strconv.Itoa(int(Season))
	key += "#"
	key += strconv.Itoa(int(clubID))

	return key
}

// ______________________ SPLIT KEY STRING ON THE BASIS OF SEPARATOR # ________________
func separateKey(key string) (int64, int32, int64, error) {
	result := strings.Split(key, "#")
	playerId, err := stoi64(result[0])
	if err != nil {
		return -1, -1, -1, err
	}

	season, err := stoi32(result[1])
	if err != nil {
		return -1, -1, -1, err
	}

	clubId, err := stoi64(result[2])
	if err != nil {
		return -1, -1, -1, err
	}
	return playerId, season, clubId, nil
}

// _________________________ FIVE NUMBER SUMMARY OF A NUMERICAL VARIABLE __________________
func (app *application) getFiveNumberSummary(list []data.Minutes) (int64, int64, int64, int64, int64, error) {

	if len(list) == 0 {
		return -1, -1, -1, -1, -1, nil
	}

	min := list[0].Minutes_Played
	max := list[0].Minutes_Played
	n := len(list)
	q1 := (n + 1) / 4
	q3 := 3 * (n + 1) / 4
	median := (n) / 2

	for index, element := range list {
		if element.Minutes_Played < min {
			min = element.Minutes_Played
		}
		if element.Minutes_Played > max {
			max = element.Minutes_Played
		}

		if q1 == index {
			q1 = int(element.Minutes_Played)
		}
		if median == index {
			median = int(element.Minutes_Played)
		}
		if q3 == index {
			q3 = int(element.Minutes_Played)
		}
	}

	return min, int64(q1), int64(median), int64(q3), max, nil

}

type Precomputations struct {
	MinutesLower int64
}
