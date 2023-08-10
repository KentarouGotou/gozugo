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

var tabs = []Tab{
	{
		Id: 1,
		Artist: "Beatles", 
		Title: "Let it be", 
		Notes: []Note{
			{
				Position: [6]string{"-", "1", "0", "2", "3", "-"}, 
				Ties: [6]string{"-", "-", "-", "-", "-", "-"}, 
				Effects: [6]string{"-", "vib", "-", "-", "-", "-"}, 
				Duration: 960
			},
			{
				Position: [6]string{"-", "1", "-", "-", "-", "-"}, 
				Ties: [6]string{"-", "t", "-", "-", "-", "-"}, 
				Effects: [6]string{"-", "vib", "-", "-", "-", "-"},
				Duration: 960
			}
			}, 
		Tempo: 120
	},
	{
		Id: 1, 
		Artist: "Beatles", 
		Title: "Let it be",
		 Notes: []Note{
			{
				Position: [6]string{"-", "-", "-", "7", "-", "-"}, 
				Ties: [6]string{"-", "-", "-", "-", "-", "-"}, 
				Effects: [6]string{"-", "-", "-", "bn2", "-", "-"}, 
				Duration: 480
			},
			{
				Position: [6]string{"-", "-", "-", "9", "10", "-"}, 
				Ties: [6]string{"-", "-", "-", "t", "-", "-"}, 
				Effects: [6]string{"-", "-", "-", "-", "-", "-"}, 
				Duration: 480
			}
		}, 
		Tempo: 120
	},
}

var DB *sql.DB
func init() {
	var err error
	DB, err = sql.Open("sqlite3", "./example.sqlite")
	if err != nil {
		panic(err)
	}
}

func getTabs()