package main

import (
	"log"

	"github.com/Jay179-sudo/FootballRecordAnalysis/internal/data"
)

// From the Clubs csv extract Club_ID, Team Name and Stadium Capacity

func getClubRecordTitles() []string {
	record_title := []string{"club_id", "name", "stadium_seats"}
	return record_title
}

// ________________________STORES CLUB DATA INTO THE CLUB TABLE________________________
func (app *application) GetClubData() error {
	records := readCsvFile("./data/clubs.csv")

	title_rows := getClubRecordTitles()
	title_index := []int{}
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
					// Store Stadium Seats
					result, err := stoi64(element)
					if err != nil {
						return err
					}
					club_record.Stadium_Capacity = result

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
