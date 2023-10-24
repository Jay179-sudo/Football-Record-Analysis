package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

// __________________________ MOVE THESE FOLLOWING FUNCTIONS TO UTILS _____________________
func search[T comparable](list []T, searchVal T) bool {
	for _, element := range list {
		if element == searchVal {
			return true
		}
	}
	return false
}

func stoi64(element string) (int64, error) {
	result, err := strconv.ParseInt(element, 10, 64)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func stoi32(element string) (int32, error) {
	result, err := strconv.ParseInt(element, 10, 32)
	if err != nil {
		return -1, err
	}
	return int32(result), nil
}

func getKey(playerId, Season, clubID int) string {
	key := strconv.Itoa(int(playerId))
	key += "#"
	key += strconv.Itoa(int(Season))
	key += "#"
	key += strconv.Itoa(int(clubID))

	return key
}

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
