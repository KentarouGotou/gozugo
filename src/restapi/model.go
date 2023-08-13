package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Note struct {
	Position [6]string `json:"position"`
	Ties     [6]string `json:"ties"`
	Effects  [6]string `json:"effects"`
	Duration int       `json:"duration"`
}

type Tab struct {
	Id     int    `json:"id"`
	Artist string `json:"artist"`
	Title  string `json:"title"`
	Notes  []Note `json:"notes"`
	Tempo  int    `json:"tempo"`
}

// var tabs = []Tab{
// 	{
// 		Id: 1,
// 		Artist: "Beatles",
// 		Title: "Let it be",
// 		Notes: []Note{
// 			{
// 				Position: [6]string{"-", "1", "0", "2", "3", "-"},
// 				Ties: [6]string{"-", "-", "-", "-", "-", "-"},
// 				Effects: [6]string{"-", "vib", "-", "-", "-", "-"},
// 				Duration: 960
// 			},
// 			{
// 				Position: [6]string{"-", "1", "-", "-", "-", "-"},
// 				Ties: [6]string{"-", "t", "-", "-", "-", "-"},
// 				Effects: [6]string{"-", "vib", "-", "-", "-", "-"},
// 				Duration: 960
// 			}
// 			},
// 		Tempo: 120
// 	},
// 	{
// 		Id: 1,
// 		Artist: "Beatles",
// 		Title: "Let it be",
// 		 Notes: []Note{
// 			{
// 				Position: [6]string{"-", "-", "-", "7", "-", "-"},
// 				Ties: [6]string{"-", "-", "-", "-", "-", "-"},
// 				Effects: [6]string{"-", "-", "-", "bn2", "-", "-"},
// 				Duration: 480
// 			},
// 			{
// 				Position: [6]string{"-", "-", "-", "9", "10", "-"},
// 				Ties: [6]string{"-", "-", "-", "t", "-", "-"},
// 				Effects: [6]string{"-", "-", "-", "-", "-", "-"},
// 				Duration: 480
// 			}
// 		},
// 		Tempo: 120
// 	},
// }

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("sqlite3", "./example.sqlite")
	if err != nil {
		panic(err)
	}
}

func GetNotesforTab(id int) []Note {
	rows, err := Db.Query("select * from notes where tab_id = ?", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var notes []Note
	for rows.Next() {
		var note Note
		err := rows.Scan(&note.Position, &note.Ties, &note.Effects, &note.Duration)
		if err != nil {
			panic(err)
		}
		notes = append(notes, note)
	}
	return notes
}
