package main

import (
	"log"

	"github.com/Jay179-sudo/FootballRecordAnalysis/internal/data"
)

// From the Clubs csv extract Club_ID, Team Name and Stadium Capacity

func getClubRecordTitles() []string {
	record_title := []string{"club_id", "name", "domestic_competition_id"}
	return record_title
}

func getCompetitionTitles() []string {
	record_title := []string{"type", "domestic_league_code", "country_name"}
	return record_title
}

// ________________________STORES CLUB DATA INTO THE CLUB TABLE________________________
func (app *application) GetClubData() error {
	records := readCsvFile("./datasource/clubs.csv")
	// store club data
	title_rows := getCompetitionTitles()
	competitionCountry := readCsvFile("./datasource/competitions.csv")
	title_index := []int{}
	mp1 := make(map[string]string)
	for i, row := range competitionCountry {
		var competition_id string
		for j, element := range row {
			if i == 0 {
				if search(title_rows, element) {
					title_index = append(title_index, j)
				}
			} else {
				if j == title_index[0] {
					// type
					if element != "domestic_league" {
						break
					}

				} else if j == title_index[1] {
					// domestic_league_code
					competition_id = element

				} else if j == title_index[2] {
					// country name

					mp1[competition_id] = element

				}
			}
		}
	}
	title_rows = getClubRecordTitles()
	title_index = []int{}
	for i, row := range records {
		club_record := &data.Club{}
		for j, element := range row {

			if i == 0 {
				if search(title_rows, element) {
					title_index = append(title_index, j)
				}
			} else {
				if j == title_index[0] {
					// Store Club ID
					result, err := stoi64(element)
					if err != nil {
						return err
					}
					club_record.Club_ID = result
				} else if j == title_index[1] {
					// Store Team Name
					if len(element) >= 30 {
						club_record.Team_Name = element[:30]
					} else {
						club_record.Team_Name = element
					}

				} else if j == title_index[2] {
					// Store Country from preprocessed map

					// club_record.Country = mp1[element]
					_, exists := mp1[element]
					if !exists {
						club_record.Country = "N/A"
					} else {
						club_record.Country = mp1[element]

					}

				}

			}

		}
		if club_record.Club_ID != 0 {
			// Push to DB
			err := app.models.Club.Insert(club_record)
			if err != nil {
				log.Fatal("Error in database write ", err)
				return err
			}

		}
	}

	return nil
}
