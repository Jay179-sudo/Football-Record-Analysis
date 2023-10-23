package main

import (
	"log"
	"time"

	"github.com/Jay179-sudo/FootballRecordAnalysis/internal/data"
)

func getPlayerRecordTitles() []string {
	record_title := []string{"player_id", "name", "date_of_birth", "position"}
	return record_title
}

// ________________________STORES Player DATA INTO THE Player Table________________________
func (app *application) GetPlayerData() error {
	records := readCsvFile("./data/players.csv")

	title_rows := getPlayerRecordTitles()
	title_index := []int{}
	for i, row := range records {
		player_record := &data.Player{}
		for j, element := range row {

			if i == 0 {
				if search(title_rows, element) {
					title_index = append(title_index, j)
					log.Print(element, j)
				}
			} else {
				if j == title_index[0] {
					// Store Club ID
					result, err := stoi64(element)
					if err != nil {
						return err
					}
					player_record.Player_ID = result
				} else if j == title_index[1] {
					// Store Team Name
					if len(element) >= 30 {
						player_record.Player_Name = element[:30]
					} else {
						player_record.Player_Name = element
					}

				} else if j == title_index[2] {
					// Store Date of Birth
					result, err := time.Parse("02-01-2006", element)
					if err != nil {
						f, _ := time.Parse("02-01-2006", "02-01-2040")
						player_record.DOB = f
					}
					player_record.DOB = result

				} else if j == title_index[3] {
					if len(element) >= 30 {
						player_record.Position = element[:30]
					} else {
						player_record.Position = element
					}
				}

			}
			log.Print(j)
		}

		if player_record.Player_ID != 0 {
			// Push to DB
			err := app.models.Player.Insert(player_record)
			if err != nil {
				log.Fatal("Error in database write ", err)
				return err
			}

		}
	}

	return nil
}
