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

var Db *sql.DB
func init() {
	var err error
	DB, err = sql.Open("sqlite3", "./example.sqlite")
	if err != nil {
		panic(err)
	}
}

func getTabs(limit int) (tabs []Tab, err error) {
	stmt := "SELECT id, artist, title, tempo FROM tabs LIMIT $1"
	rows, err := Db.Query(stmt, limit)
	if err != nil {
		return
	}

	for rows.Next() {
		tab := Tab{}
		err = rows.Scan(&tab.Id, &tab.Artist, &tab.Title, &tab.Notes, &tab.Tempo)
		if err != nil {
			return
		}
		tabs = append(tabs, tab)
	}
	rows.Close()
	return
}

func getTab(id int) (tab Tab, err error) {
	tab = Tab{}
	stmt := "SELECT id, artist, title, tempo FROM tabs WHERE id = $1"
	err = Db.QueryRow(stmt, id).Scan(&tab.Id, &tab.Artist, &tab.Title, &tab.Notes, &tab.Tempo)
	return
}

func (tab *Tab) createTab() (err error) {
	stmt := "INSERT INTO tabs (artist, title, tempo) VALUES ($1, $2, $3) RETURNING id"
	err = Db.QueryRow(stmt, tab.Artist, tab.Title, tab.Tempo).Scan(&tab.Id)
	return
}